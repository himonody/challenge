package service

import (
	authStorage "challenge/app/auth/storage"
	riskDto "challenge/app/risk/service/dto"
	userModels "challenge/app/user/models"
	userRepo "challenge/app/user/repo"
	baseConstant "challenge/config/base/constant"
	"challenge/core/dto/service"
	"time"
)

type AuthLogout struct {
	service.Service
}

// NewAuthLogoutService 实例化登出服务
func NewAuthLogoutService(s *service.Service) *AuthLogout {
	var srv = new(AuthLogout)
	srv.Orm = s.Orm
	srv.Log = s.Log
	srv.Run = s.Run
	srv.C = s.C
	return srv
}

// Logout 登出（记录完整日志）
func (a *AuthLogout) Logout(userID int, username string) error {
	ctx := a.C.Request.Context()
	now := time.Now()

	// 采集风控上下文
	rc := a.extractRiskContext()

	// 1. 清除Redis中的Token（如果有缓存）
	cache := a.Run.GetCacheAdapter()
	if cache != nil {
		_ = authStorage.DelLoginToken(ctx, cache, userID)
		// 也可以清除Session
		// _ = authStorage.DelLoginSession(ctx, cache, sessionID)
	}

	// 2. 记录登出的登录日志（可选：记录为特殊的登出状态）
	logoutLog := &userModels.UserLoginLog{
		UserID:     uint64(userID),
		LoginAt:    now,
		LoginIP:    rc.IP,
		DeviceFP:   rc.DeviceFP,
		UserAgent:  rc.UA,
		Status:     4, // 4表示登出（可以在数据库中定义）
		FailReason: "",
		CreatedAt:  &now,
	}
	if err := a.Orm.Create(logoutLog).Error; err != nil {
		a.Log.Errorf("challenge.app.auth.service.Logout.Create create user login log failed: %v", err)
	}

	// 3. 记录用户操作日志
	operLog := &userModels.AppUserOperLog{
		UserID:     userID,
		ActionType: baseConstant.UserActionLogout,
		OperateIP:  rc.IP,
		ByType:     baseConstant.UserOperByTypeApp,
		Status:     baseConstant.GeneralStatusOk,
		CreateBy:   userID,
		UpdateBy:   userID,
		CreatedAt:  now,
		UpdatedAt:  now,
		Remark:     baseConstant.OperLogRemarkLogoutSuccess,
	}
	if err := userRepo.CreateUserOperLog(a.Orm, operLog); err != nil {
		a.Log.Errorf("logout CreateUserOperLog error: %v", err)
	}

	a.Log.Infof("user %d(%s) logout from %s", userID, username, rc.IP)
	return nil
}

// extractRiskContext 提取风控上下文
func (a *AuthLogout) extractRiskContext() *riskDto.RiskContext {
	rc := &riskDto.RiskContext{}
	if a.C != nil {
		rc.IP = a.C.ClientIP()
		rc.UA = a.C.GetHeader("User-Agent")
		rc.DeviceFP = a.C.GetHeader("X-Device-FP")
	}
	return rc
}
