package service

import (
	"challenge/app/auth/service/dto"
	authStorage "challenge/app/auth/storage"
	riskDto "challenge/app/risk/service/dto"
	userModels "challenge/app/user/models"
	userRepo "challenge/app/user/repo"
	baseConstant "challenge/config/base/constant"
	"challenge/config/base/lang"
	"challenge/core/dto/service"
	"challenge/core/utils/captchautils"
	"challenge/core/utils/strutils"
	"errors"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthLogin struct {
	service.Service
}

// NewAuthLoginService 实例化登录服务
func NewAuthLoginService(s *service.Service) *AuthLogin {
	var srv = new(AuthLogin)
	srv.Orm = s.Orm
	srv.Log = s.Log
	srv.Run = s.Run
	srv.C = s.C
	return srv
}

// Login 登录（含完整风控流程）
func (a *AuthLogin) Login(req *dto.LoginReq) (*userModels.AppUser, int) {
	ctx := a.C.Request.Context()

	// 分布式锁防重复登录请求（按用户名）
	locker := a.Run.GetLockerPrefix(authStorage.AuthLockPrefix)
	lock, err := authStorage.LockAuthAction(ctx, locker, "login", req.UserName, 10)
	if err != nil {
		return nil, lang.ServerErr
	}
	if lock != nil {
		defer lock.Release(ctx)
	}

	// 0. 采集风控上下文
	rc := a.extractRiskContext()

	// 1. 参数校验
	req.UserName = strings.TrimSpace(req.UserName)
	req.Password = strings.TrimSpace(req.Password)
	req.CaptchaId = strings.TrimSpace(req.CaptchaId)
	req.CaptchaCode = strings.TrimSpace(req.CaptchaCode)

	if !strutils.IsValidUsername(req.UserName) {
		return nil, lang.AuthUsernameErrorCode
	}
	if !strutils.IsValidPassword(req.Password) {
		return nil, lang.AuthPasswordErrorCode
	}
	if req.CaptchaId == "" || req.CaptchaCode == "" {
		return nil, lang.AuthVerificationCodeErrorCode
	}
	if !captchautils.Verify(req.CaptchaId, req.CaptchaCode, true) {
		return nil, lang.AuthVerificationCodeErrorCode
	}

	// 2. 风控检查（黑名单+锁定+失败次数）
	riskSvc := NewRiskCheckService(&a.Service)
	if code, err := riskSvc.CheckLoginRisk(ctx, req.UserName, rc); code != lang.SuccessCode {
		a.Log.Warnf("login risk check failed: %v", err)

		// 记录风控拦截日志
		now := time.Now()
		riskLog := &userModels.UserLoginLog{
			UserID:     0, // 风控拦截时可能还不知道用户ID
			LoginAt:    now,
			LoginIP:    rc.IP,
			DeviceFP:   rc.DeviceFP,
			UserAgent:  rc.UA,
			Status:     3, // 风控拦截
			FailReason: baseConstant.LoginFailReasonRiskBlock,
			CreatedAt:  &now,
		}
		_ = a.Orm.Create(riskLog).Error

		return nil, code
	}

	// 3. 查询用户
	var user userModels.AppUser
	err = a.Orm.Table("app_user").
		Where("username = ?", req.UserName).
		First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 用户不存在，记录失败
			now := time.Now()

			// 记录登录失败日志
			notFoundLog := &userModels.UserLoginLog{
				UserID:     0, // 用户不存在
				LoginAt:    now,
				LoginIP:    rc.IP,
				DeviceFP:   rc.DeviceFP,
				UserAgent:  rc.UA,
				Status:     2, // 失败
				FailReason: baseConstant.LoginFailReasonUserNotFound,
				CreatedAt:  &now,
			}
			_ = a.Orm.Create(notFoundLog).Error

			// 风控记录
			_ = riskSvc.RecordLoginFail(ctx, 0, req.UserName, rc)

			return nil, lang.DataNotFoundCode
		}
		a.Log.Errorf("login query user error: %v", err)
		return nil, lang.ServerErr
	}

	// 4. 检查用户状态
	if user.Status != baseConstant.GeneralStatusOk {
		return nil, lang.ForbitErr
	}

	// 5. 校验密码
	if err = bcrypt.CompareHashAndPassword([]byte(user.Pwd), []byte(req.Password)); err != nil {
		// 密码错误，记录失败
		now := time.Now()

		// 记录登录失败日志
		failLog := &userModels.UserLoginLog{
			UserID:     uint64(user.ID),
			LoginAt:    now,
			LoginIP:    rc.IP,
			DeviceFP:   rc.DeviceFP,
			UserAgent:  rc.UA,
			Status:     2, // 失败
			FailReason: baseConstant.LoginFailReasonPasswordError,
			CreatedAt:  &now,
		}
		_ = a.Orm.Create(failLog).Error

		// 记录操作日志
		operLog := &userModels.AppUserOperLog{
			UserID:     user.ID,
			ActionType: baseConstant.UserActionLoginFail,
			OperateIP:  rc.IP,
			ByType:     baseConstant.UserOperByTypeApp,
			Status:     baseConstant.GeneralStatusAbnormal,
			CreateBy:   user.ID,
			UpdateBy:   user.ID,
			CreatedAt:  now,
			UpdatedAt:  now,
			Remark:     baseConstant.OperLogRemarkLoginFailPassword,
		}
		_ = userRepo.CreateUserOperLog(a.Orm, operLog)

		// 风控记录失败
		_ = riskSvc.RecordLoginFail(ctx, user.ID, req.UserName, rc)

		return nil, lang.AuthPasswordErrorCode
	}

	// 6. 登录成功
	// 更新最后登录时间和IP
	now := time.Now()
	_ = a.Orm.Model(&userModels.AppUser{}).
		Where("id = ?", user.ID).
		Updates(map[string]interface{}{
			"last_login_at": now,
			"last_login_ip": rc.IP,
			"updated_at":    now,
		}).Error

	// 记录登录成功（清除失败计数+记录事件）
	_ = riskSvc.RecordLoginSuccess(ctx, user.ID, req.UserName, rc)

	// 记录登录日志
	loginLog := &userModels.UserLoginLog{
		UserID:    uint64(user.ID),
		LoginAt:   now,
		LoginIP:   rc.IP,
		DeviceFP:  rc.DeviceFP,
		UserAgent: rc.UA,
		Status:    1, // 成功
		CreatedAt: &now,
	}
	_ = a.Orm.Create(loginLog).Error

	// 记录用户操作日志
	operLog := &userModels.AppUserOperLog{
		UserID:     user.ID,
		ActionType: baseConstant.UserActionLogin,
		OperateIP:  rc.IP,
		ByType:     baseConstant.UserOperByTypeApp,
		Status:     baseConstant.GeneralStatusOk,
		CreateBy:   user.ID,
		UpdateBy:   user.ID,
		CreatedAt:  now,
		UpdatedAt:  now,
		Remark:     baseConstant.OperLogRemarkLoginSuccess,
	}
	_ = userRepo.CreateUserOperLog(a.Orm, operLog)

	return &user, lang.SuccessCode
}

// extractRiskContext 提取风控上下文
func (a *AuthLogin) extractRiskContext() *riskDto.RiskContext {
	rc := &riskDto.RiskContext{}
	if a.C != nil {
		rc.IP = a.C.ClientIP()
		rc.UA = a.C.GetHeader("User-Agent")
		rc.DeviceFP = a.C.GetHeader("X-Device-FP")
	}
	return rc
}
