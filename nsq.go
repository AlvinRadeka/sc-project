package main

import (
	"log"
	"sync"

	"github.com/bitly/go-nsq"
)

func producer() {
	config := nsq.NewConfig()
	w, _ := nsq.NewProducer("127.0.0.1:32787", config)

	err := w.Publish("write_test", []byte("test"))
	if err != nil {
		log.Panic("Could not connect")
	}

	w.Stop()
}

func consumer() {

	wg := &sync.WaitGroup{}
	wg.Add(1)

	config := nsq.NewConfig()
	q, _ := nsq.NewConsumer("write_test", "ch", config)
	q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		log.Printf("Got a message: %v", *message)
		wg.Done()
		return nil
	}))
	err := q.ConnectToNSQLookupd("127.0.0.1:32784")
	if err != nil {
		log.Panic("Could not connect")
	}
	wg.Wait()

}
