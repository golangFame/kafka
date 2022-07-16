package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

func newReader(url, topic string, dialer *kafka.Dialer) *kafka.Reader {

	return kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{url},
		Topic:   topic,
		Dialer:  dialer, //not needed for local
		// GroupID: "g3",
	})
}

var ReadChannel = make(chan []byte, 100000) //unbuffereed channel needs a receiver as soon as msg is send

//reads from a topic
func read(url, topic string, dialer *kafka.Dialer) {
	reader := newReader(url, topic, dialer)
	defer reader.Close()

	for {
		msg, err := reader.ReadMessage(context.Background()) //blocking call
		if try(err, nil) {
			log.Printf("%s rec %s\n", dialer.ClientID, msg.Value)
			time.Sleep(1 * time.Second)
			//ReadChannel <- msg.Value
		}
	}
}
