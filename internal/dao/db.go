package dao

import (
	"context"
	"time"

	"github.com/bilibili/kratos/pkg/database/sql"
	"github.com/bilibili/kratos/pkg/log"
	"github.com/greyireland/shorturl/internal/conf"
	"github.com/greyireland/shorturl/internal/model"
)

var (
	AddURL      = "INSERT INTO `short_url` (incr,raw,code,ctime) values (?,?,?,?)"
	QueryURL    = "SELECT raw,code,ctime from `short_url` where code = ?"
	QueryIncrID = "SELECT incr from `short_url` where ctime < ? order by ctime desc limit 1"
)

func NewDB() (db *sql.DB, cf func(), err error) {
	db = sql.NewMySQL(&conf.Cfg.DB)
	cf = func() { db.Close() }
	return
}

func (d *dao) AddRawURL(ctx context.Context, u *model.URL) (err error) {
	res, err := d.db.Exec(ctx, AddURL, u.Incr, u.Raw, u.Code, time.Now())
	if err != nil {
		log.Errorv(ctx, log.KV("err", err.Error()))
		return err
	}
	u.ID, _ = res.LastInsertId()
	return
}

func (d *dao) GetRawURL(ctx context.Context, code string) (res *model.URL, err error) {
	row := d.db.QueryRow(ctx, QueryURL, code)
	var r model.URL
	err = row.Scan(&r.Raw, &r.Code, &r.CTime)
	if err != nil {
		return nil, err
	}
	return &r, nil
}
func (d *dao) GetDBIncrID(ctx context.Context) (id int64, err error) {
	row := d.db.QueryRow(ctx, QueryIncrID, time.Now())
	if err != nil {
		return
	}
	err = row.Scan(&id)
	return
}
