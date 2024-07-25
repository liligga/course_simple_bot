package handlers

import (
	bot "github.com/liligga/hw_tg_bot/bot"
)

func EmptyFilter(update bot.Update, theBot *bot.Bot) bool {
	return true
}

func EchoHandler(update bot.Update, theBot *bot.Bot) {
	answer := bot.NewTextAnswer(
		update.Message.Chat.ID,
		"Простите, я не понимаю",
	)

	theBot.SendMessage(answer)
}
