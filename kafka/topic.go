package kafka

import (
	"fmt"
	"github.com/segmentio/kafka-go"
	"net"
	"strconv"
)

func createTopic(url string, dialer *kafka.Dialer, topic kafka.TopicConfig) {
	conn, err := dialer.Dial("tcp", url)
	try(err, nil)

	defer conn.Close()

	controller, err := conn.Controller()
	try(err, nil)

	controllerConn, err := dialer.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	try(err, nil)
	defer controllerConn.Close()

	err = controllerConn.CreateTopics(topic)
	try(err, fmt.Println) //topic probably exists

}
