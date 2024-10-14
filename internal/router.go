package internal

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	handlers "github.com/liligga/hw_tg_bot/internal/handlers"
	"github.com/liligga/hw_tg_bot/internal/handlers/echo"
	"github.com/liligga/hw_tg_bot/internal/handlers/homework"
	randomrecipe "github.com/liligga/hw_tg_bot/internal/handlers/random_recipe"
	restaurantreview "github.com/liligga/hw_tg_bot/internal/handlers/restaurant_review"
	startinlinebuttons "github.com/liligga/hw_tg_bot/internal/handlers/start_inline_buttons"
	startwithmenu "github.com/liligga/hw_tg_bot/internal/handlers/start_with_menu"
	bot "github.com/liligga/hw_tg_bot/pkg/bot"
)

func onStartup(dpp *bot.Dispatcher) {
	// handlers.SetMyCommandsAdmin(&dpp.Bot)
	// if dpp.Bot.HasFeature(handlers.ReviewFeature) {
	// 	handlers.SetMyCommandsProducts(&dpp.Bot)
	// }
	// handlers.SetMyCommandsProducts(&dpp.Bot)
	// dpp.Bot.ToggleFeature(handlers.InlineButtonsMenuFeature)
	dpp.Bot.ToggleFeature(handlers.ReviewFeature)
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

	myDispatcher.AddHandlers(randomrecipe.AddRandomRecipeHandlers()...)
	myDispatcher.AddHandlers(startinlinebuttons.AddCategoriesHandlers()...)
	myDispatcher.AddHandlers(startwithmenu.AddMenuHandlers()...)
	myDispatcher.AddHandlers(restaurantreview.AddReviewFSMHandlers()...)
	myDispatcher.AddHandlers(homework.AddHomeworkHandlers()...)

	// Admin функционал
	myDispatcher.AddHandlers(handlers.AddAdminHandlers()...)

	// В самом конце !!!
	myDispatcher.AddHandler([2]interface{}{echo.EmptyFilter, echo.EchoHandler})

	// myDispatcher.Handlers = handlers

	return myDispatcher
}
