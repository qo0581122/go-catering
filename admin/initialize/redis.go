package initialize

import (
	"catering/global"
	"context"

	"github.com/go-redis/redis/v8"
)

func InitRedis() error {
	redisCfg := global.Config.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return err
	} else {
		global.REDIS = client
	}
	return nil
}
