package main

import (
	"context"
	"github.com/jessevdk/go-flags"
	"goBotImages/internal/bot"
	"goBotImages/internal/config"
	"goBotImages/internal/random_image"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Parse opts and config
	var opts struct {
		ConfigFile string `short:"c" long:"config-file" description:"Config file name" required:"true"`
	}
	_, err := flags.ParseArgs(&opts, os.Args)
	if err != nil {
		log.Fatalf("CLI args error: %v", err)
	}
	c := config.ParseConfig(opts.ConfigFile)

	// Bind context & signals
	ctx, cancel := context.WithCancel(context.Background())
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		log.Printf("Received %s. Terminating", sig)
		cancel()
	}()

	// Start service
	api := random_image.BuildRandomImageApi(c.RandomImageUrl)
	botObj, err := bot.New(c.Token, api)
	if err != nil {
		log.Fatal(err)
	}
	botObj.StartBotPolling(ctx, 10)
}
