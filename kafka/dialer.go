package kafka

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
	"time"
)

func newDialer(clientID, username, password string) *kafka.Dialer {
	mechanism := plain.Mechanism{
		Username: username,
		Password: password,
	}

	rootCAs, _ := x509.SystemCertPool()
	if rootCAs == nil {
		rootCAs = x509.NewCertPool()
	}

	return &kafka.Dialer{
		Timeout:       15 * time.Second,
		DualStack:     true,
		ClientID:      clientID,
		SASLMechanism: mechanism,
		TLS: &tls.Config{
			InsecureSkipVerify: false,
			RootCAs:            rootCAs,
		},
	}
}
