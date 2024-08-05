package handlers

import (
	"github.com/gin-gonic/gin"
	"gitlab.ezrpro.in/godemo/internal/pkg"
)

type PingHandler struct {
}

func NewPingHandler() *PingHandler {
	return &PingHandler{}
}

func (p *PingHandler) Ping(c *gin.Context) {
	pkg.Success(c, "ok")
}
