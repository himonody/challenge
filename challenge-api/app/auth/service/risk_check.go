package service

import (
	authStorage "challenge/app/auth/storage"
	"challenge/app/risk/models"
	riskRepo "challenge/app/risk/repo"
	riskDto "challenge/app/risk/service/dto"
	riskStorage "challenge/app/risk/storage"
	baseConstant "challenge/config/base/constant"
	baseLang "challenge/config/base/lang"
	"challenge/core/dto/service"
	"challenge/core/lang"
	"challenge/core/utils/storage"
	"context"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type RiskCheck struct {
	service.Service
}

// NewRiskCheckService 实例化风控检查服务
func NewRiskCheckService(s *service.Service) *RiskCheck {
	var srv = new(RiskCheck)
	srv.Orm = s.Orm
	srv.Log = s.Log
	srv.Run = s.Run
	srv.Lang = s.Lang
	return srv
}

// getCache 获取缓存适配器
func (r *RiskCheck) getCache() storage.AdapterCache {
	if r.Run == nil {
		return nil
	}
	return r.Run.GetCacheAdapter()
}

// CheckRegisterRisk 注册风控检查
func (r *RiskCheck) CheckRegisterRisk(ctx context.Context, rc *riskDto.RiskContext) (int, error) {
	cache := r.getCache()
	if cache == nil {
		r.Log.Warnf("risk check: cache not available, skip rate limit")
		return baseLang.SuccessCode, nil
	}

	// 1. IP 限流检查（1分钟3次）
	if hit, _ := authStorage.CheckRegisterIPLimit(ctx, cache, rc.IP, 3); hit {
		return baseLang.RiskRegisterIPLimitCode, lang.MsgErr(baseLang.RiskRegisterIPLimitCode, r.Lang)
	}

	// 2. 设备限流检查（24小时2次）
	if rc.DeviceFP != "" {
		if hit, _ := authStorage.CheckRegisterDeviceLimit(ctx, cache, rc.DeviceFP, 2); hit {
			return baseLang.RiskRegisterDeviceLimitCode, lang.MsgErr(baseLang.RiskRegisterDeviceLimitCode, r.Lang)
		}
	}

	// 3. 黑名单检查
	if hit, _ := r.checkBlacklist(ctx, rc); hit {
		return baseLang.RiskBlacklistHitCode, lang.MsgErr(baseLang.RiskBlacklistHitCode, r.Lang)
	}

	return baseLang.SuccessCode, nil
}

// RecordRegisterSuccess 记录注册成功（更新限流计数）
func (r *RiskCheck) RecordRegisterSuccess(ctx context.Context, rc *riskDto.RiskContext) error {
	cache := r.getCache()
	if cache == nil {
		return nil
	}
	_ = authStorage.IncrRegisterIPCount(ctx, cache, rc.IP, 60) // 1分钟窗口
	if rc.DeviceFP != "" {
		_ = authStorage.IncrRegisterDeviceCount(ctx, cache, rc.DeviceFP, 86400) // 24小时窗口
	}
	return nil
}

// CheckLoginRisk 登录风控检查
func (r *RiskCheck) CheckLoginRisk(ctx context.Context, username string, rc *riskDto.RiskContext) (int, error) {
	cache := r.getCache()
	if cache == nil {
		return baseLang.SuccessCode, nil
	}

	// 1. 黑名单检查
	if hit, _ := r.checkBlacklist(ctx, rc); hit {
		return baseLang.RiskBlacklistHitCode, lang.MsgErr(baseLang.RiskBlacklistHitCode, r.Lang)
	}

	// 2. 用户锁定检查
	if locked, _ := authStorage.IsLoginLocked(ctx, cache, username); locked {
		return baseLang.RiskLoginLockedCode, lang.MsgErr(baseLang.RiskLoginLockedCode, r.Lang)
	}

	// 3. 失败次数检查
	failCount, _ := authStorage.GetLoginFailCount(ctx, cache, username)
	if failCount >= 5 {
		return baseLang.RiskLoginLockedCode, lang.MsgErr(baseLang.RiskLoginLockedCode, r.Lang)
	}

	return baseLang.SuccessCode, nil
}

// RecordLoginSuccess 记录登录成功
func (r *RiskCheck) RecordLoginSuccess(ctx context.Context, userID int, username string, rc *riskDto.RiskContext) error {
	cache := r.getCache()
	if cache != nil {
		// 清除失败计数
		_ = authStorage.ClearLoginFailCount(ctx, cache, username)
		// IP和设备维度的失败计数（使用risk storage）
		_ = riskStorage.ClearLoginFailCount(ctx, cache, "ip", rc.IP)
		if rc.DeviceFP != "" {
			_ = riskStorage.ClearLoginFailCount(ctx, cache, "device", rc.DeviceFP)
		}
	}

	// 记录风控事件
	event := &models.RiskEvent{
		UserId:    int64(userID),
		EventType: baseConstant.RiskEventLoginSuccess,
		Detail:    fmt.Sprintf("login success from %s", rc.IP),
		Score:     0,
		CreatedAt: timePtr(time.Now()),
	}
	return riskRepo.CreateRiskEvent(r.Orm, event)
}

// RecordLoginFail 记录登录失败并执行风控动作
func (r *RiskCheck) RecordLoginFail(ctx context.Context, userID int, username string, rc *riskDto.RiskContext) error {
	cache := r.getCache()
	if cache == nil {
		return nil
	}

	// 1. 增加失败计数（三个维度）
	_ = authStorage.IncrLoginFailCount(ctx, cache, username, 900) // 用户维度：15分钟窗口
	_ = riskStorage.IncrLoginFailCount(ctx, cache, "ip", rc.IP, 900)
	if rc.DeviceFP != "" {
		_ = riskStorage.IncrLoginFailCount(ctx, cache, "device", rc.DeviceFP, 900)
	}

	// 2. 获取失败次数
	failCount, _ := authStorage.GetLoginFailCount(ctx, cache, username)

	// 3. 执行锁定策略
	if failCount == 3 {
		_ = authStorage.LockLogin(ctx, cache, username, 300) // 锁5分钟
		r.Log.Warnf("用户 %s 登录失败3次，锁定5分钟", username)
	} else if failCount == 4 {
		_ = authStorage.LockLogin(ctx, cache, username, 1800) // 锁30分钟
		r.Log.Warnf("用户 %s 登录失败4次，锁定30分钟", username)
	} else if failCount >= 5 {
		// 永久封禁（写入数据库）
		if userID > 0 {
			_ = r.banUser(userID, "连续登录失败5次")
		}
	}

	// 4. 记录风控事件
	if userID > 0 {
		score := int64(10) // 失败基础分
		if failCount >= 3 {
			score = 50
		}
		event := &models.RiskEvent{
			UserId:    int64(userID),
			EventType: baseConstant.RiskEventLoginFail,
			Detail:    fmt.Sprintf("login fail %d times from %s", failCount, rc.IP),
			Score:     score,
			CreatedAt: timePtr(time.Now()),
		}
		_ = riskRepo.CreateRiskEvent(r.Orm, event)

		// 累加风险分
		_ = riskRepo.IncrementRiskScore(r.Orm, uint64(userID), score)
	}

	return nil
}

// checkBlacklist 黑名单检查
func (r *RiskCheck) checkBlacklist(ctx context.Context, rc *riskDto.RiskContext) (bool, error) {
	// IP 黑名单
	item, err := riskRepo.GetActiveBlacklist(r.Orm, "ip", rc.IP)
	if err == nil && item != nil && item.ID > 0 {
		return true, nil
	}

	// 设备黑名单
	if rc.DeviceFP != "" {
		item, err = riskRepo.GetActiveBlacklist(r.Orm, "device", rc.DeviceFP)
		if err == nil && item != nil && item.ID > 0 {
			return true, nil
		}
	}

	return false, nil
}

// banUser 封禁用户
func (r *RiskCheck) banUser(userID int, reason string) error {
	riskUser := &models.RiskUser{
		UserId:    int64(userID),
		RiskLevel: 3, // 封禁
		RiskScore: 200,
		Reason:    reason,
		UpdatedAt: timePtr(time.Now()),
	}
	return riskRepo.UpsertRiskUser(r.Orm, riskUser)
}

// InitRiskUser 初始化风控用户
func (r *RiskCheck) InitRiskUser(userID int) error {
	// 检查是否已存在
	_, err := riskRepo.GetRiskUser(r.Orm, uint64(userID))
	if err == nil {
		return nil // 已存在
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	// 创建
	riskUser := &models.RiskUser{
		UserId:    int64(userID),
		RiskLevel: 0, // 正常
		RiskScore: 0,
		Reason:    "",
		UpdatedAt: timePtr(time.Now()),
	}
	return riskRepo.UpsertRiskUser(r.Orm, riskUser)
}

// BindDevice 绑定设备
func (r *RiskCheck) BindDevice(userID int, deviceFP string) error {
	if deviceFP == "" {
		return nil
	}
	device := &models.RiskDevice{
		DeviceFp:  deviceFP,
		UserId:    int64(userID),
		CreatedAt: timePtr(time.Now()),
	}
	return riskRepo.CreateRiskDevice(r.Orm, device)
}

func timePtr(t time.Time) *time.Time {
	return &t
}
