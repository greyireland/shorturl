package http

import (
	"net/http"
	"strings"

	"github.com/bilibili/kratos/pkg/ecode"
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
	pb "github.com/greyireland/shorturl/api"
	"github.com/greyireland/shorturl/internal/conf"
)

var svc pb.ShortURLBMServer

// New new a bm server.
func New(s pb.ShortURLBMServer) (engine *bm.Engine, err error) {

	svc = s
	engine = bm.DefaultServer(&conf.Cfg.HTTP)
	engine.Use(bm.HandlerFunc(CORS))
	pb.RegisterShortURLBMServer(engine, s)
	initRouter(engine)
	err = engine.Start()
	return
}

func CORS(c *bm.Context) {
	c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Add("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE,OPTIONS")
	c.Writer.Header().Add("Access-Control-Allow-Headers", "*")
	if c.Request.Method == http.MethodOptions {
		c.AbortWithStatus(204)
		return
	}
	c.Next()
}

func initRouter(e *bm.Engine) {
	e.Any("/", dispath)
}

// example for http request handler.
func dispath(c *bm.Context) {
	code := strings.Trim(c.Request.URL.Path, "/")
	if len(code) == 0 {
		// redirect to index.html
		c.JSON(nil, ecode.RequestErr)
	}

	res, err := svc.GetRawURL(c.Context, &pb.GetRawURLReq{Code: code})
	if err != nil {
		c.JSON(nil, ecode.ServerErr)
		return
	}
	c.Redirect(302, res.RawUrl)
}
