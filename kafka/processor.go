package kafka

import (
	"github.com/Shopify/sarama"
)

// Processor is the interface you need to implement to write your logic the handle the messages.
type Processor interface {
	// Handle processes a message get from a broker.
	// You must finish processing and mark offsets within
	// sarama.Config.Consumer.Group.Session.Timeout before the topic/partition is eventually
	// re-assigned to another group member.
	Handle(session sarama.ConsumerGroupSession, saramaMsg *sarama.ConsumerMessage) error
}
