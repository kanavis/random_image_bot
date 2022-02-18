package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"goBotImages/internal/random_image"
	"log"
)

type RandomImageBot struct {
	botApi         *tgbotapi.BotAPI
	randomImageApi *random_image.RandomImageApi
}

func CreateBot(token string, randomImageApi *random_image.RandomImageApi) *RandomImageBot {
	botApi, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}
	botApi.Debug = true
	log.Printf("Authorized on account %s", botApi.Self.UserName)
	return &RandomImageBot{botApi: botApi, randomImageApi: randomImageApi}
}

func (bot *RandomImageBot) StartBotPolling() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.botApi.GetUpdatesChan(u)
	for update := range updates {
		go bot.UpdateHandler(update)
	}
}

func (bot *RandomImageBot) UpdateHandler(update tgbotapi.Update) {
	log.Printf("Update %v", update.Message)

	photoBytes, err := bot.randomImageApi.GetRandomPhoto()
	if err != nil {
		log.Printf("Error getting random photo: %v", err)
		return
	}
	photo := tgbotapi.FileBytes{
		Name:  "Picture",
		Bytes: photoBytes,
	}
	msg := tgbotapi.NewPhoto(update.Message.Chat.ID, photo)
	msg.ReplyToMessageID = update.Message.MessageID
	_, err = bot.botApi.Send(msg)
	if err != nil {
		log.Printf("Error sending message %v", err)
	}
}
