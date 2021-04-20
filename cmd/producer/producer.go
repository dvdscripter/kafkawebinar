package main

import (
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/Shopify/sarama"
)

func main() {
	conf := sarama.NewConfig()
	conf.Version = sarama.V2_5_0_0
	conf.ClientID = "producer"
	conf.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer([]string{"kafka:9092"}, conf)
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	for c := 0; ; c++ {
		msg := &sarama.ProducerMessage{
			Topic: "webinar",
			Key:   sarama.StringEncoder(strconv.Itoa(c)),
			Value: sarama.StringEncoder(strconv.Itoa(rand.Intn(100))),
		}
		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			log.Println("ERRO", err)
			continue
		}

		log.Println("Partition", partition, "Offset", offset)

		time.Sleep(time.Second)
	}
}
