package main

import (
	"log"
	"strconv"

	"github.com/nsqio/go-nsq"
)

func producer(v int) error {
	config := nsq.NewConfig()
	p, err := nsq.NewProducer("127.0.0.1:4150", config)
	if err != nil {
		log.Panic(err)
		return err
	}

	val := strconv.Itoa(v)
	err = p.Publish("visitor_count", []byte(val))
	if err != nil {
		log.Panic(err)
		return err
	}

	return nil
}
