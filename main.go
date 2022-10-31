package main

import (
	"flag"
	"log"
	event_consumer "read-adviser-bot/consumer/event-consumer"

	tgClient "read-adviser-bot/clients/telegram"
	"read-adviser-bot/events/telegram"
	"read-adviser-bot/storage/files"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "storage"
	batchSize   = 100
)

func main() {

	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		files.New(storagePath),
	)
	log.Print("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("server is stopped", err)
	}
}

func mustToken() string {
	token := flag.String("token-bot-token",
		"",
		"token for accest to telegram bot",
	)

	flag.Parse()
	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token

}
