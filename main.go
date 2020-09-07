package main

import (
	"github.com/Shopify/sarama"
	"gopkg.in/alecthomas/kingpin.v2"
)

type config struct {
	c     *sarama.Config
	hosts *[]string
}

var (
	host = kingpin.Flag("host", "the host to bind to").Short('h').Default("127.0.0.1").Envar("HOST").IP()
	port = kingpin.Flag("port", "the port to bind to").Short('p').Default("3000").Envar("PORT").String()

	seedBrokers = kingpin.Flag("brokers", "the Kafka seed brokers. a string separated list: host:ip,host:ip,host:ip").Short('b').Required().Strings()

	saslUsername = kingpin.Flag("username", "the SASL username to authenticate with. Please use SASL_USERNAME env var").Envar("SASL_USERNAME").String()
	saslPassword = kingpin.Flag("password", "the SASL password to authenticate with. Please use SASL_PASSWORD env var").Envar("SASL_PASSWORD").String()

	asyncProducer sarama.AsyncProducer
)

func main() {
	kingpin.Version(currentVersion())
	kingpin.Parse()

	config := configureKafkaClient()

	asyncProducer = newAsyncProducer(config)
	defer asyncProducer.Close()

	createServer()
}

func configureKafkaClient() *config {
	saramaConfig := sarama.NewConfig()
	saramaConfig.Producer.Return.Errors = true

	if saslUsername != nil && saslPassword != nil {
		saramaConfig.Net.SASL.Enable = true
		saramaConfig.Net.SASL.User = *saslUsername
		saramaConfig.Net.SASL.Password = *saslPassword
	}

	return &config{
		hosts: seedBrokers,
		c:     saramaConfig,
	}
}
