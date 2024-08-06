package handlers

import (
	"slices"
	"strings"

	bot "github.com/liligga/hw_tg_bot/pkg/bot"
)

type DialogueStep string

const (
	Name           DialogueStep = "name"
	Group          DialogueStep = "group"
	HomeWorkNumber DialogueStep = "homework"
	Link           DialogueStep = "link"
)

func SendHomeWorkFilter(update bot.Update, theBot *bot.Bot) bool {
	if !theBot.HasFeature(TestWorkFeature) {
		return false
	}
	return update.Message.Text == "/send_hmw"
}

func StartHomeworkDialogueHandler(update bot.Update, theBot *bot.Bot) {
	answer := bot.NewTextAnswer(
		update.Message.Chat.ID,
		"Как вас зовут?",
	)
	userContext := theBot.GetUserContext(
		update.Message.From.ID,
		update.Message.Chat.ID,
	)
	userContext.CurrentState = string(Name)

	theBot.SendMessage(answer)
}

func ProcessNameFilter(update bot.Update, theBot *bot.Bot) bool {
	if !theBot.HasFeature(TestWorkFeature) {
		return false
	}
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

	reply_keyboard := bot.ReplyKeyboardMarkup{
		ResizeKeyboard: true,
		Keyboard: [][]bot.KeyboardButton{
			{
				{Text: "43-1"},
				{Text: "43-2"},
				{Text: "42-1"},
				{Text: "42-2"},
			},
		},
		InputFieldPlaceholder: "Группа",
	}
	answer.AddReplyMarkup(&reply_keyboard)

	userContext := theBot.GetUserContext(
		update.Message.From.ID,
		update.Message.Chat.ID,
	)
	userContext.CurrentState = string(Group)
	theBot.SendMessage(answer)
}

func ProcessGroupFilter(update bot.Update, theBot *bot.Bot) bool {
	if !theBot.HasFeature(TestWorkFeature) {
		return false
	}
	userID := update.Message.From.ID
	chatID := update.Message.Chat.ID

	userContext := theBot.GetUserContext(userID, chatID)
	state := userContext.CurrentState == string(Group)
	msgText := slices.Contains(
		[]string{"43-1", "43-2", "42-1", "42-2"},
		update.Message.Text,
	)

	return state && msgText
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

	reply_keyboard := bot.ReplyKeyboardMarkup{
		ResizeKeyboard: true,
		Keyboard: [][]bot.KeyboardButton{
			{
				{Text: "1"},
				{Text: "2"},
				{Text: "3"},
				{Text: "4"},
			},
			{
				{Text: "5"},
				{Text: "6"},
				{Text: "7"},
				{Text: "8"},
			},
		},
		InputFieldPlaceholder: "Выберите номер домашнего задания",
	}
	answer.AddReplyMarkup(&reply_keyboard)

	theBot.SendMessage(answer)
}

func ProcessHomeWorkNumberFilter(update bot.Update, theBot *bot.Bot) bool {
	if !theBot.HasFeature(TestWorkFeature) {
		return false
	}
	userID := update.Message.From.ID
	chatID := update.Message.Chat.ID

	userContext := theBot.GetUserContext(userID, chatID)
	state := userContext.CurrentState == string(HomeWorkNumber)
	msgText := slices.Contains(
		[]string{"1", "2", "3", "4", "5", "6", "7", "8"},
		update.Message.Text,
	)

	return state && msgText
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

	removeKb := bot.ReplyKeyboardRemove{RemoveKeyboard: true}
	answer.AddReplyRemove(&removeKb)

	theBot.SendMessage(answer)
}

func ProcessLinkFilter(update bot.Update, theBot *bot.Bot) bool {
	if !theBot.HasFeature(TestWorkFeature) {
		return false
	}
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

func AddHomeworkHandlers() [][2]interface{} {
	handlers := [][2]interface{}{
		{SendHomeWorkFilter, StartHomeworkDialogueHandler},
		{ProcessNameFilter, ProcessNameHandler},
		{ProcessGroupFilter, ProcessGroupHandler},
		{ProcessHomeWorkNumberFilter, ProcessHomeWorkNumberHandler},
		{ProcessLinkFilter, ProcessLinkHandler},
	}

	return handlers
}
