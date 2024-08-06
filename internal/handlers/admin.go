package handlers

import (
	"fmt"
	"slices"
	"strings"

	bot "github.com/liligga/hw_tg_bot/pkg/bot"
)

const (
	SimplePictureAndStartFeature bot.Feature = "SimplePictureAndStart"
	RandomRecipeFeature          bot.Feature = "RandomRecipe"
	InlineButtonsMenuFeature     bot.Feature = "InlineButtonsMenu"
	DishCategoriesFeature        bot.Feature = "DishCategories"
	ReviewFeature                bot.Feature = "Review"
	DishesPicturesFeature        bot.Feature = "DishesPictures"
	TestWorkFeature              bot.Feature = "TestWork"
)

var BotFeatures = []bot.Feature{
	SimplePictureAndStartFeature,
	RandomRecipeFeature,
	InlineButtonsMenuFeature,
	DishCategoriesFeature,
	ReviewFeature,
	DishesPicturesFeature,
	TestWorkFeature,
}

var TextFeatures = map[string]bot.Feature{
	"Простая картинка и старт":        SimplePictureAndStartFeature,
	"Счетчик и рандомные рецепты":     RandomRecipeFeature,
	"Меню со встроенными кнопками":    InlineButtonsMenuFeature,
	"Категории блюд, меню при старте": DishCategoriesFeature,
	"Отзыв в виде диалога":            ReviewFeature,
	"Блюда с картиками из БД":         DishesPicturesFeature,
	"Пример работы бота для теста":    TestWorkFeature,
}

func AdminCommandFilter(update bot.Update, theBot *bot.Bot) bool {
	isAdmin := update.Message.Chat.ID == 243154734
	command := update.Message.Text == "/admin"

	fmt.Println(isAdmin, command)
	fmt.Println("Admin command: ", update.Message.Text)

	return isAdmin && command
}

func renderFeatureButton(feature bot.Feature, theBot *bot.Bot) string {
	// featureText := map[bot.Feature]string{
	// 	SimplePictureAndStartFeature: "Простая картинка и старт",
	// 	RandomRecipeFeature:          "Счетчик и рандомные рецепты",
	// 	DishCategoriesFeature:        "Категории блюд, меню при старте",
	// 	ReviewFeature:                "Отзыв в виде диалога",
	// 	DishesPicturesFeature:        "Блюда с картиками из БД",
	// 	TestWorkFeature:              "Пример работы бота для теста",
	// }

	var theFeature string
	for text, ft := range TextFeatures {
		if ft == feature {
			theFeature = text
			break
		}
	}

	if theBot.HasFeature(feature) {
		return fmt.Sprintf("✅ %s", theFeature)
	}

	return fmt.Sprintf("❌ %s", theFeature)
}

func constructButtons(theBot *bot.Bot) [][]bot.KeyboardButton {
	var kb [][]bot.KeyboardButton
	for _, feature := range TextFeatures {
		kb = append(kb, []bot.KeyboardButton{
			{Text: renderFeatureButton(feature, theBot)},
		})
	}
	return kb
}

func AdminHandler(update bot.Update, theBot *bot.Bot) {
	answer := bot.NewTextAnswer(
		update.Message.Chat.ID,
		"Выберите фичу, которую хотите добавить/удалить",
	)

	kb := bot.ReplyKeyboardMarkup{
		Keyboard:       constructButtons(theBot),
		ResizeKeyboard: true,
	}
	answer.AddReplyMarkup(&kb)

	theBot.SendMessage(answer)
}

// func AdminHandler(update bot.Update, theBot *bot.Bot) {
// 	poll := bot.NewPollAnswer(
// 		update.Message.Chat.ID,
// 		"Выберите фичу, которую хотите активировать",
// 		[]string{
// 			"Простая картинка и старт",
// 			"Счетчик и рандомные рецепты",
// 			"Категории блюд, меню при старте",
// 			"Отзыв в виде диалога",
// 			"Блюда с картиками из БД",
// 			"Пример работы бота для теста",
// 		},
// 	)

// 	theBot.SendPoll(poll)
// }

func AdminButtonsFilter(update bot.Update, theBot *bot.Bot) bool {
	isAdmin := update.Message.Chat.ID == 243154734
	areButtonsClicked := slices.Contains(
		[]string{
			"Простая картинка и старт",
			"Счетчик и рандомные рецепты",
			"Меню со встроенными кнопками",
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
	// kb := bot.ReplyKeyboardRemove{RemoveKeyboard: true}
	var answer *bot.TextAnswer

	text := strings.Trim(update.Message.Text, "❌✅ ")
	feature := TextFeatures[text]
	// if theBot.HasFeature(feature) {
	// 	theBot.DeleteFeature(feature)
	// 	answer = bot.NewTextAnswer(update.Message.Chat.ID, "Фича удалена")
	// } else {
	// 	theBot.AddFeature(feature)
	// 	answer = bot.NewTextAnswer(update.Message.Chat.ID, "Фича добавлена")
	// }
	added := theBot.ToggleFeature(feature)
	if added {
		answer = bot.NewTextAnswer(update.Message.Chat.ID, "Фича добавлена")
	} else {
		answer = bot.NewTextAnswer(update.Message.Chat.ID, "Фича удалена")
	}

	kb := bot.ReplyKeyboardMarkup{
		Keyboard:       constructButtons(theBot),
		ResizeKeyboard: true,
	}
	answer.AddReplyMarkup(&kb)
	// answer.AddReplyRemove(&kb)
	theBot.SendMessage(answer)
}

func AdminQuitFilter(update bot.Update, theBot *bot.Bot) bool {
	isAdmin := update.Message.Chat.ID == 243154734
	return isAdmin && update.Message.Text == "/quit"
}

func AdminQuitHandler(update bot.Update, theBot *bot.Bot) {
	kb := bot.ReplyKeyboardRemove{RemoveKeyboard: true}
	answer := bot.NewTextAnswer(update.Message.Chat.ID, "Всего хорошего")
	answer.AddReplyRemove(&kb)
	theBot.SendMessage(answer)
}

func AddAdminHandlers() [][2]interface{} {
	return [][2]interface{}{
		{AdminCommandFilter, AdminHandler},
		{AdminButtonsFilter, AdminButtonHandler},
		{AdminQuitFilter, AdminQuitHandler},
	}
}
