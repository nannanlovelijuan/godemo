package api

import (
	"github.com/gin-gonic/gin"
	"gitlab.ezrpro.in/godemo/internal/api/handlers"
)

type ProjectRouter struct {
	ProjectHandler *handlers.ProjectHandler
}

func NewProjectRouter(projectHandler *handlers.ProjectHandler) *ProjectRouter {
	return &ProjectRouter{
		ProjectHandler: projectHandler,
	}
}

func (b *ProjectRouter) GetById(engine *gin.Engine) {
	engine.POST("/project/getById", b.ProjectHandler.GetById)
}
