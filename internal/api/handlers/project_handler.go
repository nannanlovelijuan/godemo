package handlers

import (
	"github.com/gin-gonic/gin"
	"gitlab.ezrpro.in/godemo/internal/model"
	"gitlab.ezrpro.in/godemo/internal/service"
	"gitlab.ezrpro.in/godemo/pkg"
)

type ProjectHandler struct {
	ProjectService service.IProjectService
}

func NewProjectHandler(projectService service.IProjectService) *ProjectHandler {
	return &ProjectHandler{
		ProjectService: projectService,
	}
}

func (p *ProjectHandler) GetById(c *gin.Context) {
	// id, _ := strconv.Atoi(c.Param("id"))

	var project model.Project
	err := c.ShouldBindJSON(&project)
	if err != nil {
		pkg.Failure(c, 500, err)
	}
	res, err := p.ProjectService.GetById(int(project.Id))
	if err != nil {
		pkg.Failure(c, 500, err)
		return
	}

	pkg.Success(c, res)
}
