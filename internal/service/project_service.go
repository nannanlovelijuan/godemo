package service

import (
	"gitlab.ezrpro.in/godemo/internal/model"
	"gitlab.ezrpro.in/godemo/internal/repo"
)

type projectService struct {
	projectRepo repo.IProjectRepo
}

func NewProjectService(projectRepo repo.IProjectRepo) IProjectService {
	return &projectService{
		projectRepo: projectRepo,
	}
}

func (p *projectService) GetById(id int) (model.Project, error) {
	return p.projectRepo.GetById(id)
}
