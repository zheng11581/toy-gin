package main

import (
	"fmt"

	"github.com/IBM/sarama"
)

func NewProducerConfig() *sarama.Config {
	config := sarama.NewConfig()
	// 发送完数据需要leader和follow都确认
	config.Producer.RequiredAcks = sarama.WaitForAll
	// 新选出一个partition
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	// 成功交付的消息将在success channel返回
	config.Producer.Return.Successes = true

	return config
}

func NewProducer(addrs []string, config *sarama.Config) (sarama.SyncProducer, error) {
	// 连接kafka
	client, err := sarama.NewSyncProducer(addrs, config)
	return client, err
}

func SendMessage(client sarama.SyncProducer, topic string, value sarama.StringEncoder) {
	// 构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = value

	defer client.Close()
	// 发送消息
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed, err:", err)
		return
	}
	fmt.Printf("pid:%v offset:%v\n", pid, offset)
}

// func main() {
// 	config := NewProducerConfig()
// 	addrs := []string{
// 		"broker1:9092",
// 		"broker2:9092",
// 		"broker3:9092",
// 	}
// 	client, err := NewProducer(addrs, config)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	SendMessage(client, "transaction", "this is test log")

// }
