package api

import (
	"github.com/gin-gonic/gin"
	"gitlab.ezrpro.in/godemo/internal/api/handlers"
)

type KafkaRouter struct {
	KafkaHandler *handlers.KafkaHandler
}

func NewKafkaRouter(kafkaHandler *handlers.KafkaHandler) *KafkaRouter {
	return &KafkaRouter{
		KafkaHandler: kafkaHandler,
	}
}

func (b *KafkaRouter) Produce(engine *gin.Engine) {
	engine.POST("/kafka/produce", b.KafkaHandler.Produce)
}
