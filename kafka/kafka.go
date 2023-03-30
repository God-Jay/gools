package kafka

import (
	"context"
)

type Config struct {
	Brokers  []string
	Version  string
	ClientID string // a user-provided string sent with every request to the brokers for logging, debugging, and auditing purposes.
	Oldest   bool   // if true, fetch oldest available offset
	Verbose  bool   // if true, logs are printed to stdout
}

type Kafka struct {
	conf       *Config
	confOption ConfOptionFunc
	Consumers  []*Consumer
}

func New(conf *Config, confOption ConfOptionFunc) *Kafka {
	return &Kafka{conf: conf, confOption: confOption}
}

// AddConsumer add a specific consumer to this receiver to handle the topics using the given group
// To handle multiple topics by this processor, use `,` to separate the topics, e.g. `"topic1,topic2"`
func (k *Kafka) AddConsumer(processor Processor, topics string, group string) *Consumer {
	c := NewConsumer(processor, topics, group)
	k.Consumers = append(k.Consumers, c)
	return c
}

// RunConsumers runs all this kafka receiver's consumers using sarama consumer group.
// Sarama consumer group runs in multiple goroutines based on the number of its topic's partition num.
// If you add 2 consumers, and each consumer's topic has 3 partitions, this will run 2*3 consumer goroutines.
func (k *Kafka) RunConsumers(ctx context.Context) {
	for _, consumer := range k.Consumers {
		go consumer.Run(ctx, k.conf, k.confOption)
	}
}
