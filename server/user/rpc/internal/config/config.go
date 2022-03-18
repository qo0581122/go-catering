package config

import (
	snow "catering/area/pkg/snowflake"

	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Mysql struct {
		DataSource  string
		MaxOpen     int
		MaxIdle     int
		MaxIdleTime int
	}
	Cache     cache.CacheConf
	Snowflake snow.SnowflakeConf
}
