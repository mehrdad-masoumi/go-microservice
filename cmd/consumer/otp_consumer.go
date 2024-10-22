package main

import (
	"fmt"
	"log"
	"mlm/config"
	rabbitmq "mlm/pkg/rabitmq"
	"time"
)

func main() {

	err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading .env file : %v ", err)
	}

	r, err := rabbitmq.Connect(config.AppConfig.Rabbitmq)
	if err != nil {
		log.Fatal(err)
	}

	msgs, err := r.Consume("otp")
	if err != nil {
		log.Fatal(err)
	}

	for msg := range msgs {
		/// send sms
		fmt.Println(string(msg.Body))
		msg.Ack(true)
		time.Sleep(1 * time.Second)
	}

}
