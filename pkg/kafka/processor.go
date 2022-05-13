package kafka

import (
	"github.com/Shopify/sarama"
)

type Processor interface {
	Handle(session sarama.ConsumerGroupSession, saramaMsg *sarama.ConsumerMessage) error
}
