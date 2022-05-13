package main

import (
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

func (mp *MsgProcessor) Handle(session sarama.ConsumerGroupSession, saramaMsg *sarama.ConsumerMessage) error {
	// do something
	log.Printf("Message claimed: value = %s, timestamp = %v, topic = %s, offset = %d, partition = %d", string(saramaMsg.Value), saramaMsg.Timestamp, saramaMsg.Topic, saramaMsg.Offset, saramaMsg.Partition)
	ctx := session.Context()
	log.Println(ctx, string(saramaMsg.Value))

	// mark if some condition
	session.MarkMessage(saramaMsg, "")

	return nil
}
