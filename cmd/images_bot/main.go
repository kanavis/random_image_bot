package main

import (
	"fmt"
	"github.com/akamensky/argparse"
	"goBotImages/internal/bot"
	"goBotImages/internal/config"
	"goBotImages/internal/random_image"
	"os"
)

func main() {
	parser := argparse.NewParser("bot", "Run images bot")
	configFile := parser.String("c", "config-file", &argparse.Options{Required: true, Help: "Config file", Default: "./config.yaml"})
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}
	c := config.ParseConfig(*configFile)
	api := random_image.BuildRandomImageApi(c.RandomImageUrl)
	botObj := bot.CreateBot(c.Token, api)
	botObj.StartBotPolling()
}
