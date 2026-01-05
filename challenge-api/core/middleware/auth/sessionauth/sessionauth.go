package sessionauth

import (
	baseLang "challenge/config/base/lang"
	"challenge/core/config"
	"challenge/core/lang"
	"challenge/core/middleware/auth/authdto"
	"challenge/core/runtime"
	"challenge/core/utils/idgen"
	"challenge/core/utils/log"
	"challenge/core/utils/strutils"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	SessionLoginPrefixTmp = "challenge:login:session:tmp" //登录中转
	SessionLoginPrefix    = "challenge:login:session:user"
)

type SessionAuth struct{}

func (s *SessionAuth) Init() {}

func (s *SessionAuth) Login(c *gin.Context) {
	errResp := authdto.Resp{
		RequestId: strutils.GenerateMsgIDFromContext(c),
		Msg:       lang.MsgByCode(baseLang.RequestErr, lang.GetAcceptLanguage(c)),
		Code:      baseLang.RequestErr,
		Data:      nil,
	}

	userId := c.GetInt64(authdto.UserId)
	if userId <= 0 {
		c.JSON(baseLang.RequestErr, errResp)
		return
	}

	cache := runtime.RuntimeConfig.GetCacheAdapter()

	//获取sid，并用sid保存userId
	sid := idgen.UUID()
	err := cache.Set(SessionLoginPrefixTmp, sid, userId, config.AuthConfig.Timeout)
	rLog := log.GetRequestLogger(c)
	if err != nil {
		rLog.Error(err)
		c.JSON(baseLang.RequestErr, errResp)
		return
	}
	if config.ApplicationConfig.IsSingleLogin {
		_ = cache.Del(SessionLoginPrefix, strconv.FormatInt(userId, 10))
	}

	//session信息
	userName, _ := c.Get(authdto.Username)
	sessionInfo, err := json.Marshal(map[string]interface{}{
		authdto.UserId:   userId,
		authdto.Username: userName,
	})
	if err != nil {
		rLog.Error(err)
		c.JSON(baseLang.RequestErr, errResp)
		return
	}
	values := map[string]interface{}{}
	vs, _ := cache.HashGetAll(SessionLoginPrefix, strconv.FormatInt(userId, 10))
	if vs != nil {
		for k, v := range vs {
			values[k] = v
		}
	}
	values[sid] = string(sessionInfo)

	//用userId保存sid，记录登录状态（此操作可用于多点登录）
	err = cache.HashSet(config.AuthConfig.Timeout, SessionLoginPrefix, strconv.FormatInt(userId, 10), values)
	if err != nil {
		rLog.Error(err)
		c.JSON(baseLang.RequestErr, errResp)
		return
	}
	//userInfo, _ := c.Get(authdto.UserInfo)
	resp := authdto.Resp{
		RequestId: strutils.GenerateMsgIDFromContext(c),
		Msg:       "",
		Code:      http.StatusOK,
		Data: authdto.Data{
			Token:    sid,
			UserName: userName.(string),
			Expire:   time.Now().Add(time.Duration(config.AuthConfig.Timeout) * time.Second).Format(time.RFC3339),
			//UserInfo: userInfo,
		},
	}
	c.JSON(http.StatusOK, resp)
}

func (s *SessionAuth) Logout(c *gin.Context) {
	userId := c.GetInt64(authdto.UserId)
	if userId <= 0 {
		return
	}
	_ = runtime.RuntimeConfig.GetCacheAdapter().Del(SessionLoginPrefix, strconv.FormatInt(userId, 10))
	c.JSON(http.StatusOK, authdto.Resp{
		RequestId: strutils.GenerateMsgIDFromContext(c),
		Msg:       "",
		Code:      http.StatusOK,
		Data:      nil,
	})
}

func (s *SessionAuth) Get(c *gin.Context, key string) (interface{}, int, error) {
	var err error
	defer func() {
		if err != nil {
			rLog := log.GetRequestLogger(c)
			rLog.Error(strutils.GetCurrentTimeStr() + " [ERROR] " + c.Request.Method + " " + c.Request.URL.Path + " Get no " + key)
		}
	}()
	cache := runtime.RuntimeConfig.GetCacheAdapter()
	sid := strings.Replace(c.Request.Header.Get(authdto.HeaderAuthorization), authdto.HeaderTokenName+" ", "", -1)
	uid, err := cache.Get(SessionLoginPrefixTmp, sid)
	if sid == "" || uid == "" || err != nil {
		err = lang.MsgErr(baseLang.AuthErr, lang.GetAcceptLanguage(c))
		return "", baseLang.AuthErr, err
	}
	userInfoStr, err := cache.HashGet(SessionLoginPrefix, uid, sid)
	userInfo := map[string]interface{}{}
	err = json.Unmarshal([]byte(userInfoStr), &userInfo)
	if err != nil || userInfo[key] == nil {
		return "", baseLang.AuthErr, lang.MsgErr(baseLang.AuthErr, lang.GetAcceptLanguage(c))
	}
	return userInfo[key], baseLang.SuccessCode, nil
}

func (s *SessionAuth) GetUserId(c *gin.Context) (int64, int, error) {
	result, respCode, err := s.Get(c, authdto.UserId)
	if err != nil {
		return 0, respCode, err
	}
	return int64(result.(float64)), respCode, err
}
func (s *SessionAuth) GetUserName(c *gin.Context) string {
	result, _, _ := s.Get(c, authdto.Username)
	if result == nil {
		return ""
	}
	return result.(string)
}
func (s *SessionAuth) AuthMiddlewareFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		cache := runtime.RuntimeConfig.GetCacheAdapter()
		sid := strings.Replace(c.Request.Header.Get(authdto.HeaderAuthorization), authdto.HeaderTokenName+" ", "", -1)
		isExist := cache.Exist(SessionLoginPrefixTmp, sid)
		errResp := authdto.Resp{
			RequestId: strutils.GenerateMsgIDFromContext(c),
			Msg:       lang.MsgByCode(baseLang.AuthErr, lang.GetAcceptLanguage(c)),
			Code:      baseLang.AuthErr,
			Data:      nil,
		}
		if !isExist {
			c.JSON(baseLang.AuthErr, errResp)
			c.Abort()
			return
		}

		// 从session中获取用户id,第一次用于缓存拿到uid，第二次用uid检测sid是否有效，可用于多端登录
		uid, err := cache.Get(SessionLoginPrefixTmp, sid)
		if err != nil {
			c.JSON(baseLang.AuthErr, errResp)
			c.Abort()
			return
		}
		_, err = cache.HashGet(SessionLoginPrefix, uid, sid)
		if err != nil {
			c.JSON(baseLang.AuthErr, errResp)
			c.Abort()
			return
		}
		c.Set(authdto.UserId, uid)
		_ = cache.Expire(SessionLoginPrefixTmp, sid, config.AuthConfig.Timeout)
		_ = cache.Expire(SessionLoginPrefix, uid, config.AuthConfig.Timeout)
		c.Next()
	}
}
