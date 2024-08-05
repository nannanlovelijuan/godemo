package global

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// server 具有start 和 stop 方法
// server 使用gin做为处理器
type Server struct {
	engine *gin.Engine
	server *http.Server
}

func NewServer(engine *gin.Engine) *Server {
	return &Server{engine: engine, server: NewHttpServer(8080)}
}

func (s *Server) Start() {

	s.server.Handler = s.engine
	s.server.ListenAndServe()
}

func (s *Server) Stop() {
	s.server.Shutdown(context.Background())
}
