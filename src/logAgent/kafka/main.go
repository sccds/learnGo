package kafka

import (
	"fmt"
	"time"

	"github.com/Shopify/sarama"
	"github.com/astaxie/beego/logs"
)

var (
	client sarama.SyncProducer
)

// init kafka client
func InitKafka(addr string) (err error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	client, err = sarama.NewSyncProducer([]string{addr}, config)
	if err != nil {
		fmt.Println("init kafka producer failed, err:", err)
		return
	}
	logs.Debug("init kafka succ")
	return
}

// send data to kafka
func SendToKafka(data, topic string) (err error) {
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(data)
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		logs.Error("send message failed, err: %v data: %v topic: %v", err, data, topic)
		return
	}
	fmt.Printf("pid: %v offset: %v \n", pid, offset)
	time.Sleep(10 * time.Millisecond)
	return
}
