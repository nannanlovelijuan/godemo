package repo

import "gitlab.ezrpro.in/godemo/internal/model"

type IProjectRepo interface {
	GetById(id int) (res model.Project, err error)
}

type IProducerRepo interface {
	Send(topic, msg string) (partition int32, offset int64, err error)
}
