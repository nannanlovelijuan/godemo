//go:build wireinject
// +build wireinject

// go install github.com/google/wire/cmd/wire@latest
package main

import (
	"github.com/google/wire"
	"gitlab.ezrpro.in/godemo/global"
)

// 开发过程 router->handler->service->repository
var ProviderRoutersSet = wire.NewSet()
var ProviderHandlersSet = wire.NewSet()
var ProviderServicesSet = wire.NewSet()
var ProviderReposSet = wire.NewSet()

func InitServer() *global.Server {

	wire.Build(
		global.NewServer,
		global.NewGinEngine,
		// ProviderRoutersSet,
		// ProviderHandlersSet,
		// ProviderServicesSet,
		// ProviderReposSet,
	)
	return &global.Server{}
}
