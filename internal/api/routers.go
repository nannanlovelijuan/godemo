package api

import (
	"github.com/gin-gonic/gin"
	"gitlab.ezrpro.in/godemo/internal/api/handlers"
)

type Routers struct {
	PingRouter  *PingRouter
	PingHandler *handlers.PingHandler
}

func NewRouters(pingHandler *handlers.PingHandler) *Routers {

	pingRouter := NewPingRouter(pingHandler)

	return &Routers{PingRouter: pingRouter}
}

func RegisterRouters(engine *gin.Engine, routers *Routers) {
	routers.PingRouter.Ping(engine)
}
