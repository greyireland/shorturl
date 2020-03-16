package conf

import (
	"fmt"

	"github.com/bilibili/kratos/pkg/cache/redis"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/database/sql"
	"github.com/bilibili/kratos/pkg/net/http/blademaster"
	"github.com/bilibili/kratos/pkg/net/rpc/warden"
	"github.com/bilibili/kratos/pkg/time"
)

var (
	Cfg Config
)

type Config struct {
	App   App                      `toml:"app"`
	HTTP  blademaster.ServerConfig `toml:"http"`
	GRPC  warden.ServerConfig      `toml:"grpc"`
	DB    sql.Config               `toml:"db"`
	Redis redis.Config             `toml:"redis"`
}
type App struct {
	Domain       string        `toml:"domain"`
	CacheSize    int           `toml:"cache_size"`
	Expire       time.Duration `toml:"expire"`
	RedirectPort string        `toml:"redirect_port"`
}

func Init() (err error) {
	var (
		ct paladin.TOML
	)
	if err = paladin.Get("config.toml").Unmarshal(&ct); err != nil {
		return
	}
	if err = ct.Get("app").UnmarshalTOML(&Cfg.App); err != nil {
		return
	}
	if err = ct.Get("db").UnmarshalTOML(&Cfg.DB); err != nil {
		return
	}
	if err = ct.Get("redis").UnmarshalTOML(&Cfg.Redis); err != nil {
		return
	}
	if err = ct.Get("http").UnmarshalTOML(&Cfg.HTTP); err != nil {
		return
	}
	if err = ct.Get("grpc").UnmarshalTOML(&Cfg.GRPC); err != nil {
		return
	}
	fmt.Print("conf:", Cfg)
	return
}
