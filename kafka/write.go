package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
	"strconv"
	"time"
)

func newWriter(url string, topic string, dialer *kafka.Dialer) *kafka.Writer {
	x := kafka.NewWriter(kafka.WriterConfig{
		Brokers:      []string{url},
		Topic:        topic,
		Balancer:     &kafka.CRC32Balancer{},
		Dialer:       dialer,
		BatchSize:    10,
		BatchTimeout: 1 * time.Millisecond,
	})
	return x
}

func write(url string, topic string, dialer *kafka.Dialer, ranges ...int) {
	i, y := 0, 100
	if len(ranges) > 1 {
		i, y = ranges[0], ranges[1]
	}
	writer := newWriter(url, topic, dialer)
	defer writer.Close()

	for i := i; i < y; i++ {
		v := []byte("V" + strconv.Itoa(i))
		log.Printf("ID-%s send:\t%s\n", dialer.ClientID, v)
		msg := kafka.Message{Key: []byte("testing"), Value: v}
		err := writer.WriteMessages(context.Background(), msg)
		try(err, nil)

		// go func(msg kafka.Message) {
		// 	err := writer.WriteMessages(context.Background(), msg)
		// 	try(err, nil)
		// }(msg)
		// time.Sleep(100 * time.Millisecond)
	}
}
