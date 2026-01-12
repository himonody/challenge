package service

import (
	"challenge/app/risk/models"
	"challenge/app/risk/repo"
	riskStorage "challenge/app/risk/storage"
	baseConstant "challenge/config/base/constant"
	baseLang "challenge/config/base/lang"
	"challenge/core/dto/service"
	"challenge/core/lang"

	"context"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Risk struct {
	service.Service
}

// NewRiskService 实例化风控服务
func NewRiskService(s *service.Service) *Risk {
	var srv = new(Risk)
	srv.Orm = s.Orm
	srv.Log = s.Log
	srv.Run = s.Run
	srv.Lang = s.Lang
	return srv
}

// LoadStrategies 获取场景策略（带缓存，下发策略缓存表 & redis）
func (r *Risk) LoadStrategies(ctx context.Context, scene string) ([]models.RiskStrategyCache, int, error) {
	cache := r.Run.GetCacheAdapter()
	if cache == nil {
		r.Log.Warnf("risk.LoadStrategies: cache not available")
		// 无缓存时直接查表
		return r.loadStrategiesFromDB(ctx, scene)
	}

	// 1) redis 缓存
	if list, err := riskStorage.GetStrategies(ctx, cache, scene); err == nil && len(list) > 0 {
		return list, baseLang.SuccessCode, nil
	}

	// 2) 表缓存
	list, err := repo.ListStrategyCacheByScene(r.Orm, scene)
	if err != nil {
		return nil, baseLang.DataQueryLogCode, lang.MsgLogErrf(r.Log, r.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	if len(list) > 0 {
		_ = riskStorage.CacheStrategies(ctx, cache, scene, list)
		return list, baseLang.SuccessCode, nil
	}

	// 3) 主表并回填缓存表
	return r.loadStrategiesFromDB(ctx, scene)
}

// loadStrategiesFromDB 从主表加载策略并回填缓存
func (r *Risk) loadStrategiesFromDB(ctx context.Context, scene string) ([]models.RiskStrategyCache, int, error) {
	strategies, err := repo.ListStrategiesByScene(r.Orm, scene)
	if err != nil {
		return nil, baseLang.DataQueryLogCode, lang.MsgLogErrf(r.Log, r.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	if len(strategies) == 0 {
		return nil, baseLang.RiskStrategyNotFoundCode, lang.MsgErr(baseLang.RiskStrategyNotFoundCode, r.Lang)
	}

	caches := make([]models.RiskStrategyCache, 0, len(strategies))
	for _, item := range strategies {
		caches = append(caches, models.RiskStrategyCache{
			Scene:         item.Scene,
			IdentityType:  item.IdentityType,
			RuleCode:      item.RuleCode,
			WindowSeconds: item.WindowSeconds,
			Threshold:     item.Threshold,
			Action:        item.Action,
			ActionValue:   item.ActionValue,
		})
	}

	if err = repo.UpsertStrategyCache(r.Orm, caches); err != nil {
		return nil, baseLang.DataUpdateLogCode, lang.MsgLogErrf(r.Log, r.Lang, baseLang.DataUpdateCode, baseLang.DataUpdateLogCode, err)
	}

	cache := r.Run.GetCacheAdapter()
	if cache != nil {
		_ = riskStorage.CacheStrategies(ctx, cache, scene, caches)
	}

	return caches, baseLang.SuccessCode, nil
}

// ListActions 读取动作列表
func (r *Risk) ListActions() (map[string]models.RiskAction, int, error) {
	list, err := repo.ListActions(r.Orm)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, baseLang.DataQueryLogCode, lang.MsgLogErrf(r.Log, r.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	res := make(map[string]models.RiskAction, len(list))
	for _, item := range list {
		res[item.Code] = item
	}
	return res, baseLang.SuccessCode, nil
}

// CheckBlacklist 判断是否命中黑名单，缓存结果
func (r *Risk) CheckBlacklist(ctx context.Context, typ, value string) (bool, int, error) {
	cache := r.Run.GetCacheAdapter()
	if cache != nil {
		if hit, ok := riskStorage.GetBlacklistFlag(ctx, cache, typ, value); ok {
			if hit {
				return true, baseLang.RiskBlacklistHitCode, lang.MsgErr(baseLang.RiskBlacklistHitCode, r.Lang)
			}
			return false, baseLang.SuccessCode, nil
		}
	}

	item, err := repo.GetActiveBlacklist(r.Orm, typ, value)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return false, baseLang.DataQueryLogCode, lang.MsgLogErrf(r.Log, r.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}

	hit := err == nil && item != nil && item.ID > 0
	if cache != nil {
		_ = riskStorage.CacheBlacklistFlag(ctx, cache, typ, value, hit, 10*time.Minute)
	}

	if hit {
		return true, baseLang.RiskBlacklistHitCode, lang.MsgErr(baseLang.RiskBlacklistHitCode, r.Lang)
	}
	return false, baseLang.SuccessCode, nil
}

// ================ 用户风险管理 ================

// GetUserRiskLevel 获取用户风险等级
func (r *Risk) GetUserRiskLevel(ctx context.Context, userID uint64) (int64, int64, error) {
	// 先从缓存获取
	cache := r.Run.GetCacheAdapter()
	if cache != nil {
		if score, ok := riskStorage.GetRiskScore(ctx, cache, userID); ok {
			level := r.calculateRiskLevel(score)
			return level, score, nil
		}
	}

	// 从数据库获取
	riskUser, err := repo.GetRiskUser(r.Orm, userID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, 0, err
	}

	if riskUser == nil || riskUser.UserId == 0 {
		// 用户不存在风控记录，返回正常
		return 0, 0, nil // 0 = 正常
	}

	// 缓存风险分数
	if cache != nil {
		_ = riskStorage.CacheRiskScore(ctx, cache, userID, riskUser.RiskScore, time.Hour)
	}

	return riskUser.RiskLevel, riskUser.RiskScore, nil
}

// calculateRiskLevel 根据分数计算风险等级（返回数字：0正常 1观察 2限制 3封禁）
func (r *Risk) calculateRiskLevel(score int64) int64 {
	switch {
	case score >= 80:
		return 3 // 封禁
	case score >= 50:
		return 2 // 限制
	case score >= 20:
		return 1 // 观察
	default:
		return 0 // 正常
	}
}

// UpdateUserRiskScore 更新用户风险分数
func (r *Risk) UpdateUserRiskScore(ctx context.Context, userID uint64, deltaScore int64, reason string) error {
	// 获取当前分数
	_, currentScore, err := r.GetUserRiskLevel(ctx, userID)
	if err != nil {
		return err
	}

	newScore := currentScore + deltaScore
	if newScore < 0 {
		newScore = 0
	}

	// 计算新等级
	newLevel := r.calculateRiskLevel(newScore)

	// 更新数据库
	riskUser := &models.RiskUser{
		UserId:    int64(userID),
		RiskLevel: newLevel,
		RiskScore: newScore,
		Reason:    reason,
		UpdatedAt: timePtr(time.Now()),
	}
	err = repo.UpsertRiskUser(r.Orm, riskUser)
	if err != nil {
		return err
	}

	// 清除缓存
	cache := r.Run.GetCacheAdapter()
	if cache != nil {
		_ = riskStorage.CacheRiskScore(ctx, cache, userID, newScore, time.Hour)
	}

	// 记录风控事件
	event := &models.RiskEvent{
		UserId:    int64(userID),
		EventType: baseConstant.RiskEventScoreChange,
		Detail:    fmt.Sprintf("分数变化: %+d, 原因: %s, 新分数: %d, 新等级: %d", deltaScore, reason, newScore, newLevel),
		Score:     deltaScore,
		CreatedAt: timePtr(time.Now()),
	}
	_ = repo.CreateRiskEvent(r.Orm, event)

	r.Log.Infof("用户 %d 风险分数更新: %d -> %d, 等级: %d", userID, currentScore, newScore, newLevel)

	return nil
}

// ================ 黑名单管理 ================

// AddToBlacklist 添加到黑名单
func (r *Risk) AddToBlacklist(ctx context.Context, typ, value, reason string) error {
	now := time.Now()
	blacklist := &models.RiskBlacklist{
		Type:      typ,
		Value:     value,
		RiskLevel: 3, // 最高风险
		Reason:    reason,
		Status:    baseConstant.GeneralStatusOk,
		CreatedAt: &now,
	}

	err := r.Orm.Create(blacklist).Error
	if err != nil {
		return err
	}

	// 清除缓存，强制重新查询
	cache := r.Run.GetCacheAdapter()
	if cache != nil {
		_ = riskStorage.ClearBlacklistCache(ctx, cache, typ, value)
	}

	r.Log.Infof("添加黑名单: 类型=%s, 值=%s, 原因=%s", typ, value, reason)
	return nil
}

// RemoveFromBlacklist 从黑名单移除
func (r *Risk) RemoveFromBlacklist(ctx context.Context, typ, value string) error {
	err := r.Orm.Model(&models.RiskBlacklist{}).
		Where("type = ? AND value = ?", typ, value).
		Update("status", baseConstant.GeneralStatusBlock).
		Error

	if err != nil {
		return err
	}

	// 清除缓存
	cache := r.Run.GetCacheAdapter()
	if cache != nil {
		_ = riskStorage.ClearBlacklistCache(ctx, cache, typ, value)
	}

	r.Log.Infof("移除黑名单: 类型=%s, 值=%s", typ, value)
	return nil
}

// ================ 策略管理 ================

// RefreshStrategyCache 刷新策略缓存
func (r *Risk) RefreshStrategyCache(ctx context.Context, scene string) error {
	cache := r.Run.GetCacheAdapter()
	if cache == nil {
		return errors.New("缓存不可用")
	}

	// 清除旧缓存
	_ = riskStorage.ClearStrategyCache(ctx, cache, scene)

	// 重新加载
	_, code, err := r.LoadStrategies(ctx, scene)
	if code != baseLang.SuccessCode {
		return err
	}

	r.Log.Infof("刷新策略缓存成功: scene=%s", scene)
	return nil
}

// ================ 辅助函数 ================

func timePtr(t time.Time) *time.Time {
	return &t
}
