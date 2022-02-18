package run

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"goBotImages/internal/random_image"
	"log"
)

func UpdateHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update, api *random_image.RandomImageApi) {
	log.Printf("Update %v", update.Message)

	photoBytes, err := api.GetRandomPhoto()
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
	_, err = bot.Send(msg)
	if err != nil {
		log.Printf("Error sending message %v", err)
	}
}
