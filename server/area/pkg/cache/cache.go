package cache

import (
	"database/sql"

	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/syncx"
)

var (
	// can't use one SingleFlight per conn, because multiple conns may share the same cache key.
	exclusiveCalls = syncx.NewSingleFlight()
	stats          = cache.NewStat("sqlc")
)

func NewCache(c cache.CacheConf) *cache.Cache {
	cache := cache.New(c, exclusiveCalls, stats, sql.ErrNoRows)
	return &cache
}
