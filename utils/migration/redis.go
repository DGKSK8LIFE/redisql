package migration

import "github.com/go-redis/redis/v8"

func OpenRedis(rdb *redis.Client, redisAddress, redisPassword string) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: redisPassword,
		DB:       0,
	})
}
