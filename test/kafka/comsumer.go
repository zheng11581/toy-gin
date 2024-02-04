package main

import (
	"fmt"

	"github.com/IBM/sarama"
)

// kafka consumer

func NewConsumerConfig() *sarama.Config {
	config := sarama.NewConfig()
	return config
}

func main() {
	consumer, err := sarama.NewConsumer([]string{"172.20.208.166:9092"}, nil)
	if err != nil {
		fmt.Printf("fail to start consumer, err:%v\\n", err)
		return
	}
	// 根据topic取到所有的分区
	partitionList, err := consumer.Partitions("web_log")
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\\n", err)
		return
	}
	fmt.Printf("get partitions:%v\\n", partitionList)
	// 遍历所有的分区
	for partition := range partitionList {
		// 针对每个分区创建一个对应的分区消费者
		fmt.Printf("create consumer for %v \\n", partition)
		pc, err := consumer.ConsumePartition("web_log", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d,err:%v\\n", partition, err)
			return
		}
		defer pc.AsyncClose()
		// 异步从每个分区消费信息
		fmt.Printf("star to get message from %v \\n", partition)
		for msg := range pc.Messages() {
			fmt.Printf("Partition:%d Offset:%d Key:%v Value:%v",
				msg.Partition, msg.Offset, msg.Key, msg.Value)
		}

	}
}
