package storage

import (
	"challenge/core/utils/storage"
	"context"
	"fmt"
	"time"

	"github.com/bsm/redislock"
)

const ChallengeLockPrefix = "challenge:challenge:checkin"
const ChallengeCheckinUniqueLockPrefix = "challenge:challenge:checkin:%s:%s"

func LockCheckin(ctx context.Context, locker storage.AdapterLocker, action, key string, ttlSec int64) (*redislock.Lock, error) {
	if locker == nil {
		return nil, nil
	}
	if ttlSec <= 0 {
		ttlSec = 10
	}
	lockKey := fmt.Sprintf(ChallengeCheckinUniqueLockPrefix, action, key)
	return locker.Lock(lockKey, ttlSec, &redislock.Options{
		RetryStrategy: redislock.LimitRetry(redislock.LinearBackoff(200*time.Millisecond), 5),
	})
}

// WithCheckinLock 执行带认证分布式锁的函数
func WithCheckinLock(ctx context.Context, locker storage.AdapterLocker, action, key string, ttlSec int64, fn func() error) error {
	if locker == nil {
		return fn()
	}
	lock, err := LockCheckin(ctx, locker, action, key, ttlSec)
	if err != nil {
		return err
	}
	if lock != nil {
		defer lock.Release(ctx)
	}
	return fn()
}
