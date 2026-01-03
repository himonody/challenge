package service

import (
	"challenge/app/auth/service/dto"
	baseLang "challenge/config/base/lang"
	"challenge/core/dto/response"
	"challenge/core/lang"
	"challenge/core/middleware/auth"
	"challenge/core/middleware/auth/authdto"
	"challenge/core/utils/captchautils"
	"challenge/core/utils/encrypt"
	"challenge/core/utils/idgen"
	"errors"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	userPwdRegex = regexp.MustCompile(`^[A-Za-z0-9!@#$%^&*()\-_=+,.?/:;{}\[\]` + "`" + `~]{4,12}$`)
)

type userRow struct {
	Id       int64
	UserName string
	Pwd      string
	RefCode  string
	Status   string
}

type showColumnsRow struct {
	Field string `gorm:"column:Field"`
}

func Captcha(c *gin.Context) {
	id, b64s, _, err := captchautils.DriverStringFunc()
	if err != nil {
		response.Error(c, baseLang.RequestErr, "captcha generate failed")
		return
	}
	response.OK(c, &dto.CaptchaResp{CaptchaId: id, ImageBase64: b64s}, "")
}

func Register(c *gin.Context) {
	var req dto.RegisterReq
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

	// check duplicate username
	var existCnt int64
	db.Table("app_user").Where(userNameCol+" = ?", req.UserName).Count(&existCnt)
	if existCnt > 0 {
		response.Error(c, baseLang.RequestErr, "用户名已存在")
		return
	}

	parentId := int64(0)
	if req.RefCode != "" {
		var parent userRow
		if err := db.Table("app_user").Select("id").Where("ref_code = ?", req.RefCode).Take(&parent).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				response.Error(c, baseLang.RequestErr, "推荐码不存在")
				return
			}
			response.Error(c, baseLang.ServerErr, "db error")
			return
		}
		parentId = parent.Id
	}

	pwdHash, err := encrypt.HashEncrypt(req.Password)
	if err != nil {
		response.Error(c, baseLang.ServerErr, "password encrypt failed")
		return
	}

	refCode := idgen.InviteId()
	for i := 0; i < 3; i++ {
		var cnt int64
		db.Table("app_user").Where("ref_code = ?", refCode).Count(&cnt)
		if cnt == 0 {
			break
		}
		refCode = idgen.InviteId()
	}

	now := time.Now()
	create := map[string]interface{}{
		"level_id":   1,
		userNameCol:  req.UserName,
		pwdCol:       pwdHash,
		"nickname":   req.UserName,
		"ref_code":   refCode,
		"parent_id":  parentId,
		"status":     "1",
		"created_at": now,
		"updated_at": now,
		"create_by":  0,
		"update_by":  0,
		"tree_sort":  0,
		"tree_sorts": "0",
		"tree_leaf":  "1",
		"tree_level": 0,
		"parent_ids": "0,",
	}

	res := db.Table("app_user").Create(create)
	if res.Error != nil {
		response.Error(c, baseLang.DataInsertCode, "注册失败")
		return
	}

	// Auto login after register
	uid := int64(0)
	if idVal, ok := create["id"]; ok {
		_ = idVal
	}
	// fetch id
	var u userRow
	if err := db.Table("app_user").Select("id").Where(userNameCol+" = ?", req.UserName).Take(&u).Error; err == nil {
		uid = u.Id
	}
	if uid <= 0 {
		response.OK(c, gin.H{"userName": req.UserName}, "")
		return
	}
	c.Set(authdto.LoginUserId, uid)
	c.Set(authdto.UserName, req.UserName)
	auth.Auth.Login(c)
}

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

func getDb(c *gin.Context) (*gorm.DB, error) {
	v, ok := c.Get("db")
	if !ok || v == nil {
		return nil, errors.New("db not found")
	}
	db, ok := v.(*gorm.DB)
	if !ok || db == nil {
		return nil, errors.New("db invalid")
	}
	return db, nil
}

func resolveColumn(db *gorm.DB, table string, preferred string, fallback string) string {
	var rows []showColumnsRow
	_ = db.Raw("SHOW COLUMNS FROM "+table+" LIKE ?", preferred).Scan(&rows).Error
	if len(rows) > 0 {
		return preferred
	}
	return fallback
}
