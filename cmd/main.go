package main

import "gitlab.ezrpro.in/godemo/global"

func main() {

	// 初始化应用
	app := global.NewApplication("EZR.Arch.FlinkManager.ApiHost", "1.0.0")
	// 初始化Http服务
	server := InitServer(app)

	app.Run(func(a *global.Application) {
		server.Start()
		a.Wait()
		server.Stop()
	})

}
