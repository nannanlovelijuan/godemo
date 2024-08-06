package api

import (
	"github.com/gin-gonic/gin"
	"gitlab.ezrpro.in/godemo/internal/api/handlers"
)

type TestRouter struct {
	pingHandler *handlers.TestHandler
}

func NewTestRouter(pingHandler *handlers.TestHandler) *TestRouter {
	return &TestRouter{
		pingHandler: pingHandler,
	}
}

// 服务验证
func (b *TestRouter) Ping(engine *gin.Engine) {
	engine.GET("/test/ping", b.pingHandler.Ping)
}

// 优雅停机测试
func (b *TestRouter) Gracefully(engine *gin.Engine) {
	engine.GET("/test/gracefully", b.pingHandler.Gracefully)
}
