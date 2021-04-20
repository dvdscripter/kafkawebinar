package main

import (
	"context"
	"log"

	"github.com/Shopify/sarama"
)

type kafkaHandler struct{}

func (kafkaHandler) Setup(_ sarama.ConsumerGroupSession) error {
	log.Println("READY TO CONSUME")
	return nil
}
func (kafkaHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (kafkaHandler) ConsumeClaim(cgs sarama.ConsumerGroupSession, cgc sarama.ConsumerGroupClaim) error {
	for msg := range cgc.Messages() {
		log.Println("MSG", string(msg.Key), string(msg.Value))
		cgs.MarkMessage(msg, "")
	}

	return nil
}

func main() {
	conf := sarama.NewConfig()
	conf.Version = sarama.V2_5_0_0
	conf.ClientID = "consumer"
	conf.Consumer.Offsets.Initial = sarama.OffsetOldest

	cg, err := sarama.NewConsumerGroup([]string{"kafka:9092"}, "kafkawebinar-group", conf)
	if err != nil {
		panic(err)
	}
	defer cg.Close()

	ctx := context.Background()
	handler := kafkaHandler{}
	for {
		if err := cg.Consume(ctx, []string{"webinar"}, handler); err != nil {
			log.Println("ERROR", err)
		}
	}

}
