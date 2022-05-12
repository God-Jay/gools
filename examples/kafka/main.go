package main

import (
	"context"
	"github.com/god-jay/gools/pkg/kafka"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	kfk := kafka.New(&kafka.Config{
		Brokers:  []string{"localhost:9092", "localhost:9093"},
		Version:  "2.7.1",
		ClientID: "example_client",
		Oldest:   true,
		Verbose:  true,
	})

	kfk.AddConsumer(NewMsgProcessor(), "topic", "group")

	kfk.RunConsumer(ctx)

	// Wait for SIGINT and SIGTERM (HIT CTRL-C)
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-quit

	cancel()
}
