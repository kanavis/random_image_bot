package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"goBotImages/internal/random_image"
	"log"
	"sync"
)

type RandomImageBot struct {
	botApi         *tgbotapi.BotAPI
	randomImageApi *random_image.RandomImageApi
}

func New(token string, randomImageApi *random_image.RandomImageApi) (*RandomImageBot, error) {
	botApi, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	botApi.Debug = true
	log.Printf("Authorized on account %s", botApi.Self.UserName)
	return &RandomImageBot{botApi: botApi, randomImageApi: randomImageApi}, nil
}

func (bot *RandomImageBot) StartBotPolling(workers int) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.botApi.GetUpdatesChan(u)
	wg := &sync.WaitGroup{}
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go bot.pollUpdatesChannel(wg, updates)
	}
	wg.Wait()
}

func (bot *RandomImageBot) pollUpdatesChannel(wg *sync.WaitGroup, updates tgbotapi.UpdatesChannel) {
	defer wg.Done()
	for update := range updates {
		bot.UpdateHandler(update)
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
	if _, err := bot.botApi.Send(msg); err != nil {
		log.Printf("Error sending message %v", err)
	}
}
