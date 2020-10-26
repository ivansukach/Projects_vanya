package main

import (
	"fmt"
	"github.com/ivansukach/rabbitmq-server2/config"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"golang.org/x/net/websocket"
	"time"
)
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func changeRole(c echo.Context, selector *bool) error {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()
		for {
			log.Println("Handler")
			msg := ""
			err := websocket.Message.Receive(ws, &msg)
			if err != nil {
				c.Logger().Error(err)
				return
			}
			fmt.Printf("%s\n", msg)
			switch msg {
			case "server2:producer":
				log.Println("Let's produce messages")
				*selector = true
			case "server1:producer":
				log.Println("Let's read messages")
				*selector = false
			}
		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}

func main() {
	cfg := config.Load()
	selector := true
	go messageExchange(&selector, cfg)
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/", "public")
	e.GET("/ws", func(c echo.Context) error {
		return changeRole(c, &selector)
	})
	e.Logger.Fatal(e.Start(fmt.Sprintf(":"+cfg.Port)))

}

func messageExchange(selector *bool, cfg config.Config) {
	log.Println("Port:", cfg.Port)
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
	err = channel.ExchangeDeclare("messages", "topic", true, false, false, false, nil)
	if err != nil {
		log.Error(err)
	}

	// We create a queue named server1
	_, err = channel.QueueDeclare("server1", true, false, false, false, nil)

	if err != nil {
		log.Error("error declaring the queue: " + err.Error())
	}

	// We bind the queue to the exchange to send and receive data from the queue
	err = channel.QueueBind("server1", "#", "messages", false, nil)

	if err != nil {
		log.Error("error binding to the queue: " + err.Error())
	}

	// We create a queue named server2
	_, err = channel.QueueDeclare("server2", true, false, false, false, nil)

	if err != nil {
		log.Error("error declaring the queue: " + err.Error())
	}

	// We bind the queue to the exchange to send and receive data from the queue
	err = channel.QueueBind("server2", "#", "messages", false, nil)

	if err != nil {
		log.Error("error binding to the queue: " + err.Error())
	}

	// We create a queue named nod
	_, err = channel.QueueDeclare("nod", true, false, false, false, nil)

	if err != nil {
		log.Error("error declaring the queue: " + err.Error())
	}

	// We bind the queue to the exchange to send and receive data from the queue
	err = channel.QueueBind("nod", "#", "messages", false, nil)

	if err != nil {
		log.Error("error binding to the queue: " + err.Error())
	}


	tMes := time.NewTicker(time.Second * 4)
	for {
		select {
		case <-tMes.C:
			if *selector {
				log.Println("write")
				message := amqp.Publishing{
					Body: []byte((fmt.Sprintf("server2 say: %s", time.Now().UTC().String()))),
				}

				// We publish the message to the exchange we created earlier
				err = channel.Publish("messages", "random-key", false, false, message)

				if err != nil {
					log.Error("error publishing a message to the queue:" + err.Error())
				}
			} else {
				log.Println("read")
				messages, err := channel.Consume("server2", "", false, false, false, false, nil)

				if err != nil {
					log.Error("error consuming the queue: " + err.Error())
				}

				// We loop through the messages in the queue and print them to the console.
				// The msgs will be a go channel, not an amqp channel
				for msg := range messages {
					//print the message to the console
					fmt.Println("message received: " + string(msg.Body))
					// Acknowledge that we have received the message so it can be removed from the queue
					msg.Ack(false)
				}

			}
		}
	}

}

