package storage

import (
	"challenge/app/risk/models"
	baseConstant "challenge/config/base/constant"
	"challenge/core/utils/storage"
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

// 推荐的过期时间常量（秒）
const (
	DefaultStrategyTTL  = 300  // 策略缓存：5分钟
	DefaultBlacklistTTL = 600  // 黑名单缓存：10分钟
	DefaultRiskScoreTTL = 3600 // 风险分数缓存：1小时
)

// 缓存Key前缀定义
const (
	// 策略缓存前缀（使用 baseConstant.RiskStrategyCachePrefix）
	// 格式：challenge:risk:strategy:{scene}

	// 黑名单缓存前缀（使用 baseConstant.RiskBlacklistPrefix）
	// 格式：challenge:risk:blacklist:{type}:{value}

	// 风险分数缓存前缀
	RiskScorePrefix = "challenge:risk:score" // 风险分数缓存，格式：challenge:risk:score:{userID}
)

// CacheStrategies 缓存策略集合
func CacheStrategies(ctx context.Context, cache storage.AdapterCache, scene string, items []models.RiskStrategyCache) error {
	if cache == nil {
		return nil
	}
	b, err := json.Marshal(items)
	if err != nil {
		return err
	}

	return cache.Set(baseConstant.RiskStrategyCachePrefix, scene, string(b), DefaultStrategyTTL)
}

// GetStrategies 从缓存取策略集合
func GetStrategies(ctx context.Context, cache storage.AdapterCache, scene string) ([]models.RiskStrategyCache, error) {
	var out []models.RiskStrategyCache
	if cache == nil {
		return out, nil
	}

	val, err := cache.Get(baseConstant.RiskStrategyCachePrefix, scene)
	if err != nil || val == "" {
		return out, err
	}
	err = json.Unmarshal([]byte(val), &out)
	return out, err
}

// CacheBlacklistFlag 缓存黑名单命中结果
func CacheBlacklistFlag(ctx context.Context, cache storage.AdapterCache, typ, value string, blocked bool, ttl time.Duration) error {
	if cache == nil {
		return nil
	}
	key := fmt.Sprintf("%s:%s", typ, value)
	flag := "0"
	if blocked {
		flag = "1"
	}
	return cache.Set(baseConstant.RiskBlacklistPrefix, key, flag, int(ttl.Seconds()))
}

// GetBlacklistFlag 读取黑名单命中缓存
func GetBlacklistFlag(ctx context.Context, cache storage.AdapterCache, typ, value string) (bool, bool) {
	if cache == nil {
		return false, false
	}
	key := fmt.Sprintf("%s:%s", typ, value)
	val, err := cache.Get(baseConstant.RiskBlacklistPrefix, key)
	if err != nil || val == "" {
		return false, false
	}
	return val == "1", true
}

// ================ 风险分数缓存 ================

// CacheRiskScore 缓存用户风险分数
func CacheRiskScore(ctx context.Context, cache storage.AdapterCache, userID uint64, score int64, ttl time.Duration) error {
	if cache == nil {
		return nil
	}

	return cache.Set(RiskScorePrefix, strconv.FormatUint(userID, 10), fmt.Sprintf("%d", score), int(ttl.Seconds()))
}

// GetRiskScore 获取缓存的风险分数
func GetRiskScore(ctx context.Context, cache storage.AdapterCache, userID uint64) (int64, bool) {
	if cache == nil {
		return 0, false
	}
	val, err := cache.Get(RiskScorePrefix, strconv.FormatUint(userID, 10))
	if err != nil || val == "" {
		return 0, false
	}
	var score int64
	fmt.Sscanf(val, "%d", &score)
	return score, true
}

// ================ 清理函数 ================

// ClearStrategyCache 清除策略缓存
func ClearStrategyCache(ctx context.Context, cache storage.AdapterCache, scene string) error {
	if cache == nil {
		return nil
	}

	return cache.Del(baseConstant.RiskStrategyCachePrefix, scene)
}

// ClearBlacklistCache 清除黑名单缓存
func ClearBlacklistCache(ctx context.Context, cache storage.AdapterCache, typ, value string) error {
	if cache == nil {
		return nil
	}
	key := fmt.Sprintf("%s:%s", typ, value)
	return cache.Del(baseConstant.RiskBlacklistPrefix, key)
}
