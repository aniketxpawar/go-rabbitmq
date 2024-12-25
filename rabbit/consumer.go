package rabbit

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/aniketxpawar/go-rabbitmq/config"
	"github.com/streadway/amqp"
)

func RunConsumer1(queueName string) {
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
	
	q,err := ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil{
		log.Fatalf("Failed to declare a queue: %v",err)
	}
	
	routingKey := "info"
	err = ch.QueueBind(
		q.Name,
		routingKey,
		exchange,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to bind queue: %v", err)
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	log.Printf("Consumer running on queue '%s'. Press CTRL+C to stop.", queueName)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func(){
		for msg := range msgs {
			log.Printf(" [x] Received %s", msg.Body)
		}
	}()

	<- sigChan
	log.Println("Consumer shutting down.")
}

func RunConsumer2(queueName string){
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
	
	q,err := ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil{
		log.Fatalf("Failed to declare a queue: %v",err)
	}
	
	routingKey := "error"
	err = ch.QueueBind(
		q.Name,
		routingKey,
		exchange,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to bind queue: %v", err)
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	log.Printf("Consumer running on queue '%s'. Press CTRL+C to stop.", queueName)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func(){
		for msg := range msgs {
			log.Printf(" [x] Received %s", msg.Body)
		}
	}()

	<- sigChan
	log.Println("Consumer shutting down.")
}