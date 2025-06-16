package stores

import (
	"github.com/redis/go-redis/v9"
	"github.com/vaibhavgupta04/marketdatafeeder/config"
)

func NewRedisClient() *redis.Client{
	redisCfg := config.LoadRedisConfig()

	rdb := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Host,
		Username: redisCfg.Username,
		Password: redisCfg.Password,
		DB:       redisCfg.DB,
	})
	return rdb
}



