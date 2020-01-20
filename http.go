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

// TODO::Organize http requests to the amqp server

