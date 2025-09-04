package initialize

import (
	"os"
	"server/global"

	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

func ConnectRedis() redis.Client {
	redisConfig := global.Config.Redis

	client := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Address,
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
	})

	_, err := client.Ping().Result()

	if err != nil {
		global.Log.Error("连接Redis异常", zap.Error(err))
		os.Exit(0)
	}

	return *client
}
