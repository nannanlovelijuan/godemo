package service

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"gitlab.ezrpro.in/godemo/internal/model"
)

type IProjectService interface {
	GetById(id int) (model.Project, error)
}
type IProducerService interface {
	Send(topic, msg string) (partition int32, offset int64, err error)
}

type IKafkaService interface {
	ProducerByConfluent(client *kafka.Producer, topic, msg string) (partition int32, offset int64, err error)

	// Producer(topic []string, msg string) (partition int32, offset int64, err error)
}
