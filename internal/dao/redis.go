package dao

import (
	"context"
	"fmt"
	"time"

	"github.com/bilibili/kratos/pkg/cache/redis"
	"github.com/bilibili/kratos/pkg/log"
	"github.com/greyireland/shorturl/internal/conf"
)

var (
	URL    = "url_%s"
	IncrID = "incrID"
)

func NewRedis() (r *redis.Redis, cf func(), err error) {
	r = redis.NewRedis(&conf.Cfg.Redis)
	cf = func() { r.Close() }
	return
}

func (d *dao) PingRedis(ctx context.Context) (err error) {
	if _, err = d.redis.Do(ctx, "SET", "ping", "pong"); err != nil {
		log.Error("conn.Set(PING) error(%v)", err)
	}
	return
}
func (d *dao) GetCacheURL(ctx context.Context, code string) (string, error) {
	return redis.String(d.redis.Do(ctx, "GET", key(URL, code)))

}
func (d *dao) AddCacheURL(ctx context.Context, code, raw string) (string, error) {
	return redis.String(d.redis.Do(ctx,
		"SETEX",
		key(URL, code),
		int(time.Duration(d.cfg.App.Expire).Seconds()), raw))
}
func (d *dao) Incr(ctx context.Context) (id int64, err error) {
	return redis.Int64(d.redis.Do(ctx, "INCR", IncrID))
}
func key(format string, args ...interface{}) string {
	return fmt.Sprintf(format, args...)
}
