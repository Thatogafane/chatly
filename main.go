package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/ably/ably-go/ably"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	username := os.Args[1]

	client, err := ably.NewRealtime(
		ably.WithKey(os.Getenv("ABLY_API_KEY")),
		ably.WithClientID(username),
		ably.WithEchoMessages(false),
	)

	if err != nil {
		log.Fatal(err)
	}

	topic := client.Channels.Get("Fun Golang")

	_, _ = topic.SubscribeAll(context.Background(), func(message *ably.Message) {
		fmt.Printf("Recieved message from %v: '%v\n", message.ClientID, message.Data)
	})

	reader := bufio.NewReader(os.Stdin)

	for {
		text, _ := reader.ReadString('\n')

		text = strings.ReplaceAll(text, "\n", "")

		err = topic.Publish(context.Background(), "Fun Golang", text)
		if err != nil {
			err = fmt.Errorf("publishing to topic: %w", err)
			fmt.Println(err)
		} else {
			fmt.Printf("Published message: %v\n", text)git
		}
	}
}
