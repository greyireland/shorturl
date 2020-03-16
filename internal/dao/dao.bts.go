package dao

import (
	"context"

	"github.com/greyireland/shorturl/internal/model"
)

func (d *dao) AddURL(ctx context.Context, u *model.URL) (err error) {
	err = d.addCode(u.Code, u.Raw)
	if err != nil {
		return
	}
	_, err = d.AddCacheURL(ctx, u.Code, u.Raw)
	if err != nil {
		return
	}
	return d.AddRawURL(ctx, u)
}
func (d *dao) GetURL(ctx context.Context, code string) (u *model.URL, err error) {
	raw, _ := d.GetCacheURL(ctx, code)
	if len(raw) != 0 {
		u = &model.URL{Raw: raw}
		return
	}

	u, err = d.GetRawURL(ctx, code)
	if err != nil {
		return
	}
	d.cache.Do(ctx, func(ctx context.Context) {
		d.AddCacheURL(ctx, u.Code, u.Raw)
	})
	return
}
func (d *dao) GetCode(ctx context.Context, raw string) (res *model.URL, err error) {
	var code string
	code, err = d.getCode(raw)
	if err != nil || len(code) == 0 {
		return
	}
	return &model.URL{Code: code}, nil
}
func (d *dao) GetIncrID(ctx context.Context) (int64, error) {
	id, err := d.Incr(ctx)
	if err == nil && id != 0 {
		return id, nil
	}
	return d.GetDBIncrID(ctx)
}
