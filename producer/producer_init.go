package producer

import (
	"github.com/IBM/sarama"
	"github.com/sirupsen/logrus"
)

var producer sarama.SyncProducer

func InitProducer() {
	logrus.Info("Initializing Kafka producer")
	brokers := []string{"localhost:9092"}
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	var err error
	producer, err = sarama.NewSyncProducer(brokers, config)
	if err != nil {
		logrus.Error("Error initializing Kafka producer: ", err)
	}
	logrus.Info("Kafka producer initialized successfully")
}
