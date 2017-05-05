// Experimental Kafka consumer written in Go. Code has mostly been taken from
// sarama's consumer example (github.com/Shopify/sarama).
package main

import (
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/Shopify/sarama"
)

var brokerList = []string{"kafka1:9091", "kafka2:9092", "kafka3:9093"}

func main() {
	time.Sleep(30 * time.Second)

	consumer, err := sarama.NewConsumer(brokerList, nil)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := consumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	partitionConsumer, err := consumer.ConsumePartition("test", 0, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	// Trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	log.Println("Starting consumer loop")

	consumed := 0
ConsumerLoop:
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			log.Printf("Consumed message offset %d\n", msg.Offset)
			log.Printf("Key: %s, Value: %s\n", msg.Key, msg.Value)
			consumed++
		case <-signals:
			break ConsumerLoop
		}
	}

	log.Printf("Consumed: %d\n", consumed)
}
