// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"github.com/greyireland/shorturl/internal/dao"
	"github.com/greyireland/shorturl/internal/server/grpc"
	"github.com/greyireland/shorturl/internal/server/http"
	"github.com/greyireland/shorturl/internal/server/redirect"
	"github.com/greyireland/shorturl/internal/service"

	"github.com/google/wire"
)

//go:generate kratos t wire
func InitApp() (*App, func(), error) {
	panic(wire.Build(dao.Provider, service.Provider, http.New, grpc.New, redirect.New, NewApp))
}
