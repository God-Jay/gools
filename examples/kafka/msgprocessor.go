package main

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/god-jay/gools/pkg/kafka"
	"log"
)

var _ kafka.Processor = (*MsgProcessor)(nil)

type MsgProcessor struct {
}

func NewMsgProcessor() *MsgProcessor {
	return &MsgProcessor{}
}

func (mp *MsgProcessor) Handle(ctx context.Context, saramaMsg *sarama.ConsumerMessage) error {
	// do something
	log.Println(string(saramaMsg.Value))
	return nil
}
