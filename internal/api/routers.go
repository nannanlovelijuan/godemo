package api

import (
	"github.com/gin-gonic/gin"
	"gitlab.ezrpro.in/godemo/internal/api/handlers"
)

type Routers struct {
	testRouter     *TestRouter
	projectRouter  *ProjectRouter
	producerRouter *producerRouter
	kafkaRouter    *KafkaRouter
}

func NewRouters(testHandler *handlers.TestHandler, projectHandler *handlers.ProjectHandler, producerHandler *handlers.ProducerHandler, kafkaHandler *handlers.KafkaHandler) *Routers {

	testRouter := NewTestRouter(testHandler)
	projectRouter := NewProjectRouter(projectHandler)
	producerRouter := NewProducerRouter(producerHandler)
	kafkaRouter := NewKafkaRouter(kafkaHandler)
	return &Routers{testRouter: testRouter, projectRouter: projectRouter, producerRouter: producerRouter, kafkaRouter: kafkaRouter}
}

func RegisterRouters(engine *gin.Engine, routers *Routers) {
	routers.testRouter.Ping(engine)
	routers.testRouter.Gracefully(engine)

	routers.projectRouter.GetById(engine)

	routers.producerRouter.SendTest(engine)

	routers.kafkaRouter.Produce(engine)
}
