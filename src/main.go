package main

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
	"os"
	"qmailer/src/email"
)

var log = logrus.New()

func logInfo(msg string) {
	log.WithFields(logrus.Fields{}).Info(msg)
}

func logError(msg string, err error) {
	log.WithError(err).Error(msg)
}

func failOnError(err error, msg string) {
	if err != nil {
		logError(msg, err)
	}
}

func main() {
	emailConfig := email.Config{
		Host: os.Getenv("EMAIL_HOST"),
		Port: os.Getenv("EMAIL_PORT"),
		User: os.Getenv("EMAIL_USER"),
		Pass: os.Getenv("EMAIL_PASS"),
		From: os.Getenv("EMAIL_FROM"),
	}

	smtpWrapper := email.NewSmtpWrapper()
	emailer := email.NewEmailer(smtpWrapper)

	host := os.Getenv("RABBITMQ_HOST")
	queue := os.Getenv("RABBITMQ_QUEUE")

	conn, err := amqp.Dial(host)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer func(conn *amqp.Connection) {
		err := conn.Close()
		if err != nil {
			log.Panicf("Error occurred when closing RabbitMQ connection: %s", err)
		}
	}(conn)

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer func(ch *amqp.Channel) {
		err := ch.Close()
		if err != nil {
			log.Panicf("Error occurred when closing RabbitMQ channel: %s", err)
		}
	}(ch)

	q, err := ch.QueueDeclare(
		queue,
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare the queue")

	messages, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range messages {
			log.Printf("Received a message: %s", d.Body)

			message := email.Email{
				To:      nil,
				Subject: "",
				Body:    "",
			}

			err := emailer.Send(message, emailConfig)
			if err != nil {
				logError("An error occurred when sending email", err)
			}
		}
	}()

	logInfo("Awaiting messages.")
	<-forever
}
