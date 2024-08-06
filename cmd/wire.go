//go:build wireinject
// +build wireinject

// go install github.com/google/wire/cmd/wire@latest
package main

import (
	"github.com/google/wire"
	"gitlab.ezrpro.in/godemo/global"
	"gitlab.ezrpro.in/godemo/internal/api"
	"gitlab.ezrpro.in/godemo/internal/api/handlers"
	"gitlab.ezrpro.in/godemo/internal/repo"
	"gitlab.ezrpro.in/godemo/internal/service"
)

// 开发过程 router->handler->service->repository
var ProviderRoutersSet = wire.NewSet(
	api.NewRouters,
	api.NewTestRouter,
	api.NewProjectRouter,
)
var ProviderHandlersSet = wire.NewSet(
	handlers.NewTestHandler,
	handlers.NewProjectHandler,
)
var ProviderServicesSet = wire.NewSet(
	service.NewProjectService,
)
var ProviderReposSet = wire.NewSet(
	repo.NewMysqlProjectRepo,
)

func InitServer() *global.Server {

	wire.Build(
		global.NewServer,
		global.NewGinEngine,
		global.InitDB,
		ProviderRoutersSet,
		ProviderHandlersSet,
		ProviderServicesSet,
		ProviderReposSet,
	)
	return &global.Server{}
}
