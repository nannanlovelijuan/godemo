package repo

import "gitlab.ezrpro.in/godemo/internal/model"

type IProjectRepo interface {
	GetById(id int) (res model.Project, err error)
}

type IProducerRepo interface {
	//同步发送
	Send(topic, msg string) (partition int32, offset int64, err error)

	//异步发送
	AsyncSend(topic, msg string) (partition int32, offset int64, err error)
}
