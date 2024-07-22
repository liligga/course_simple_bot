package handlers

import (
	"strings"

	bot "github.com/liligga/hw_tg_bot/bot"
)

type DialogueStep string

const (
	Name           DialogueStep = "name"
	Group          DialogueStep = "group"
	HomeWorkNumber DialogueStep = "homework"
	Link           DialogueStep = "link"
)

func SendHomeWorkFilter(update bot.Update, theBot *bot.Bot) bool {
	return update.Message.Text == "/send_hmw"
}

func StartHomeworkDialogueHandler(update bot.Update, theBot *bot.Bot) {
	answer := bot.NewTextAnswer(
		update.Message.Chat.ID,
		"Как вас зовут?",
	)
	// newDialogue := Dialogue{
	// 	CurrentStep: Name,
	// 	UserID:      update.Message.From.ID,
	// }
	userContext := theBot.GetUserContext(
		update.Message.From.ID,
		update.Message.Chat.ID,
	)
	// userContext.Data["dialogue"] = newDialogue
	userContext.CurrentState = string(Name)

	theBot.SendMessage(answer)
}

func ProcessNameFilter(update bot.Update, theBot *bot.Bot) bool {
	userId := update.Message.From.ID
	chatId := update.Message.Chat.ID

	userContext := theBot.GetUserContext(userId, chatId)

	return userContext.CurrentState == string(Name)
}

func ProcessNameHandler(update bot.Update, theBot *bot.Bot) {
	answer := bot.NewTextAnswer(
		update.Message.Chat.ID,
		"Какая у вас группа?",
	)

	userContext := theBot.GetUserContext(
		update.Message.From.ID,
		update.Message.Chat.ID,
	)
	userContext.CurrentState = string(Group)
	theBot.SendMessage(answer)
}

func ProcessGroupFilter(update bot.Update, theBot *bot.Bot) bool {
	userID := update.Message.From.ID
	chatID := update.Message.Chat.ID

	userContext := theBot.GetUserContext(userID, chatID)
	return userContext.CurrentState == string(Group)
}

func ProcessGroupHandler(update bot.Update, theBot *bot.Bot) {
	answer := bot.NewTextAnswer(
		update.Message.Chat.ID,
		"Номер домашнего задания?",
	)

	userContext := theBot.GetUserContext(
		update.Message.From.ID,
		update.Message.Chat.ID,
	)
	userContext.CurrentState = string(HomeWorkNumber)
	theBot.SendMessage(answer)
}

func ProcessHomeWorkNumberFilter(update bot.Update, theBot *bot.Bot) bool {
	userID := update.Message.From.ID
	chatID := update.Message.Chat.ID

	userContext := theBot.GetUserContext(userID, chatID)
	return userContext.CurrentState == string(HomeWorkNumber)
}

func ProcessHomeWorkNumberHandler(update bot.Update, theBot *bot.Bot) {
	answer := bot.NewTextAnswer(
		update.Message.Chat.ID,
		"Ссылка на домашнее задание?",
	)

	userContext := theBot.GetUserContext(
		update.Message.From.ID,
		update.Message.Chat.ID,
	)
	userContext.CurrentState = string(Link)
	theBot.SendMessage(answer)
}

func ProcessLinkFilter(update bot.Update, theBot *bot.Bot) bool {
	userID := update.Message.From.ID
	chatID := update.Message.Chat.ID

	userContext := theBot.GetUserContext(userID, chatID)
	correctState := userContext.CurrentState == string(Link)

	link := strings.Contains(update.Message.Text, "https://github.com/")

	return correctState && link
}

func ProcessLinkHandler(update bot.Update, theBot *bot.Bot) {
	answer := bot.NewTextAnswer(
		update.Message.Chat.ID,
		"Спасибо, мы получили ваше домашнее задание! Хорошего дня!",
	)

	theBot.DeleteUserContext(update.Message.From.ID, update.Message.Chat.ID)
	theBot.SendMessage(answer)
}
