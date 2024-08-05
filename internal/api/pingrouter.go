package api

import (
	"github.com/gin-gonic/gin"
	"gitlab.ezrpro.in/godemo/internal/api/handlers"
)

type PingRouter struct {
	pingHandler *handlers.PingHandler
}

func NewPingRouter(pingHandler *handlers.PingHandler) *PingRouter {
	return &PingRouter{
		pingHandler: pingHandler,
	}
}

func (b *PingRouter) Ping(engine *gin.Engine) {
	engine.GET("/", b.pingHandler.Ping)
}
