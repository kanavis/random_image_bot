package main

import (
	"fmt"
	"github.com/akamensky/argparse"
	"goBotImages/internal/bot"
	"goBotImages/internal/config"
	"goBotImages/internal/random_image"
	"log"
	"os"
)

func main() {
	parser := argparse.NewParser("bot", "Run images bot")
	configFile := parser.String("c", "config-file", &argparse.Options{Required: true, Help: "Config file", Default: "./config.yaml"})
	if err := parser.Parse(os.Args); err != nil {
		fmt.Print(parser.Usage(err))
	}
	c := config.ParseConfig(*configFile)
	api := random_image.BuildRandomImageApi(c.RandomImageUrl)
	botObj, err := bot.New(c.Token, api)
	if err != nil {
		log.Fatal(err)
	}
	botObj.StartBotPolling(10)
}
