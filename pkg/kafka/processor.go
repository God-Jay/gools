package kafka

import (
	"context"
	"github.com/Shopify/sarama"
)

type Processor interface {
	Handle(ctx context.Context, msg *sarama.ConsumerMessage) error
}
