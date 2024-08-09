package global

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

// conflunt-kafka生产者对象
func InitConfluntKafkaProducer(app *Application) *kafka.Producer {

	//初始化，从config/db中获取配置
	configMap := setKafkaProducerConfig(app)

	producer, err := kafka.NewProducer(configMap)
	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
	}

	return producer
}

func setKafkaProducerConfig(app *Application) *kafka.ConfigMap {
	_, bootstrapServer := app.Config.GetKv("kafka.bootstrap_server")

	configMap := kafka.ConfigMap{
		"bootstrap.servers": bootstrapServer,
	}
	if b, v := app.Config.GetKv("kafka.acks"); b {
		configMap.SetKey("acks", v)
	}

	if b, v := app.Config.GetKv("kafka.retries"); b {
		configMap.SetKey("retries", v)
	}

	if b, v := app.Config.GetKv("kafka.linger.ms"); b {
		configMap.SetKey("linger.ms", v)
	}

	return &configMap
}

//kafka消费者对象

//初始化kafka生产者

//初始化kafka消费者

// 发送消息
