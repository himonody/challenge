package storage

import (
	"challenge/core/utils/storage"
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

var rateLimitLua = redis.NewScript(`
local current = redis.call("INCR", KEYS[1])
if current == 1 then
    redis.call("EXPIRE", KEYS[1], ARGV[2])
end
if current > tonumber(ARGV[1]) then
    return 0
end
return 1
`)

func RateLimitCheck(c context.Context, cache storage.AdapterCache, key string, limit int, ttl time.Duration) {
	if r, ok := cache.(interface{ GetClient() *redis.Client }); ok && r.GetClient() != nil {
		client := r.GetClient()
		// 在这里使用原生 client，例如：
		client.Get(c, key)
	}
	//	rateLimitLua.Run(c, cache.String(), []string{key}, limit, int(ttl.Seconds()))

}
