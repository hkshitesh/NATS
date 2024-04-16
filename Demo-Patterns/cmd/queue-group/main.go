package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)

	if err != nil {
		log.Fatalln()
	}

	defer nc.Close()

	nc.QueueSubscribe("demo_queue", "my_queue", func(msg *nats.Msg) {
		log.Println("Subscriber 1: ", string(msg.Data))
	})
	nc.QueueSubscribe("demo_queue", "my_queue", func(msg *nats.Msg) {
		log.Println("Subscriber 2: ", string(msg.Data))
	})
	nc.QueueSubscribe("demo_queue", "my_queue", func(msg *nats.Msg) {
		log.Println("Subscriber 3: ", string(msg.Data))
	})

	for i := 1; i <= 3; i++ {
		message := fmt.Sprintf("Message %d", i)
		if err := nc.Publish("demo_queue", []byte(message)); err != nil {
			log.Fatalln(err)
		}
	}
	time.Sleep((2 * time.Second))
}
