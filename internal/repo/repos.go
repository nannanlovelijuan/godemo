package repo

import "gitlab.ezrpro.in/godemo/internal/model"

type IProjectRepo interface {
	GetById(id int) (res model.Project, err error)
}
