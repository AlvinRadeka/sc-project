package main

import (
	"log"
	"strconv"
	"sync"

	"github.com/bitly/go-nsq"
	"github.com/garyburd/redigo/redis"
	nsqio "github.com/nsqio/go-nsq"
)

func saveCount(msg string) error {
	val, err := strconv.Atoi(msg)
	if err != nil {
		log.Fatalln("Error converting message to int! Please check your message")
		log.Fatalln("Message received: ", msg)
		return err
	}

	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatalln(err)
		return err
	}
	defer conn.Close()

	_, err = conn.Do("SET", "visitor_count", val)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	return nil
}

func consumer() error {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	decodeConfig := nsq.NewConfig()
	c, err := nsq.NewConsumer("visitor_count", "channel_one", decodeConfig)
	if err != nil {
		log.Panic("Could not create consumer!")
		return err
	}

	c.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		msg := string(message.Body)
		log.Println("NSQ message received: ", msg)
		saveCount(msg)
		return nil
	}))

	err = c.ConnectToNSQD("127.0.0.1:4150")
	if err != nil {
		log.Panic("Could not connect")
		return err
	}

	wg.Wait()

	return nil
}

func producer(v int) error {
	config := nsqio.NewConfig()
	p, err := nsqio.NewProducer("127.0.0.1:4150", config)
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
