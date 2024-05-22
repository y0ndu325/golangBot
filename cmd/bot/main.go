package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"main.go/pkg/telegram"
)
const TOKEN = "token"

func main() {
	bot, err := tgbotapi.NewBotAPI(TOKEN)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	telegramBot := telegram.NewBot(bot)
	if err := telegramBot.Start();err!=nil {
		log.Fatal(err)
	}
}