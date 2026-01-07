package storage

import (
	"context"
	"fmt"
	"time"

	"challenge/core/utils/storage"

	"github.com/bsm/redislock"
)

// 认证相关的分布式锁前缀
const AuthLockPrefix = "challenge:auth"
const AuthUniqueLockPrefix = "challenge:auth:%s:%s"

// LockAuthAction 获取认证相关分布式锁
// action 示例：login / register / captcha 等
// key     示例：用户名或唯一标识
func LockAuthAction(ctx context.Context, locker storage.AdapterLocker, action, key string, ttlSec int64) (*redislock.Lock, error) {
	if locker == nil {
		return nil, nil
	}
	if ttlSec <= 0 {
		ttlSec = 10
	}
	lockKey := fmt.Sprintf(AuthUniqueLockPrefix, action, key)
	return locker.Lock(lockKey, ttlSec, &redislock.Options{
		RetryStrategy: redislock.LimitRetry(redislock.LinearBackoff(200*time.Millisecond), 5),
	})
}

// WithAuthLock 执行带认证分布式锁的函数
func WithAuthLock(ctx context.Context, locker storage.AdapterLocker, action, key string, ttlSec int64, fn func() error) error {
	if locker == nil {
		return fn()
	}
	lock, err := LockAuthAction(ctx, locker, action, key, ttlSec)
	if err != nil {
		return err
	}
	if lock != nil {
		defer lock.Release(ctx)
	}
	return fn()
}
