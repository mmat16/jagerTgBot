package main

import (
	"flag"
	tgClient "jagerTgBot/clients/telegram"
	event_consumer "jagerTgBot/consumer/event-consumer"
	"jagerTgBot/events/telegram"
	"jagerTgBot/storage/files"
	"log"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "storage"
	batchSize   = 100
)

func main() {
	//tgClient := telegram.New(tgBotHost, mustToken())
	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		files.New(storagePath),
	)
	log.Print("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)
	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}

func mustToken() string {
	token := flag.String(
		"token-bot-token",
		"",
		"enter bot token to access telegram",
	)
	flag.Parse()

	if *token == "" {
		log.Fatal("token were not specified")
	}
	return *token
}
