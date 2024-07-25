package handlers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/liligga/hw_tg_bot/bot"
)

func onStartup(dpp *bot.Dispatcher) {

	fmt.Println("Bot started")
}

func NewApp(client *http.Client) bot.Dispatcher {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	tok := os.Getenv("BOT_TOKEN")
	if tok == "" {
		fmt.Println("BOT_TOKEN is not set")
		os.Exit(1)
	}

	myDispatcher := bot.NewDispatcher(tok, client)
	myDispatcher.OnStartup(onStartup)

	myDispatcher.AddHandlers(AddMenuHandlers()...)
	myDispatcher.AddHandlers(AddHomeworkHandlers()...)

	// Admin функционал
	myDispatcher.AddHandlers(AddAdminHandlers()...)

	// В самом конце !!!
	myDispatcher.AddHandler([2]interface{}{EmptyFilter, EchoHandler})

	// myDispatcher.Handlers = handlers

	return myDispatcher
}
