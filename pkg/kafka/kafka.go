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
	conf      *Config
	Consumers []*Consumer
}

func New(conf *Config) *Kafka {
	return &Kafka{conf: conf}
}

func (k *Kafka) AddConsumer(processor Processor, topic string, group string) {
	k.Consumers = append(k.Consumers, NewConsumer(processor, topic, group))
}

func (k *Kafka) RunConsumer(ctx context.Context) {
	for _, consumer := range k.Consumers {
		go consumer.Run(ctx, k.conf)
	}
}
