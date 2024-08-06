package api

import (
	"github.com/gin-gonic/gin"
	"gitlab.ezrpro.in/godemo/internal/api/handlers"
)

type Routers struct {
	TestRouter     *TestRouter
	TestHandler    *handlers.TestHandler
	ProjectRouter  *ProjectRouter
	ProjectHandler *handlers.ProjectHandler
}

func NewRouters(testHandler *handlers.TestHandler, projectHandler *handlers.ProjectHandler) *Routers {

	testRouter := NewTestRouter(testHandler)
	projectRouter := NewProjectRouter(projectHandler)
	return &Routers{TestRouter: testRouter, ProjectRouter: projectRouter}
}

func RegisterRouters(engine *gin.Engine, routers *Routers) {
	routers.TestRouter.Ping(engine)
	routers.TestRouter.Gracefully(engine)

	routers.ProjectRouter.GetById(engine)
}
