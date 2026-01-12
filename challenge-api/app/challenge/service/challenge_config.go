package service

import (
	"challenge/app/challenge/models"
	"challenge/app/challenge/repo"
	"challenge/app/challenge/storage"
	"challenge/core/dto/service"
)

type ChallengeConfig struct {
	service.Service
}

func NewChallengeConfig(s *service.Service) *ChallengeConfig {
	return &ChallengeConfig{
		Service: *s,
	}
}
func (c *ChallengeConfig) GetChallengeConfig() (*models.AppChallengeConfig, error) {
	cache := c.Run.GetCacheAdapter()
	configCache, err := storage.GetChallengeConfigCache(c.C.Request.Context(), cache)
	if err == nil && configCache != nil {
		return configCache, nil
	}
	config, err := repo.GetChallengeConfig(c.Orm)
	if err != nil {
		return nil, err
	}
	if config != nil {
		go storage.SetChallengeConfigCache(c.C.Request.Context(), cache, config)
	}

	return config, nil
}
