package service

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"gitlab.ezrpro.in/godemo/internal/repo"
)

type KafkaService struct {
	// ProducerRepo repo.Confl
}

func NewKafkaService() IKafkaService {
	return &KafkaService{}
}

// ProducerByConfluent implements IKafkaService.
func (k *KafkaService) ProducerByConfluent(client *kafka.Producer, topic string, msg string) (partition int32, offset int64, err error) {
	repo := repo.NewConfluentProducerRepo(client)

	partition, offset, err = repo.Send(topic, msg)
	client.Close()
	return
}
