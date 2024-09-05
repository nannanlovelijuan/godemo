package handlers

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/gin-gonic/gin"
	"gitlab.ezrpro.in/godemo/internal/service"
	"gitlab.ezrpro.in/godemo/pkg"
)

type KafkaHandler struct {
	KafkaService service.IKafkaService
}

func NewKafkaHandler(kf service.IKafkaService) *KafkaHandler {
	return &KafkaHandler{
		KafkaService: kf,
	}
}

func (k *KafkaHandler) Produce(c *gin.Context) {
	// 获取参数

	var req struct {
		Client  int      `json:"client"`
		Topics  []string `json:"topics"`
		Message string   `json:"message"`
		Server  int      `json:"server"`
	}
	c.ShouldBindJSON(&req)

	// todo校验参数
	if req.Client == 1 {
		server := ""
		if req.Server == 1 {
			server = "192.168.128.216:9092"
		}

		client, err := newConfluentProducerClient(server)
		if err != nil {
			pkg.Failure(c, 500, fmt.Errorf("create confluent client failed %v", err))
			return
		}

		//todo 优化的点发送两个topic，存在一个成功一个失败的场景，所以尽量发送一个
		resData := make(map[string]interface{})
		for _, topic := range req.Topics {
			partition, offset, err := k.KafkaService.ProducerByConfluent(client, topic, req.Message)
			if err != nil {
				pkg.Failure(c, 500, fmt.Errorf("send topic %s message failed %v", topic, err))
				return
			}
			resData[topic] = gin.H{
				"partition": partition,
				"offset":    offset,
			}
		}

		pkg.Success(c, resData)
	} else {
		//todo sarama client
		pkg.Failure(c, 500, fmt.Errorf("client type not support"))
	}
}

func newConfluentProducerClient(server string) (producer *kafka.Producer, err error) {
	configMap := kafka.ConfigMap{
		"bootstrap.servers": server,
		"client.id":         "confluent-client",
		"acks":              "all",
	}
	producer, err = kafka.NewProducer(&configMap)
	return
}
