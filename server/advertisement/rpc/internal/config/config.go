package config

import (
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Mysql struct {
		DataSource string
		MaxOpen    int
		MaxIdle    int
	}
	Cache   cache.CacheConf
	Product zrpc.RpcClientConf
}
