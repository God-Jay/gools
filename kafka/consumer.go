package kafka

import (
	"context"
	"github.com/Shopify/sarama"
	"log"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
)

type Consumer struct {
	Processor Processor
	Topics    string // topic1,topic2,topic3
	Group     string
	Ready     chan bool
}

type ConfOptionFunc func(c *sarama.Config)

func NewConsumer(processor Processor, topics string, group string) *Consumer {
	return &Consumer{
		Processor: processor,
		Topics:    topics,
		Group:     group,
		Ready:     make(chan bool),
	}
}

func (c *Consumer) Run(ctx context.Context, conf *Config, option ConfOptionFunc) {
	consumerGroup, err := c.saramaConsumerGroup(conf, c.Group, option)
	if err != nil {
		log.Panicf("Error creating consumer group client: %v", err)
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if err := consumerGroup.Consume(ctx, strings.Split(c.Topics, ","), c); err != nil {
				log.Panicf("Error from consumer: %v", err)
			}
			// check if context was cancelled, signaling that the consumer should stop
			if ctx.Err() != nil {
				return
			}
			c.Ready = make(chan bool)
		}
	}()

	<-c.Ready // Await till the consumer has been set up
	log.Println("Sarama consumer up and running!...")

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-ctx.Done():
		log.Println("terminating: context cancelled")
	case <-sigterm:
		log.Println("terminating: via signal")
	}
	wg.Wait()
	if err = consumerGroup.Close(); err != nil {
		log.Panicf("Error closing client: %v", err)
	}
}

func (c *Consumer) saramaConsumerGroup(conf *Config, group string, option ConfOptionFunc) (sarama.ConsumerGroup, error) {
	if conf.Verbose {
		sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)
	}

	config := sarama.NewConfig()

	version, err := sarama.ParseKafkaVersion(conf.Version)
	if err != nil {
		log.Panicf("Error parsing Kafka version: %v", err)
	}
	config.Version = version
	config.ClientID = conf.ClientID

	if conf.Oldest {
		config.Consumer.Offsets.Initial = sarama.OffsetOldest
	}

	option(config)

	client, err := sarama.NewConsumerGroup(conf.Brokers, group, config)

	return client, err
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (c *Consumer) Setup(session sarama.ConsumerGroupSession) error {
	// Mark the consumer as ready
	close(c.Ready)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (c *Consumer) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (c *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		err := c.Processor.Handle(session, msg)
		if err != nil {
			return err
		}
	}
	return nil
}
