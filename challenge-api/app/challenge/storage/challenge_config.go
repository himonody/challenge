package storage

import (
	"challenge/app/challenge/models"
	"challenge/core/utils/storage"
	"context"
	"encoding/json"
)

const (
	ChallengeConfigCachePrefix = "challenge:challenge:config"
	DefaultStrategyTTL         = 24 * 3600
)

// SetChallengeConfigCache 缓存策略集合
func SetChallengeConfigCache(ctx context.Context, cache storage.AdapterCache, items []models.AppChallengeConfig) error {
	if cache == nil {
		return nil
	}
	b, err := json.Marshal(items)
	if err != nil {
		return err
	}

	return cache.Set(ChallengeConfigCachePrefix, "", string(b), DefaultStrategyTTL)
}

// GetChallengeConfigCache 从缓存取策略集合
func GetChallengeConfigCache(ctx context.Context, cache storage.AdapterCache) ([]models.AppChallengeConfig, error) {
	var out []models.AppChallengeConfig
	if cache == nil {
		return out, nil
	}

	val, err := cache.Get(ChallengeConfigCachePrefix, "")
	if err != nil || val == "" {
		return out, err
	}
	err = json.Unmarshal([]byte(val), &out)
	return out, err
}
func DelChallengeConfigCache(ctx context.Context, cache storage.AdapterCache) error {
	return cache.Del(ChallengeConfigCachePrefix, "")
}
