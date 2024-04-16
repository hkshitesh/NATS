package main

import (
	"fmt"

	"github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"
)

func main() {
	opts := &server.Options{}

	ns, err := server.NewServer(opts)

	if err != nil {
		panic(err)
	}
	go ns.Start()

	nc, err := nats.Connect(ns.ClientURL())
	if err != nil {
		panic(err)
	}
	nc.Subscribe("my_new_subject", func(msg *nats.Msg) {
		fmt.Println(string(msg.Data))
		ns.Shutdown()
	})
	nc.Publish("my_new_subject", []byte("Hello embedded NATS! Hitesh"))
	ns.WaitForShutdown()

}
