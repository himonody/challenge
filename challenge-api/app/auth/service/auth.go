package service

import (
	"challenge/app/auth/service/dto"
	"challenge/app/user/models"
	m "challenge/app/user/models"
	"challenge/app/user/repo"
	baseLang "challenge/config/base/lang"
	"challenge/core/dto/response"
	"challenge/core/dto/service"
	"challenge/core/lang"
	"challenge/core/middleware/auth"
	"challenge/core/middleware/auth/authdto"
	"challenge/core/utils/captchautils"
	"challenge/core/utils/encrypt"
	"challenge/core/utils/idgen"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Auth struct {
	service.Service
}

// NewAuthService app-实例用户配置管理记录
func NewAuthService(s *service.Service) *Auth {
	var srv = new(Auth)
	srv.Orm = s.Orm
	srv.Log = s.Log
	srv.Run = s.Run
	return srv
}

func (a *Auth) Register(req *dto.RegisterReq) (models.AppUser, int, error) {
	req.UserName = strings.TrimSpace(req.UserName)
	req.Password = strings.TrimSpace(req.Password)
	req.RefCode = strings.TrimSpace(req.RefCode)
	req.CaptchaId = strings.TrimSpace(req.CaptchaId)
	req.CaptchaCode = strings.TrimSpace(req.CaptchaCode)

	if !userPwdRegex.MatchString(req.UserName) {
		return errors.New("用户名格式错误")
	}
	if !userPwdRegex.MatchString(req.Password) {
		return errors.New("密码格式错误")
	}
	if req.CaptchaId == "" || req.CaptchaCode == "" {
		return errors.New("验证码错误")
	}
	if !captchautils.Verify(req.CaptchaId, req.CaptchaCode, true) {
		return errors.New("验证码错误")
	}
	var existCnt int64
	a.Orm.Table("app_user").Where("username = ?", req.UserName).Count(&existCnt)
	if existCnt > 0 {
		return errors.New("用户名已存在")
	}
	parentId := 0
	if req.RefCode != "" {
		var parent models.AppUser
		if err := a.Orm.Table("app_user").Select("id").Where("ref_code = ?", req.RefCode).Take(&parent).Error; err != nil {
			return errors.New("推荐码不存在")
		}
		parentId = parent.ID
	}
	pwdHash, err := encrypt.HashEncrypt(req.Password)
	if err != nil {
		return errors.New("密码错误")
	}

	refCode := idgen.InviteId()
	for i := 0; i < 3; i++ {
		var cnt int64
		a.Orm.Table("app_user").Where("ref_code = ?", refCode).Count(&cnt)
		if cnt == 0 {
			break
		}
		refCode = idgen.InviteId()
	}

	now := time.Now()
	user := new(models.AppUser)
	user.Username = req.UserName
	user.Nickname = req.UserName
	user.Pwd = pwdHash
	user.RefCode = refCode
	user.ParentID = parentId
	user.Status = "1"
	user.CreatedAt = now
	user.UpdatedAt = now
	user.CreateBy = 0
	user.UpdateBy = 0
	user.TreeSort = 0
	user.TreeSorts = "0"
	user.TreeLeaf = "1"
	user.TreeLevel = 0
	if parentId != 0 {
		user.ParentIDs = fmt.Sprintf("%d,", parentId)
	}

	tx := a.Orm.Begin()

	err = repo.CreateUser(tx, user)
	if err != nil {
		a.Log.Errorf("app.auth.service.Register  CreateUser req:%v error:%w", user, err)
		tx.Rollback()
		return err
	}
	log := new(models.AppUserOperLog)
	log.UserID = user.ID
	log.ActionType = m.Register
	log.ByType = "1"
	log.Status = "1"
	log.CreateBy = user.ID
	log.UpdateBy = user.ID
	log.CreatedAt = now
	log.UpdatedAt = now
	if err = repo.CreateUserOperLog(tx, log); err != nil {
		a.Log.Errorf("app.auth.service.Register  CreateUserOperLog req:%v error:%w", log, err)
		tx.Rollback()
		return err
	}
	if err = tx.Commit().Error; err != nil {
		a.Log.Errorf("app.auth.service.Register  Commit req:%v error:%w", user, err)
		tx.Rollback()
		return err
	}
	a.C.Set(authdto.LoginUserId, user.ID)
	a.C.Set(authdto.UserName, user.Username)
	auth.Auth.Login(c)
	return nil
}

var (
	userPwdRegex = regexp.MustCompile(`^[A-Za-z0-9!@#$%^&*()\-_=+,.?/:;{}\[\]` + "`" + `~]{4,12}$`)
)

func Login(c *gin.Context) {
	var req dto.LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, baseLang.ParamErrCode, lang.MsgByCode(baseLang.ParamErrCode, lang.GetAcceptLanguage(c)))
		return
	}
	if !userPwdRegex.MatchString(req.UserName) {
		response.Error(c, baseLang.ParamErrCode, "用户名格式错误")
		return
	}
	if !userPwdRegex.MatchString(req.Password) {
		response.Error(c, baseLang.ParamErrCode, "密码格式错误")
		return
	}
	if req.CaptchaId == "" || req.CaptchaCode == "" {
		response.Error(c, baseLang.ParamErrCode, "验证码不能为空")
		return
	}
	if !captchautils.Verify(req.CaptchaId, req.CaptchaCode, true) {
		response.Error(c, baseLang.ParamErrCode, "验证码错误")
		return
	}

	db, err := getDb(c)
	if err != nil {
		response.Error(c, baseLang.ServerErr, "db error")
		return
	}

	userNameCol := resolveColumn(db, "app_user", "user_name", "username")
	pwdCol := resolveColumn(db, "app_user", "pwd", "password")

	var u userRow
	q := db.Table("app_user").Select("id, " + userNameCol + " as user_name, " + pwdCol + " as pwd, status")
	if err := q.Where(userNameCol+" = ?", req.UserName).Take(&u).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(c, baseLang.RequestErr, "用户名或密码错误")
			return
		}
		response.Error(c, baseLang.ServerErr, "db error")
		return
	}

	if u.Status != "" && u.Status != "1" {
		response.Error(c, baseLang.ForbitErr, "账号已禁用")
		return
	}
	if u.Pwd == "" {
		response.Error(c, baseLang.RequestErr, "用户名或密码错误")
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(u.Pwd), []byte(req.Password)); err != nil {
		response.Error(c, baseLang.RequestErr, "用户名或密码错误")
		return
	}

	c.Set(authdto.LoginUserId, u.Id)
	c.Set(authdto.UserName, req.UserName)
	auth.Auth.Login(c)
}

func Logout(c *gin.Context) {
	auth.Auth.Logout(c)
}
