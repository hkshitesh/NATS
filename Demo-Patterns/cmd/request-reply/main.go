package main

import (
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)

	if err != nil {
		log.Fatalln(err)
	}
	defer nc.Close()

	nc.Subscribe("demo_reply", func(msg *nats.Msg) {
		log.Println("Request Received", string(msg.Data))
		msg.Respond([]byte("Here you go!"))
	})

	reply, err := nc.Request("demo_reply", []byte("Give me data"), 4*time.Second)

	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Got Reply:", string(reply.Data))

}
