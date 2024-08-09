package api

import (
	"github.com/gin-gonic/gin"
	"gitlab.ezrpro.in/godemo/internal/api/handlers"
)

type Routers struct {
	testRouter     *TestRouter
	projectRouter  *ProjectRouter
	producerRouter *producerRouter
}

func NewRouters(testHandler *handlers.TestHandler, projectHandler *handlers.ProjectHandler, producerHandler *handlers.ProducerHandler) *Routers {

	testRouter := NewTestRouter(testHandler)
	projectRouter := NewProjectRouter(projectHandler)
	producerRouter := NewProducerRouter(producerHandler)
	return &Routers{testRouter: testRouter, projectRouter: projectRouter, producerRouter: producerRouter}
}

func RegisterRouters(engine *gin.Engine, routers *Routers) {
	routers.testRouter.Ping(engine)
	routers.testRouter.Gracefully(engine)

	routers.projectRouter.GetById(engine)

	routers.producerRouter.SendTest(engine)
}
