package main

import (
	"log"
	"os"

	"github.com/nsqio/go-nsq"
)

// TailHandler ...
type TailHandler struct {
	messagesShown int
	topicName     string
	totalMessages int
}

// HandleMessage (for `go-nsq`'s `Consumer`)
func (th *TailHandler) HandleMessage(m *nsq.Message) error {
	th.messagesShown++

	// print topic
	_, err := os.Stdout.WriteString(th.topicName)
	if err != nil {
		log.Fatalf("ERROR: failed to write to os.Stdout - %s", err)
	}
	_, err = os.Stdout.WriteString(" | ")
	if err != nil {
		log.Fatalf("ERROR: failed to write to os.Stdout - %s", err)
	}

	// print message
	_, err = os.Stdout.Write(m.Body)
	if err != nil {
		log.Fatalf("ERROR: failed to write to os.Stdout - %s", err)
	}
	_, err = os.Stdout.WriteString("\n")
	if err != nil {
		log.Fatalf("ERROR: failed to write to os.Stdout - %s", err)
	}
	if th.totalMessages > 0 && th.messagesShown >= th.totalMessages {
		os.Exit(0)
	}

	return nil
}
