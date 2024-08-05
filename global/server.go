package global

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.ezrpro.in/godemo/internal/api"
)

// server 具有start 和 stop 方法
// server 使用gin做为处理器
type Server struct {
	engine     *gin.Engine
	server     *http.Server
	pingrouter *api.PingRouter
}

func NewServer(engine *gin.Engine, pingRouter *api.PingRouter) *Server {
	return &Server{engine: engine, server: NewHttpServer(8080), pingrouter: pingRouter}
}

func (s *Server) Start() {

	s.pingrouter.Ping(s.engine)

	s.server.Handler = s.engine
	s.server.ListenAndServe()
}

func (s *Server) Stop() {
	s.server.Shutdown(context.Background())
}
