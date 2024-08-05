//go:build wireinject
// +build wireinject

// go install github.com/google/wire/cmd/wire@latest
package main

import (
	"github.com/google/wire"
	"gitlab.ezrpro.in/godemo/global"
	"gitlab.ezrpro.in/godemo/internal/api"
	"gitlab.ezrpro.in/godemo/internal/api/handlers"
)

// 开发过程 router->handler->service->repository
var ProviderRoutersSet = wire.NewSet(
	api.NewRouters,
	api.NewPingRouter,
)
var ProviderHandlersSet = wire.NewSet(
	handlers.NewPingHandler,
)
var ProviderServicesSet = wire.NewSet()
var ProviderReposSet = wire.NewSet()

func InitServer() *global.Server {

	wire.Build(
		global.NewServer,
		global.NewGinEngine,
		// api.NewPingRouter,
		ProviderRoutersSet,
		ProviderHandlersSet,
		// ProviderServicesSet,
		// ProviderReposSet,
	)
	return &global.Server{}
}
