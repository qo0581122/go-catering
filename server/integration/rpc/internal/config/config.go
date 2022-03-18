package config

import (
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Mysql struct {
		DataSource string
		MaxIdle    int
		MaxOpen    int
	}
	Cache   cache.CacheConf
	Product zrpc.RpcClientConf
}

func (c *Config) GetDataSource() string {
	return c.Mysql.DataSource
}

func (c *Config) GetMaxIdle() int {
	return c.Mysql.MaxIdle
}

func (c *Config) GetMaxOpen() int {
	return c.Mysql.MaxOpen
}
