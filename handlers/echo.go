package handlers

import (
	"github.com/liligga/hw_tg_bot/bot"
)


func EmptyFilter(update bot.Update) bool {
	return true
}

func EchoHandler(update bot.Update, theBot bot.Bot) {
	theBot.SendMessage(
		update.Message.Chat.ID, 
		update.Message.Text,
		nil,
		nil,
	)
}

func CommandStartFilter(update bot.Update) bool {
    return update.Message.Text == "/start"
}

func StartHandler(update bot.Update, theBot bot.Bot) {
	kb := bot.ReplyKeyboardMarkup{
		Keyboard: [][]bot.KeyboardButton{
			{
				{Text: "Фантастика"},
				{Text: "Фэнтези"},
			},
			{
				{Text: "Романтика"},
				{Text: "Детектив"},
			},
			{
				{Text: "Боевик"},
				{Text: "Хоррор"},
			},
		},
		ResizeKeyboard: true,
	}
	
	theBot.SendMessage(
		update.Message.Chat.ID, 
		"Hello, World! This is handler for /start command",
		&kb,
		nil,
	)
}

func CategoryFilter(update bot.Update) bool {
	msgText := update.Message.Text

	if msgText == "Фантастика" || msgText == "Фэнтези" || msgText == "Романтика" || msgText == "Детектив" || msgText == "Боевик" || msgText == "Хоррор" {
		return true
	}

	return false
}

func CategoryHandler(update bot.Update, theBot bot.Bot) {
	kb := bot.ReplyKeyboardRemove{
		RemoveKeyboard: true,}
	theBot.SendMessage(
		update.Message.Chat.ID,
		"Ваша категория: "+update.Message.Text,
		nil,
		&kb,
	)
}