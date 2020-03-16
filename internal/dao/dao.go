package dao

import (
	"context"

	"github.com/bilibili/kratos/pkg/cache/redis"
	"github.com/bilibili/kratos/pkg/database/sql"
	"github.com/bilibili/kratos/pkg/sync/pipeline/fanout"
	"github.com/google/wire"
	"github.com/greyireland/shorturl/internal/conf"
	"github.com/greyireland/shorturl/internal/model"
	"github.com/hashicorp/golang-lru"
)

var Provider = wire.NewSet(New, NewDB, NewRedis)

type Dao interface {
	Close()
	Ping(ctx context.Context) (err error)
	AddURL(ctx context.Context, url *model.URL) error
	GetURL(ctx context.Context, code string) (*model.URL, error)
	GetCode(ctx context.Context, raw string) (*model.URL, error)
	GetIncrID(ctx context.Context) (int64, error)
}

// dao dao.
type dao struct {
	db    *sql.DB
	redis *redis.Redis
	cache *fanout.Fanout
	lru   *lru.Cache
	cfg   conf.Config
}

// New new a dao and return.
func New(r *redis.Redis, db *sql.DB) (d Dao, cf func(), err error) {
	return newDao(r, db)
}

func newDao(r *redis.Redis, db *sql.DB) (d *dao, cf func(), err error) {

	lruCache, err := lru.New(conf.Cfg.App.CacheSize)
	if err != nil {
		return
	}

	d = &dao{
		db:    db,
		redis: r,
		cache: fanout.New("cache"),
		lru:   lruCache,
		cfg:   conf.Cfg,
	}
	cf = d.Close
	return
}

// Close close the resource.
func (d *dao) Close() {
	d.cache.Close()
}

// Ping ping the resource.
func (d *dao) Ping(ctx context.Context) (err error) {
	return nil
}
