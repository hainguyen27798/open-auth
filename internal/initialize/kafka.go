package initialize

import (
	"github.com/open-auth/global"
	"github.com/segmentio/kafka-go"
)

func InitKafka() {
	global.SMTPProducer = &kafka.Writer{
		Addr:     kafka.TCP("localhost:9094"),
		Topic:    "send-mail-topic",
		Balancer: &kafka.LeastBytes{},
	}
}
