package rabbit

import (
	"log"
	"math/rand"
	"time"

	"github.com/aniketxpawar/go-rabbitmq/config"
	"github.com/streadway/amqp"
)

func RunProducer() {
	rabbitmqURL := config.GetEnv("RABBITMQ_URL","amqp://guest:guest@localhost:5672/")
	exchange := config.GetEnv("EXCHANGE","logs_exchange")

	conn, err := amqp.Dial(rabbitmqURL)
	if err != nil{
		log.Fatalf("Failed to connect to RabbitMQ: %v",err)
	}
	defer conn.Close()

	ch,err := conn.Channel()
	if err != nil{
		log.Fatalf("Failed to open a channel: %v",err)
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		exchange,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil{
		log.Fatalf("Failed to declare a exchange: %v",err)
	}

	log.Println("Producer running. Press CTRL+C to stop.")
	for {
		time.Sleep(2 * time.Second)
		routingKey := "info"
		if rand.Intn(2) == 0 { // 50% chance to switch to "error"
			routingKey = "error"
		}

		body := ""
		if routingKey == "info"{
			body = "Info Message"
		} else {
			body = "Error Message"
		}

		err = ch.Publish(
			exchange,routingKey,false,false,
			amqp.Publishing{
				ContentType: "plain/text",
				Body: []byte(body),
			},
		)

		if err != nil {
			log.Printf("Failed to publish a message: %v", err)
		} else {
			log.Printf("Sent message with routing key '%s': %s", routingKey, body)
		}
	}
}