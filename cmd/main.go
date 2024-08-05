package main

import (
	"gitlab.ezrpro.in/godemo/global"
)

func main() {

	// 初始化应用
	app := global.NewApplication("demo", "1.0.0")
	// 初始化Http服务
	server := InitServer()
	// app.Run(server.Start)
	app.Run(func(a *global.Application) {
		server.Start()
		a.Wait()
		server.Stop()
	})

	// app.Wait()

	// server.Stop()
}
