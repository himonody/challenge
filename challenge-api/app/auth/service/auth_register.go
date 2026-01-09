package service

import (
	"challenge/app/auth/service/dto"
	authStorage "challenge/app/auth/storage"
	riskDto "challenge/app/risk/service/dto"
	userModels "challenge/app/user/models"
	"challenge/app/user/repo"
	baseConstant "challenge/config/base/constant"
	"challenge/config/base/lang"
	"challenge/core/dto/service"
	"challenge/core/utils/captchautils"
	"challenge/core/utils/encrypt"
	"challenge/core/utils/idgen"
	"challenge/core/utils/strutils"
	"fmt"
	"strings"
	"time"
)

type AuthRegister struct {
	service.Service
}

// NewAuthRegisterService 实例化注册服务
func NewAuthRegisterService(s *service.Service) *AuthRegister {
	var srv = new(AuthRegister)
	srv.Orm = s.Orm
	srv.Log = s.Log
	srv.Run = s.Run
	srv.C = s.C
	return srv
}

// Register 注册（含完整风控流程）
func (a *AuthRegister) Register(req *dto.RegisterReq) (*userModels.AppUser, int) {
	ctx := a.C.Request.Context()

	// 分布式锁防重复注册请求（按用户名）
	locker := a.Run.GetLockerPrefix(authStorage.AuthLockPrefix)
	lock, err := authStorage.LockAuthAction(ctx, locker, baseConstant.AuthRegisterAction, strings.TrimSpace(req.UserName), 10)
	if err != nil {
		return nil, lang.RepeatOperationCode
	}
	if lock != nil {
		defer lock.Release(ctx)
	}

	// 0. 采集风控上下文
	rc := a.extractRiskContext()

	// 1. 参数校验
	req.UserName = strings.TrimSpace(req.UserName)
	req.Password = strings.TrimSpace(req.Password)
	req.RefCode = strings.TrimSpace(req.RefCode)
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

	// 2. 风控检查（黑名单+限流）
	riskSvc := NewRiskCheckService(&a.Service)
	if code, err := riskSvc.CheckRegisterRisk(ctx, rc); code != lang.SuccessCode {
		a.Log.Errorf("challenge.app.auth.service.Register.CheckRegisterRisk risk check failed: %v", err)
		return nil, code
	}

	// 3. 检查用户名是否存在
	var existCnt int64
	a.Orm.Table("app_user").Where("username = ?", req.UserName).Count(&existCnt)
	if existCnt > 0 {
		return nil, lang.AuthUserAlreadyExistsCode
	}

	// 4. 处理邀请码（带 FOR UPDATE 锁）
	parentId := 0
	if req.RefCode != "" {
		var parent userModels.AppUser
		if err := a.Orm.Clauses( /*clause.Locking{Strength: "UPDATE"}*/ ).
			Table("app_user").Select("id").
			Where("ref_code = ?", req.RefCode).
			Take(&parent).Error; err != nil {
			return nil, lang.AuthInviteCodeNotFoundErrorCode
		}
		parentId = parent.ID
	}

	// 5. 密码hash
	pwdHash, err := encrypt.HashEncrypt(req.Password)
	if err != nil {
		return nil, lang.DataDecodeCode
	}

	// 6. 生成推荐码
	refCode := idgen.InviteId()
	for i := 0; i < 3; i++ {
		var cnt int64
		a.Orm.Table("app_user").Where("ref_code = ?", refCode).Count(&cnt)
		if cnt == 0 {
			break
		}
		refCode = idgen.InviteId()
	}

	// 7. 开启事务
	tx := a.Orm.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 8. 创建用户
	now := time.Now()
	user := &userModels.AppUser{
		Username:  req.UserName,
		Nickname:  req.UserName,
		Pwd:       pwdHash,
		RefCode:   refCode,
		ParentID:  parentId,
		Status:    baseConstant.GeneralStatusOk,
		CreatedAt: now,
		UpdatedAt: now,
		TreeSort:  baseConstant.UserTreeSort0,
		TreeLeaf:  baseConstant.UserTreeLeafYes,
		TreeLevel: baseConstant.UserTreeLevel,
	}
	if parentId != 0 {
		user.ParentIDs = fmt.Sprintf("%d,", parentId)
	}

	if err = repo.CreateUser(tx, user); err != nil {
		a.Log.Errorf("challenge.app.auth.service.Register.CreateUser create user failed: %v,user:%v", err, user)
		tx.Rollback()
		return nil, lang.DataInsertCode
	}

	// 9. 记录用户操作日志（注册）
	operLog := &userModels.AppUserOperLog{
		UserID:     user.ID,
		ActionType: baseConstant.UserActionRegister,
		OperateIP:  rc.IP,
		ByType:     baseConstant.UserOperByTypeApp,
		Status:     baseConstant.GeneralStatusOk,
		CreateBy:   user.ID,
		UpdateBy:   user.ID,
		CreatedAt:  now,
		UpdatedAt:  now,
		Remark:     baseConstant.OperLogRemarkRegisterSuccess,
	}
	if err = repo.CreateUserOperLog(tx, operLog); err != nil {
		a.Log.Errorf("register CreateUserOperLog error: %v", err)
		tx.Rollback()
		return nil, lang.ServerErr
	}

	// 记录初始登录日志（注册即登录）
	loginLog := &userModels.UserLoginLog{
		UserID:     uint64(user.ID),
		LoginAt:    now,
		LoginIP:    rc.IP,
		DeviceFP:   rc.DeviceFP,
		UserAgent:  rc.UA,
		Status:     1, // 成功
		FailReason: "",
		CreatedAt:  &now,
	}
	if err = tx.Create(loginLog).Error; err != nil {
		a.Log.Errorf("register CreateLoginLog error: %v", err)
		tx.Rollback()
		return nil, lang.ServerErr
	}

	// 10. 绑定设备
	if err = riskSvc.BindDevice(user.ID, rc.DeviceFP); err != nil {
		a.Log.Warnf("register BindDevice error: %v", err)
	}

	// 11. 初始化风控用户
	if err = riskSvc.InitRiskUser(user.ID); err != nil {
		a.Log.Warnf("register InitRiskUser error: %v", err)
	}

	// 12. 提交事务
	if err = tx.Commit().Error; err != nil {
		a.Log.Errorf("register Commit error: %v", err)
		tx.Rollback()
		return nil, lang.ServerErr
	}

	// 13. 记录注册成功（更新限流计数）
	_ = riskSvc.RecordRegisterSuccess(ctx, rc)

	return user, lang.SuccessCode
}

// extractRiskContext 提取风控上下文
func (a *AuthRegister) extractRiskContext() *riskDto.RiskContext {
	rc := &riskDto.RiskContext{}
	if a.C != nil {
		rc.IP = a.C.ClientIP()
		rc.UA = a.C.GetHeader("User-Agent")
		rc.DeviceFP = a.C.GetHeader("X-Device-FP")
	}
	return rc
}
