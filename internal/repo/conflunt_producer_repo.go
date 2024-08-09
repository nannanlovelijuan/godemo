package repo

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type confluentProducer struct {
	producer *kafka.Producer
}

func NewConfluentProducerRepo(producer *kafka.Producer) IProducerRepo {
	return &confluentProducer{producer: producer}
}

// Send implements IMsgRepo.
func (p *confluentProducer) Send(topic, msg string) (partition int32, offset int64, err error) {
	// defer p.producer.Close()

	//优化成协程发送

	deliveryChan := make(chan kafka.Event)

	err = p.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(msg),
	}, deliveryChan)

	if err != nil {
		fmt.Printf("Produce failed: %v\n", err)
		return -1, -1, err
	}

	e := <-deliveryChan
	m := e.(*kafka.Message)

	err = m.TopicPartition.Error

	if err != nil {
		fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
		return -1, -1, err
	} else {
		fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
			*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
		return m.TopicPartition.Partition, int64(m.TopicPartition.Offset), err
	}

}
