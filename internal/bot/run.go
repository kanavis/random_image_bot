package run

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"goBotImages/internal/random_image"
	"log"
)

func StartBotPolling(token string, api *random_image.RandomImageApi) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		go UpdateHandler(bot, update, api)
	}
}
