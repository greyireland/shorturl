package redirect

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/greyireland/shorturl/api"
	"github.com/greyireland/shorturl/internal/conf"
)

func New(svc api.ShortURLBMServer) (server *http.Server, err error) {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		fmt.Println("path:", path)
		code := strings.Trim(path, "/")
		if len(code) == 0 {
			w.Write(indexHTML)
			w.WriteHeader(http.StatusOK)
			return
		}
		res, err := svc.GetRawURL(context.TODO(), &api.GetRawURLReq{Code: code})
		fmt.Println("raw url:", res)
		if err != nil {
			w.Write(indexHTML)
			w.WriteHeader(http.StatusOK)
			return
		}
		w.Header().Set("Location", res.RawUrl)
		w.WriteHeader(302)
	})

	fmt.Println("redirect serve port:", conf.Cfg.App.RedirectPort)
	go func() {
		http.ListenAndServe("0.0.0.0:"+conf.Cfg.App.RedirectPort, nil)
	}()
	return
}
