package handlers

import (
	"math/rand"
	"strconv"

	bot "github.com/liligga/hw_tg_bot/bot"
)

var userSet = []int{}

// var Counter = map[int]int{}

func StartCommandFilter2(update bot.Update, theBot *bot.Bot) bool {
	if !theBot.HasFeature(bot.RandomRecipeFeature) {
		return false
	}
	return update.Message.Text == "/start"
}

func StartCommandHandler2(update bot.Update, theBot *bot.Bot) {
	userId := update.Message.From.ID
	found := false
	for _, el := range userSet {
		// fmt.Println(el)
		if el == userId {
			found = true
			break
		}
	}

	if !found {
		userSet = append(userSet, int(userId))
	}

	answer := bot.NewTextAnswer(
		update.Message.Chat.ID,
		"Привет, "+update.Message.From.FirstName+"!\n"+"Наш бот обслуживает уже "+strconv.Itoa(len(userSet))+" пользователей",
	)
	theBot.SendMessage(answer)
}

func RandomRecipeFilter(update bot.Update, theBot *bot.Bot) bool {
	if !theBot.HasFeature(bot.RandomRecipeFeature) {
		return false
	}
	return update.Message.Text == "/random_recipe"
}

func RandomRecipeHandler(update bot.Update, theBot *bot.Bot) {
	recipies := [4]string{
		"Омлет: Для приготовления омлета разбейте 3 куринных яйца и добавьте пол стакана молока, все взбейте и запеките",
		"Котлета: Для приготовления котлеты смешайте фарш с яйцом, тестом, добавьте соль и запекайте",
		"Борщ: Для приготовления борща поставьте мясо вариться на 40 минут, сделайте зажарку. Добавьте капусту и картофель в бульон. Добавьте зажарку",
		"Ризотто: В сотейнике растопите масло и обжарьте на нём измельчённый лук. Выложите рис и обжаривайте пару минут.Постепенно влейте в рис вино и горячий бульон. Затем добавьте рыбу, нарезанную небольшими кусочками, рубленый зелёный лук и шпинат и перемешайте.  Готовьте ещё немного, пока шпинат не размягчится. При необходимости добавьте в ризотто соль.",
	}
	selectedRecipe := recipies[rand.Intn(len(recipies))]
	answer := bot.NewTextAnswer(
		update.Message.Chat.ID,
		"Вот ваш рецепт:\n"+selectedRecipe,
	)
	theBot.SendMessage(answer)
}

func InfoCommandFilter(update bot.Update, theBot *bot.Bot) bool {
	if !theBot.HasFeature(bot.RandomRecipeFeature) {
		return false
	}
	return update.Message.Text == "/info"
}

func InfoCommandHandler(update bot.Update, theBot *bot.Bot) {
	username := update.Message.From.Username
	if username == "" {
		username = "У вас нет ника"
	} else {
		username = "Ваш ник: " + username
	}
	msgText := "Ваш ID: " + strconv.FormatInt(int64((update.Message.From.ID)), 10) + "\nВаше имя: " + update.Message.From.FirstName + "\n" + username
	answer := bot.NewTextAnswer(
		update.Message.Chat.ID,
		msgText,
	)

	theBot.SendMessage(answer)
}

func AddRandomRecipeHandlers() [][2]interface{} {
	handlers := [][2]interface{}{
		{StartCommandFilter2, StartCommandHandler2},
		{RandomRecipeFilter, RandomRecipeHandler},
		{InfoCommandFilter, InfoCommandHandler},
	}

	return handlers
}
