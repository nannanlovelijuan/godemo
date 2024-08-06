package global

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gitlab.ezrpro.in/godemo/internal/api"
)

// server 具有start 和 stop 方法
// server 使用gin做为处理器
type Server struct {
	engine  *gin.Engine
	server  *http.Server
	routers *api.Routers
}

func NewServer(engine *gin.Engine, routers *api.Routers) *Server {
	return &Server{engine: engine, server: NewHttpServer(8080), routers: routers}
}

func (s *Server) Start() {

	//启动新线程启动服务
	go func() {
		//加载所有路由
		api.RegisterRouters(s.engine, s.routers)

		s.server.Handler = s.engine
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("server start error:%v\n", err)
			os.Exit(1)
		}
	}()
}

func (s *Server) Stop() {
	fmt.Println("server shutdown...")
	if err := s.server.Shutdown(context.Background()); err != nil {
		fmt.Printf("server shutdown error:%v\n", err)
	}
}
