package main

import (
	"flag"
	"jagerTgBot/clients/telegram"
	"log"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {
	tgClient := telegram.New(tgBotHost, mustToken())

	// fetcher = fetcher.New()

	// processor = processor.New()

	// consumer.Start(fetcher, processor)
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
