package storage

import (
	"context"
	"fmt"
	"time"

	"challenge/core/utils/storage"

	"github.com/bsm/redislock"
)

// 分布式锁前缀（保持与全局约定一致）
const (
	UserLockPrefix = "challenge:user"
)

// LockUserAction 获取用户级分布式锁
func LockUserAction(ctx context.Context, locker storage.AdapterLocker, userID uint64, action string, ttlSec int64) (*redislock.Lock, error) {
	if locker == nil {
		return nil, nil
	}
	if ttlSec <= 0 {
		ttlSec = 10 // 默认10秒
	}
	key := fmt.Sprintf("%s:%s:%d", UserLockPrefix, action, userID)
	return locker.Lock(key, ttlSec, &redislock.Options{
		RetryStrategy: redislock.LimitRetry(redislock.LinearBackoff(200*time.Millisecond), 5),
	})
}

// WithUserLock 执行带用户级分布式锁的函数
func WithUserLock(ctx context.Context, locker storage.AdapterLocker, userID uint64, action string, ttlSec int64, fn func() error) error {
	if locker == nil {
		return fn()
	}
	lock, err := LockUserAction(ctx, locker, userID, action, ttlSec)
	if err != nil {
		return err
	}
	if lock != nil {
		defer lock.Release(ctx)
	}
	return fn()
}
