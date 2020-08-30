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

	return asyncProducer
}
