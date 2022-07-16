package kafka

import (
	"fmt"
	"github.com/segmentio/kafka-go"
	"os"
	"strconv"
)

func StartKafka(topic, cGID string) {

	username := os.Getenv("username")
	password := os.Getenv("pass")
	url := os.Getenv("bootstrap_servers")
	fmt.Println("Starting kafka")

	dialer := newDialer("kafka-test", username, password)

	topicConfig := kafka.TopicConfig{
		Topic:             topic,
		NumPartitions:     220,
		ReplicationFactor: 3, //1 doesnt work
	}
	createTopic(url, dialer, topicConfig)

	inc := 100
	z, y := 0, inc

	for i := 0; i < 100; i++ {
		dialer := newDialer(strconv.Itoa(i), username, password)
		go write(url, topic, dialer, z, y)
		//go read(url, topic, dialer)
		z, y = z+inc, y+inc
	}
}

func try(err error, errorHandler func(...interface{}) (int, error)) bool {
	if err == nil {
		return true
	}
	if errorHandler == nil {
		panic(err.Error())
	}
	errorHandler(string(err.Error()))
	return false
}
