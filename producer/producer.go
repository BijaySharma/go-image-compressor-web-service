package producer

import (
	"github.com/IBM/sarama"
	"github.com/sirupsen/logrus"
)

func PublishMessage(topic, message string) {
	if producer == nil {
		logrus.Info("Producer is not initialized, initializing...")
		InitProducer()
	}
	logrus.Infof("Publishing message %s to topic %s", message, topic)
	_, _, err := producer.SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	})
	if err != nil {
		logrus.Error("Error publishing message to topic ", topic, ": ", err)
	}
}
