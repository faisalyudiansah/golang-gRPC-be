package redis

import (
	"fmt"

	"server/pkg/config"
	"server/pkg/logger"

	"github.com/RediSearch/redisearch-go/redisearch"
)

func InitRedisSearch(cfg *config.RedisConfig) *redisearch.Client {
	rds := redisearch.NewClient(fmt.Sprintf("%v:%v", cfg.Host, cfg.Port), "clean-arch-v2Index")

	if err := rds.Drop(); err != nil {
		logger.Log.Info("no existing index...")
	}
	logger.Log.Info("redisearch is ready...")

	return rds
}
