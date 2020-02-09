package main

import (
	"log"
	"os"

	"github.com/nsqio/go-nsq"
)

// TailHandler ...
type TailHandler struct {
	topic string
}

// HandleMessage (for `go-nsq`'s `Consumer`)
func (th *TailHandler) HandleMessage(m *nsq.Message) error {
	// print topic
	_, err := os.Stdout.WriteString(th.topic)
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

	return nil
}
