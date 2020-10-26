package main

import (
	"fmt"
	"github.com/ivansukach/rabbitmq-nod-server/config"
	"github.com/ivansukach/rabbitmq-nod-server/repository"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"time"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	cfg:=config.Load()
	db, err := sqlx.Connect("postgres", "user=su password=su dbname=messages sslmode=disable")
	if err != nil {
		log.Error(err)
	}
	rps := repository.New(db)
	defer db.Close()


	log.Println("URL of connection to RabbitMQ:", cfg.RabbitMQUrl)
	connection, err := amqp.Dial(cfg.RabbitMQUrl)
	defer connection.Close()
	if err != nil {
		log.Error("could not establish connection with RabbitMQ:", err.Error())
	}
	channel, err := connection.Channel()
	if err != nil {
		log.Error("could not open RabbitMQ channel:" + err.Error())
	}


	tMes := time.NewTicker(time.Second * 1)
	for {
		select {
		case <-tMes.C:
			log.Println("read")
			messages, err := channel.Consume("nod", "", false, false, false, false, nil)

			if err != nil {
				log.Error("error consuming the queue: " + err.Error())
			}

			// We loop through the messages in the queue and print them to the console.
			// The msgs will be a go channel, not an amqp channel
			for msg := range messages {
				//print the message to the console
				message:=repository.Message{Content: string(msg.Body)}
				fmt.Println("message received: " + message.Content)
				rps.Create(&message)
				// Acknowledge that we have received the message so it can be removed from the queue
				msg.Ack(false)
			}
			}

	}
}