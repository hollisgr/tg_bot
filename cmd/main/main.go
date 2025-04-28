package main

import (
	"context"
	"os"
	"os/signal"
	"tg_bot/internal/bot/handlers"
	"tg_bot/internal/cfg"

	"github.com/go-telegram/bot"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(handlers.Help),
	}

	b, err := bot.New(cfg.GetToken(), opts...)
	if nil != err {
		panic(err)
	}

	handlers.Register(b)
	b.Start(ctx)
}
