package main

import (
	"log"
	"os"

	"github.com/aniketxpawar/go-rabbitmq/config"
	"github.com/aniketxpawar/go-rabbitmq/rabbit"
)

func main() {
	config.LoadConfig()

	if len(os.Args) < 2{
		log.Fatalf("Usage: %s [producer|consumer1|consumer2]",os.Args[0])
	}

	switch os.Args[1] {
	case "producer":
		rabbit.RunProducer()
	case "consumer1":
		rabbit.RunConsumer1(config.GetEnv("QUEUE1","info_queue"))
	case "consumer2":
		rabbit.RunConsumer2(config.GetEnv("QUEUE2","error_queue"))
	default:
		log.Fatalf("Invalid Argument: %s. Use 'producer' or 'consumer'.",os.Args[1])
		
	}
}