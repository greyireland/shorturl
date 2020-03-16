package grpc

import (
	pb "github.com/greyireland/shorturl/api"
	"github.com/greyireland/shorturl/internal/conf"

	"github.com/bilibili/kratos/pkg/net/rpc/warden"
)

// New new a grpc server.
func New(svc pb.ShortURLBMServer) (ws *warden.Server, err error) {
	ws = warden.NewServer(&conf.Cfg.GRPC)
	pb.RegisterShortURLServer(ws.Server(), svc)
	ws, err = ws.Start()
	return
}
