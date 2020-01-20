package main;

import (
        "encoding/json"
    	"flag"
    	"fmt"
    	"github.com/streadway/amqp"
    	"io/ioutil"
    	"log"
    	"net/http"
)

var (
        address = flag.String("address", "bind host:port")
        amqpUri = flag.String("amqp", "amqp://guest:guest@127.0.0.1:5672/", "amqp uri")
)

func init() {
        flag.parse()
}

type MessageEntity struct {
        Exchange     string `json:"exchange"`
        Key          string `json:"key"`
        DeliveryMode uint8  `json:"deliverymode"`
        Priority     uint8  `json:"priority"`
        Body         string `json:"body"`
}

type ExchangeEntity struct {
	Name       string `json:"name"`
	Type       string `json:"type"`
	Durable    bool   `json:"durable"`
	AutoDelete bool   `json:"autodelete"`
	NoWait     bool   `json:"nowait"`
}

type QueueEntity struct {
	Name       string `json:"name"`
	Durable    bool   `json:"durable"`
	AutoDelete bool   `json:"autodelete"`
	Exclusive  bool   `json:"exclusive"`
	NoWait     bool   `json:"nowait"`
}

type QueueBindEntity struct {
	Queue    string   `json:"queue"`
	Exchange string   `json:"exchange"`
	NoWait   bool     `json:"nowait"`
	Keys     []string `json:"keys"` // bind/routing keys
}

// RabbitMQ Operate Wrapper
type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	done    chan error
}

// TODO::Organize http requests to the amqp server
func main() {
        http.HandleFunc("/exchange", ExchangeHandler)
        http.HandleFunc("/queue/bind", QueueBindHandler)
        http.HandleFunc("/queue", QueueHandler)
        http.HandleFunc("/publish", PublishHandler)
        // TODO::make these handler functions within the main class

        // Start HTTP Server
        log.Printf("server run %s (listen %s)\n", *address, *amqpUri)
        err := http.ListenAndServe(*address, nil)

        if err != nil {
        	log.Fatal(err)
        }
}
