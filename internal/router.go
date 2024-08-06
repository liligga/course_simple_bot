package internal

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	handlers "github.com/liligga/hw_tg_bot/internal/handlers"
	bot "github.com/liligga/hw_tg_bot/pkg/bot"
)

func onStartup(dpp *bot.Dispatcher) {
	handlers.SetMyCommandsAdmin(&dpp.Bot)
	if dpp.Bot.HasFeature(handlers.ReviewFeature) {
		handlers.SetMyCommandsProducts(&dpp.Bot)
	}
	dpp.Bot.ToggleFeature(handlers.InlineButtonsMenuFeature)
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

	myDispatcher.AddHandlers(handlers.AddRandomRecipeHandlers()...)
	myDispatcher.AddHandlers(handlers.AddCategoriesHandlers()...)
	myDispatcher.AddHandlers(handlers.AddMenuHandlers()...)
	myDispatcher.AddHandlers(handlers.AddHomeworkHandlers()...)

	// Admin функционал
	myDispatcher.AddHandlers(handlers.AddAdminHandlers()...)

	// В самом конце !!!
	myDispatcher.AddHandler([2]interface{}{handlers.EmptyFilter, handlers.EchoHandler})

	// myDispatcher.Handlers = handlers

	return myDispatcher
}
