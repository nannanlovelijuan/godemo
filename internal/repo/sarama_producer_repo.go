package repo

import (
	"fmt"

	"github.com/IBM/sarama"
)

type saramaKafkaProducer struct {
	producer *sarama.SyncProducer
}

// AsyncSend implements IProducerRepo.
func (i *saramaKafkaProducer) AsyncSend(topic string, msg string) (partition int32, offset int64, err error) {
	panic("unimplemented")
}

// Send implements IProducerRepo.
func (i *saramaKafkaProducer) Send(topic string, msg string) (partition int32, offset int64, err error) {

	fmt.Println("Send msg by sarama...")

	if i == nil || i.producer == nil {
		return 0, 0, fmt.Errorf("producer is not initialized")
	}

	producer := *i.producer
	message := &sarama.ProducerMessage{Topic: topic, Value: sarama.StringEncoder(msg)}
	partition, offset, err = producer.SendMessage(message)
	if err != nil {
		fmt.Printf("Failed to send message: %v", err)
	}
	return
}

func NewSaramaKafkaProducer(producer *sarama.SyncProducer) IProducerRepo {
	return &saramaKafkaProducer{
		producer: producer,
	}
}
