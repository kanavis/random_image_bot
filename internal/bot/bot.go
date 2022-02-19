package bot

import (
	"context"
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

func (bot *RandomImageBot) StartBotPolling(ctx context.Context, workers int) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updatesChan := bot.botApi.GetUpdatesChan(u)
	wg := &sync.WaitGroup{}
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go bot.pollUpdatesChannel(ctx, wg, updatesChan)
	}
	wg.Wait()
}

func (bot *RandomImageBot) pollUpdatesChannel(ctx context.Context, wg *sync.WaitGroup, updatesChan tgbotapi.UpdatesChannel) {
	defer wg.Done()
	for {
		select {
		case update := <-updatesChan:
			bot.UpdateHandler(update)
		case <-ctx.Done():
			return
		}
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
