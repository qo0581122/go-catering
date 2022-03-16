package global

import (
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"

	"catering/config"

	"gorm.io/gorm"
)

var (
	DB                  *gorm.DB
	Config              config.Config
	Log                 *zap.Logger
	REDIS               *redis.Client
	Concurrency_Control = &singleflight.Group{}
)
