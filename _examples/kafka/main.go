package main

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/god-jay/gools/kafka"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	cfg := &kafka.Config{
		Brokers:  []string{"localhost:9092", "localhost:9093"},
		Version:  "2.7.1",
		ClientID: "example_client",
		Oldest:   true,
		Verbose:  true,
	}
	kfk := kafka.New(cfg, func(c *sarama.Config) {
		c.Producer.Return.Successes = true
		c.Consumer.Group.Session.Timeout = 10 * time.Second
	})

	kfk.AddConsumer(NewMsgProcessor(), "topic", "group")

	kfk.RunConsumers(ctx)

	// Wait for SIGINT and SIGTERM (HIT CTRL-C)
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-quit

	cancel()
}
