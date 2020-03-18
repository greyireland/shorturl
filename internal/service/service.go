package service

import (
	"context"
	"fmt"

	pb "github.com/greyireland/shorturl/api"
	"github.com/greyireland/shorturl/internal/conf"
	"github.com/greyireland/shorturl/internal/dao"
	"github.com/greyireland/shorturl/internal/model"

	"github.com/google/wire"
)

var Provider = wire.NewSet(New, wire.Bind(new(pb.ShortURLBMServer), new(*Service)))

// Service service.
type Service struct {
	dao dao.Dao
}

// New new a service and return.
func New(d dao.Dao) (s *Service, cf func(), err error) {
	s = &Service{
		dao: d,
	}
	cf = s.Close
	return
}

func (s *Service) Shorten(ctx context.Context, req *pb.URLReq) (reply *pb.URLResp, err error) {
	var (
		code string
		incr int64
	)
	res, err := s.dao.GetCode(ctx, req.RawUrl)
	if res != nil && err == nil {
		code = res.Code
	} else {
		incr, err = s.dao.GetIncrID(ctx)
		if err != nil {
			incr = 0
		}
		code = DecimalToAny(genID(incr), 64)
		u := &model.URL{Raw: req.RawUrl, Incr: incr, Code: code}
		err = s.dao.AddURL(ctx, u)
		if err != nil {
			return
		}
	}

	content := fmt.Sprintf("%s/%s", conf.Cfg.App.Domain, code)
	reply = &pb.URLResp{Code: content}
	return
}
func genID(incr int64) int64 {
	return incr + 1
}
func (s *Service) GetRawURL(ctx context.Context, req *pb.GetRawURLReq) (res *pb.GetRawURLResp, err error) {
	var r *model.URL
	r, err = s.dao.GetURL(ctx, req.Code)
	if err != nil {
		return
	}
	res = &pb.GetRawURLResp{RawUrl: r.Raw}
	return
}

// Close close the resource.
func (s *Service) Close() {
}
