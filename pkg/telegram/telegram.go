package telegram

import (
	"fmt"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"main.go/pkg/parse"
)

type Bot struct {
	bot *tgbotapi.BotAPI
}

func NewBot(bot *tgbotapi.BotAPI) *Bot {
	return &Bot{bot: bot}
}

func (b *Bot) Start() error {
	log.Printf("Authorized on account %s", b.bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates:= b.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // If we got a message
			continue
		}
		switch update.Message.Command(){
		case "start":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет!, я бот для статистике NHL")
			b.bot.Send(msg)
		case "parse":
			url := "https://www.championat.com/hockey/_nhl/tournament/5918/calendar/"
			data := parse.FetchDataParse(url)
			response := "результаты\n" + fmt.Sprint("%v", data)
			messages := splitMessage(response, 4096)
			for _, msgTxt := range messages {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgTxt) 
				b.bot.Send(msg)
			}
		default:
			msg:= tgbotapi.NewMessage(update.Message.Chat.ID, "Извините, я не знаю такой команды.")
			b.bot.Send(msg)
		}
	}
	return nil
}

func splitMessage(text string, limit int) []string{
	var chunks []string

	for len(text) > limit {
		chunk := text[:limit]
		lastSpace := strings.LastIndex(chunk, " ")
		if lastSpace != 1 {
			chunk = chunk[:lastSpace]
		}
		chunks = append(chunks, chunk)
		text = text[len(chunk):]
	}
	chunks = append(chunks, text)
	return chunks
}