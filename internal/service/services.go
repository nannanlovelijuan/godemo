package service

import "gitlab.ezrpro.in/godemo/internal/model"

type IProjectService interface {
	GetById(id int) (model.Project, error)
}
