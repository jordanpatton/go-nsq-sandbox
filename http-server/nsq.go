package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nsqio/go-nsq"
	// "github.com/nsqio/nsq/internal/version"
)

const nsqdTCPAddress = "127.0.0.1:4150"

// PublishToNsq ...
func PublishToNsq(topic string, message string) {
	cfg := nsq.NewConfig()
	// cfg.UserAgent = fmt.Sprintf("to_nsq/%s go-nsq/%s", version.Binary, nsq.VERSION)

	stopChan := make(chan bool)
	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)

	producer, err := nsq.NewProducer(nsqdTCPAddress, cfg)
	if err != nil {
		log.Fatalf("failed to create nsq.Producer - %s", err)
	}

	go func() {
		err := producer.Publish(topic, []byte(message))
		if err != nil {
			log.Fatal(err)
		}
		close(stopChan)
	}()

	select {
	case <-termChan:
	case <-stopChan:
	}

	producer.Stop()
}
