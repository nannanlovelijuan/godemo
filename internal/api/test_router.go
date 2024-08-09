package api

import (
	"github.com/gin-gonic/gin"
	"gitlab.ezrpro.in/godemo/internal/api/handlers"
)

type TestRouter struct {
	testHandler *handlers.TestHandler
}

func NewTestRouter(testHandler *handlers.TestHandler) *TestRouter {
	return &TestRouter{
		testHandler: testHandler,
	}
}

// 服务验证
func (b *TestRouter) Ping(engine *gin.Engine) {
	engine.GET("/test/ping", b.testHandler.Ping)
}

// 优雅停机测试
func (b *TestRouter) Gracefully(engine *gin.Engine) {
	engine.GET("/test/gracefully", b.testHandler.Gracefully)
}
