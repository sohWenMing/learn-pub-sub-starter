package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	ampq "github.com/rabbitmq/amqp091-go"
)

const ampqConnectionString = "amqp://guest:guest@localhost:5672/"

func main() {
	conn, err := ampq.Dial(ampqConnectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	sigChan := make(chan os.Signal, 1)

	go func(sigChan chan<- os.Signal) {
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	}(sigChan)
	fmt.Println("Successfully connected to the AMPQ server")
	<-sigChan
	fmt.Println("System is shutting down ...")

}
