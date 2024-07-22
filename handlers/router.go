package handlers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/liligga/hw_tg_bot/bot"
)

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

	handlers := make([][2]interface{}, 0)
	// handlers = append(handlers, [2]interface{}{CommandStartFilter, StartHandler})
	// handlers = append(handlers, [2]interface{}{ButtonMenuFilter, MenuHandler})
	// handlers = append(handlers, [2]interface{}{CommandMenuFilter, MenuHandler})
	// handlers = append(handlers, [2]interface{}{CategoryFilter, CategoryHandler})
	handlers = append(handlers, [2]interface{}{SendHomeWorkFilter, StartHomeworkDialogueHandler})
	handlers = append(handlers, [2]interface{}{ProcessNameFilter, ProcessNameHandler})
	handlers = append(handlers, [2]interface{}{ProcessGroupFilter, ProcessGroupHandler})
	handlers = append(handlers, [2]interface{}{ProcessHomeWorkNumberFilter, ProcessHomeWorkNumberHandler})
	handlers = append(handlers, [2]interface{}{ProcessLinkFilter, ProcessLinkHandler})
	handlers = append(handlers, [2]interface{}{EmptyFilter, EchoHandler})

	myDispatcher.Handlers = handlers

	return myDispatcher
}
