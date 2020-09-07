package main

import (
	"encoding/json"
	"net/http"

	"github.com/Shopify/sarama"
)

// ReqBody ...
type ReqBody struct {
	Topic           string `json:"topic"`
	Payload         string `json:"payload"`
	PartitioningKey string `json:"key,omitempty"`
}

// http localhost:3000/produceAsync topic='stam' payload='{"a": 1, "b": 2}'
func produceAsync(w http.ResponseWriter, req *http.Request) {
	var body ReqBody

	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	producePayload := sarama.ProducerMessage{
		Topic: body.Topic,
		Value: sarama.StringEncoder(body.Payload),
		Key:   sarama.StringEncoder(body.PartitioningKey),
	}

	asyncProducer.Input() <- &producePayload

	w.WriteHeader(http.StatusOK)
}
