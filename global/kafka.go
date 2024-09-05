package global

import (
	"fmt"
	"hash/fnv"

	"github.com/IBM/sarama"
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
	// 0：不确认，1：learder确认，3：leleader和follower确认
	if b, v := app.Config.GetKv("kafka.acks"); b {
		configMap.SetKey("acks", v)
	}

	if b, v := app.Config.GetKv("kafka.retries"); b {
		configMap.SetKey("retries", v)
	}

	if b, v := app.Config.GetKv("kafka.linger.ms"); b {
		configMap.SetKey("linger.ms", v)
	}

	// 注册自定义分区器
	partitioner := &confluentPartitioner{}
	configMap.SetKey("partitioner", partitioner)

	return &configMap
}

func InitSaramaKafkaProducer(app *Application) *sarama.SyncProducer {

	_, bootstrapServer := app.Config.GetKv("kafka.bootstrap_server")
	brokerList := []string{bootstrapServer}

	config := setSaramaKafkaProducerConfig(app)
	producer, err := sarama.NewSyncProducer(brokerList, config)

	if err != nil {
		fmt.Printf("Failed to create InitSaramaKafkaProducer: %v\n", err)
	}
	return &producer
}

func setSaramaKafkaProducerConfig(_ *Application) *sarama.Config {

	config := sarama.NewConfig()
	config.Version = sarama.V0_11_0_2
	config.Producer.Return.Successes = true

	//自定义分区器
	config.Producer.Partitioner = NewSaramaPartitioner
	return config
}

//kafka消费者对象

//初始化kafka生产者

//初始化kafka消费者

// 发送消息

// sarama自定义分区器
type saramaPartitioner struct {
	partition int32
}

// Partition 返回的是分区的位置或者索引，并不是具体的分区号。比如有十个分区[0,1，2,3...9] 这里返回 0 表示取数组中的第0个位置的分区。在 Go 客户端中是这样实现的，具体见下文源码分析
// 轮询发送的自定义分区
func (p *saramaPartitioner) Partition(message *sarama.ProducerMessage, numPartitions int32) (int32, error) {
	if p.partition >= numPartitions {
		p.partition = 0
	}
	ret := p.partition
	p.partition++
	return ret, nil
}

// 该方法的作用在下文源码分析中有详细解释
func (p *saramaPartitioner) RequiresConsistency() bool {
	return false
}

func NewSaramaPartitioner(topic string) sarama.Partitioner {
	return &saramaPartitioner{}
}

type confluentPartitioner struct{}

func (p *confluentPartitioner) Partition(topic string, key, value []byte, numPartitions int32) int32 {
	// 自定义分区逻辑，例如根据 key 或 value 的哈希值来决定分区

	hasher := fnv.New64()
	// 使用 Write 方法写入数据到哈希器
	_, err := hasher.Write(key)
	if err != nil {
		fmt.Println("Error writing data:", err)
		return 0
	}
	hash := hasher.Sum64()
	return int32(hash % uint64(numPartitions))
}
