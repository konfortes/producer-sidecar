package main

import (
	"log"

	"github.com/Shopify/sarama"
)

func newAsyncProducer(conf *config) sarama.AsyncProducer {
	asyncProducer, err := sarama.NewAsyncProducer(*conf.hosts, conf.c)
	if err != nil {
		log.Fatal(err)
	}

	go handleErrors(asyncProducer)

	return asyncProducer
}

func handleErrors(producer sarama.AsyncProducer) {
	for {
		err := <-producer.Errors()
		if err != nil {
			log.Println("Failed to produce message", err)
		}
	}
}
