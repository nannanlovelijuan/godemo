package service

import "gitlab.ezrpro.in/godemo/internal/model"

type IProjectService interface {
	GetById(id int) (model.Project, error)
}
type IProducerService interface {
	Send(topic, msg string) (partition int32, offset int64, err error)
}
