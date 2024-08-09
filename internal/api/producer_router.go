package api

import (
	"github.com/gin-gonic/gin"
	"gitlab.ezrpro.in/godemo/internal/api/handlers"
)

type producerRouter struct {
	ProducerHandler *handlers.ProducerHandler
}

func NewProducerRouter(producerHandler *handlers.ProducerHandler) *producerRouter {
	return &producerRouter{
		ProducerHandler: producerHandler,
	}
}

func (b *producerRouter) SendTest(engine *gin.Engine) {
	engine.POST("/producer/:topic", b.ProducerHandler.SendTest)
}
