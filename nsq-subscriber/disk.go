package main

import (
	"io/ioutil"
	"log"
	"path"

	"github.com/nsqio/go-nsq"
)

var pathToFile = path.Join("pages", "nsq.txt")

// DiskHandler ...
type DiskHandler struct {
	topic string
}

// HandleMessage (for `go-nsq`'s `Consumer`)
func (dh *DiskHandler) HandleMessage(m *nsq.Message) error {
	data := dh.topic + " | " + string(m.Body)
	err := ioutil.WriteFile(pathToFile, []byte(data), 0600)
	if err != nil {
		log.Fatalf("ERROR: failed to write file - %s", err)
	}

	return nil
}
