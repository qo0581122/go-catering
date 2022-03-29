package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/sessions"
	"go.uber.org/zap"

	"catering/config"

	"gorm.io/gorm"
)

var (
	DB      *gorm.DB
	Config  config.Config
	Log     *zap.Logger
	REDIS   *redis.Client
	SESSION *sessions.CookieStore
)
