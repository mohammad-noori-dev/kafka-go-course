package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	topic := "user-profiles-segio"

	writer := &kafka.Writer{
		Addr:  kafka.TCP("localhost:9094"),
		Topic: topic,
	}

	defer func() {
		err := writer.Close()
		if err != nil {
			log.Fatal("failed to close writer", err)
		}
	}()
	fmt.Println("Producer started. Sending message to topic: ", topic)

	// We'll send 10 messages
	for i := 0; i < 10; i++ {
		message := fmt.Sprintf("Message #%d for user", i)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		err := writer.WriteMessages(ctx, kafka.Message{
			Value: []byte(message),
		})
		if err != nil {
			log.Fatalf("failed to write message %d: %v", i, err)
		}

		fmt.Printf("Sent message: %s\n", message)

		// A small delay between messages
		time.Sleep(500 * time.Millisecond)
	}

	fmt.Println("Finished sending messages!")
}
