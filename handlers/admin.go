package handlers

import (
	"fmt"
	"slices"
	"strings"

	bot "github.com/liligga/hw_tg_bot/bot"
)

var BotFeatures = []bot.Feature{
	bot.SimplePictureAndStartFeature,
	bot.RandomRecipeFeature,
	bot.DishCategoriesFeature,
	bot.ReviewFeature,
	bot.DishesPicturesFeature,
	bot.TestWorkFeature,
}

var TextFeatures = map[string]bot.Feature{
	"Простая картинка и старт":        bot.SimplePictureAndStartFeature,
	"Счетчик и рандомные рецепты":     bot.RandomRecipeFeature,
	"Категории блюд, меню при старте": bot.DishCategoriesFeature,
	"Отзыв в виде диалога":            bot.ReviewFeature,
	"Блюда с картиками из БД":         bot.DishesPicturesFeature,
	"Пример работы бота для теста":    bot.TestWorkFeature,
}

func AdminCommandFilter(update bot.Update, theBot *bot.Bot) bool {
	isAdmin := update.Message.Chat.ID == 243154734
	command := update.Message.Text == "/admin"

	fmt.Println(isAdmin, command)
	fmt.Println("Admin command: ", update.Message.Text)

	return isAdmin && command
}

func renderFeatureButton(feature bot.Feature, theBot *bot.Bot) string {
	featureText := map[bot.Feature]string{
		bot.SimplePictureAndStartFeature: "Простая картинка и старт",
		bot.RandomRecipeFeature:          "Счетчик и рандомные рецепты",
		bot.DishCategoriesFeature:        "Категории блюд, меню при старте",
		bot.ReviewFeature:                "Отзыв в виде диалога",
		bot.DishesPicturesFeature:        "Блюда с картиками из БД",
		bot.TestWorkFeature:              "Пример работы бота для теста",
	}
	if theBot.HasFeature(feature) {
		return fmt.Sprintf("✅ %s", featureText[feature])
	}
	return fmt.Sprintf("❌ %s", featureText[feature])
}

func constructButtons(theBot *bot.Bot) [][]bot.KeyboardButton {
	var kb [][]bot.KeyboardButton
	for _, feature := range BotFeatures {
		kb = append(kb, []bot.KeyboardButton{
			{Text: renderFeatureButton(feature, theBot)},
		})
	}
	return kb
}

func AdminHandler(update bot.Update, theBot *bot.Bot) {
	answer := bot.NewTextAnswer(
		update.Message.Chat.ID,
		"Выберите фичу, которую хотите активировать",
	)

	kb := bot.ReplyKeyboardMarkup{
		Keyboard:       constructButtons(theBot),
		ResizeKeyboard: true,
	}
	answer.AddReplyMarkup(&kb)

	theBot.SendMessage(answer)
}

func AdminButtonsFilter(update bot.Update, theBot *bot.Bot) bool {
	isAdmin := update.Message.Chat.ID == 243154734
	areButtonsClicked := slices.Contains(
		[]string{
			"Простая картинка и старт",
			"Счетчик и рандомные рецепты",
			"Категории блюд, меню при старте",
			"Отзыв в виде диалога",
			"Блюда с картиками из БД",
			"Пример работы бота для теста",
		},
		strings.Trim(update.Message.Text, "❌✅ "),
	)
	return isAdmin && areButtonsClicked
}

func AdminButtonHandler(update bot.Update, theBot *bot.Bot) {
	kb := bot.ReplyKeyboardRemove{RemoveKeyboard: true}
	var answer *bot.TextAnswer

	text := strings.Trim(update.Message.Text, "❌✅ ")
	feature := TextFeatures[text]
	if theBot.HasFeature(feature) {
		theBot.DeleteFeature(feature)
		answer = bot.NewTextAnswer(update.Message.Chat.ID, "Фича удалена")
	} else {
		theBot.AddFeature(feature)
		answer = bot.NewTextAnswer(update.Message.Chat.ID, "Фича добавлена")
	}

	answer.AddReplyRemove(&kb)
	theBot.SendMessage(answer)
}

func AddAdminHandlers() [][2]interface{} {
	return [][2]interface{}{
		{AdminCommandFilter, AdminHandler},
		{AdminButtonsFilter, AdminButtonHandler},
	}
}
