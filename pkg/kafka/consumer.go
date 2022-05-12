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

func NewConsumer(processor Processor, topics string, group string) *Consumer {
	return &Consumer{
		Processor: processor,
		Topics:    topics,
		Group:     group,
		Ready:     make(chan bool),
	}
}

func (c *Consumer) Run(ctx context.Context, conf *Config) {
	consumerGroup, err := c.saramaConsumerGroup(conf, c.Group)
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

func (c *Consumer) saramaConsumerGroup(conf *Config, group string) (sarama.ConsumerGroup, error) {
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

	client, err := sarama.NewConsumerGroup(conf.Brokers, group, config)

	return client, err
}

func (c *Consumer) Setup(session sarama.ConsumerGroupSession) error {
	// Mark the consumer as ready
	close(c.Ready)
	return nil
}

func (c *Consumer) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (c *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		log.Printf("Message claimed: value = %s, timestamp = %v, topic = %s, offset = %d, partition = %d", string(msg.Value), msg.Timestamp, msg.Topic, msg.Offset, msg.Partition)

		err := c.Processor.Handle(session.Context(), msg)
		if err != nil {
			return err
		}
		session.MarkMessage(msg, "")
	}
	return nil
}
