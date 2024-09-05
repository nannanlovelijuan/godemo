package repo

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type confluentProducer struct {
	producer *kafka.Producer
}

// AsyncSend implements IProducerRepo.
func (p *confluentProducer) AsyncSend(topic string, msg string) (partition int32, offset int64, err error) {
	fmt.Println("Send async msg by Confluent...")

	err = p.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(msg),
	}, nil)

	if err != nil {
		fmt.Printf("Delivery failed: %v\n", err)
		return -1, -1, err
	}

	go func() {
		for e := range p.producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					partition, offset, err = -1, -1, ev.TopicPartition.Error
					fmt.Printf("Failed to deliver message: %v,msg:%s\n", ev.TopicPartition, msg)
				} else {
					partition, offset = ev.TopicPartition.Partition, int64(ev.TopicPartition.Offset)
					fmt.Printf("Successfully produced record to topic %s partition [%d] @ offset %v\n",
						topic, partition, offset)
				}
			}
		}
	}()

	return
}

// Send implements IMsgRepo.
func (p *confluentProducer) Send(topic, msg string) (partition int32, offset int64, err error) {
	// defer p.producer.Close()

	//优化成协程发送

	fmt.Println("Send msg by Confluent...")

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

	defer close(deliveryChan)
	if err != nil {
		fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
		return -1, -1, err
	} else {
		fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
			*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
		return m.TopicPartition.Partition, int64(m.TopicPartition.Offset), err
	}

}

func NewConfluentProducerRepo(producer *kafka.Producer) IProducerRepo {
	return &confluentProducer{producer: producer}
}
