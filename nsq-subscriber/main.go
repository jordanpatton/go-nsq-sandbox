package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nsqio/go-nsq"
	// "github.com/nsqio/nsq/internal/app"
	// "github.com/nsqio/nsq/internal/version"
)

const (
	nsqdTCPAddress        = "127.0.0.1:4150"
	nsqlookupdHTTPAddress = "127.0.0.1:4161"
	topic                 = "test"
)

var (
	channel     = ""
	maxInFlight = 1
)

func main() {
	cfg := nsq.NewConfig()
	cfg.MaxInFlight = maxInFlight
	// cfg.UserAgent = fmt.Sprintf("nsq_tail/%s go-nsq/%s", version.Binary, nsq.VERSION)

	if channel == "" {
		rand.Seed(time.Now().UnixNano())
		channel = fmt.Sprintf("tail%06d#ephemeral", rand.Int()%999999)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	consumer, err := nsq.NewConsumer(topic, channel, cfg)
	if err != nil {
		log.Fatal(err)
	}

	consumer.AddHandler(&TailHandler{topic: topic})

	err = consumer.ConnectToNSQD(nsqdTCPAddress)
	if err != nil {
		log.Fatal(err)
	}

	err = consumer.ConnectToNSQLookupd(nsqlookupdHTTPAddress)
	if err != nil {
		log.Fatal(err)
	}

	<-sigChan

	consumer.Stop()
	<-consumer.StopChan
}
