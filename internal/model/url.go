package model

import (
	"github.com/bilibili/kratos/pkg/time"
)

type URL struct {
	ID    int64     `json:"id"`
	Incr  int64     `json:"incr"`
	Raw   string    `json:"raw"`
	Code  string    `json:"code"`
	CTime time.Time `json:"ctime"`
	MTime time.Time `json:"mtime"`
}
