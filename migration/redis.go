package migration

import "github.com/go-redis/redis/v8"

// OpenRedis opens a redis connection with a desired address and password
func OpenRedis(redisAddress, redisPassword string) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: redisPassword,
		DB:       0,
	})
	return rdb
}
