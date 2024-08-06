package service

import (
	"gitlab.ezrpro.in/godemo/internal/model"
	"gitlab.ezrpro.in/godemo/internal/repo"
)

type projectService struct {
	mysqlProjectRepo repo.IProjectRepo
}

func NewProjectService(mysqlProjectRepo repo.IProjectRepo) IProjectService {
	return &projectService{
		mysqlProjectRepo: mysqlProjectRepo,
	}
}

func (p *projectService) GetById(id int) (model.Project, error) {
	return p.mysqlProjectRepo.GetById(id)
}
