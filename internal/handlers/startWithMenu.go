package handlers

import (
	"fmt"
	"os"

	bot "github.com/liligga/hw_tg_bot/pkg/bot"
)

func CommandStartFilter(update bot.Update, theBot *bot.Bot) bool {
	if !theBot.HasFeature(DishCategoriesFeature) {
		return false
	}
	return update.Message.Text == "/start"
}

func StartHandler(update bot.Update, theBot *bot.Bot) {
	// kb := bot.InlineKeyboardMarkup{
	// 	InlineKeyboard: [][]bot.InlineKeyboardButton{
	// 		{
	// 			{Text: "В меню", CallbackData: "menu"},
	// 		},
	// 		{
	// 			{Text: "Контакты", CallbackData: "contacts"},
	// 			{Text: "Расписание", CallbackData: "schedule"},
	// 		},
	// 		{
	// 			{Text: "Вакансии", CallbackData: "vacancies"},
	// 			{Text: "О нас", CallbackData: "about"},
	// 			{Text: "Наш сайт", URL: "https://github.com/liligga"},
	// 		},
	// 	},
	// }
	answer := bot.NewTextAnswer(
		update.Message.Chat.ID,
		"Привет, я бот ресторана 'Ресторан'. Я помогу тебе выбрать что заказать",
	)
	// answer.AddInlineMarkup(&kb)
	theBot.SendMessage(answer)
}

func ButtonMenuFilter(update bot.Update, theBot *bot.Bot) bool {
	if !theBot.HasFeature(DishCategoriesFeature) {
		return false
	}
	return update.CallbackQuery.Data == "menu"
}

func CommandMenuFilter(update bot.Update, theBot *bot.Bot) bool {
	if !theBot.HasFeature(DishCategoriesFeature) {
		return false
	}
	return update.Message.Text == "/menu"
}

func MenuHandler(update bot.Update, theBot *bot.Bot) {
	kb := bot.ReplyKeyboardMarkup{
		Keyboard: [][]bot.KeyboardButton{
			{
				{Text: "Супы"},
				{Text: "Салаты"},
			},
			{
				{Text: "Закуски"},
				{Text: "Охладительные напитки"},
			},
			{
				{Text: "Десерты"},
				{Text: "Горячие напитки"},
			},
		},
		ResizeKeyboard: true,
	}

	var theMessage bot.Message

	if update.Message.Chat.ID != 0 {
		theMessage = update.Message
	}
	if update.CallbackQuery.Message.Chat.ID != 0 {
		theMessage = update.CallbackQuery.Message
	}

	answer := bot.NewTextAnswer(
		theMessage.Chat.ID,
		"Выберите категорию",
	)
	answer.AddReplyMarkup(&kb)

	theBot.SendMessage(answer)
}

func CategoryFilter(update bot.Update, theBot *bot.Bot) bool {
	if !theBot.HasFeature(DishCategoriesFeature) {
		return false
	}
	msgText := update.Message.Text

	if msgText == "Супы" || msgText == "Салаты" || msgText == "Закуски" || msgText == "Охладительные напитки" || msgText == "Горячие напитки" || msgText == "Десерты" {
		return true
	}

	return false
}

func CategoryHandler(update bot.Update, theBot *bot.Bot) {
	kb := bot.ReplyKeyboardRemove{
		RemoveKeyboard: true,
	}
	menu := make(map[string][2]string)
	menu["Супы"] = [2]string{
		"Луковый суп",
		"images/soups1.jpg",
	}
	menu["Салаты"] = [2]string{
		"Цезарь",
		"images/salads1.jpg",
	}
	menu["Закуски"] = [2]string{
		"Салат из краба",
		"images/snacks1.jpg",
	}
	menu["Охладительные напитки"] = [2]string{
		"Кола",
		"images/beverages1.jpg",
	}
	menu["Горячие напитки"] = [2]string{
		"Кофе",
		"images/beverages2.jpg",
	}
	menu["Десерты"] = [2]string{
		"Капкейки",
		"images/desserts1.jpg",
	}
	answer := bot.NewTextAnswer(
		update.Message.Chat.ID,
		"Вот блюда из выбранной категории:",
	)
	answer.AddReplyRemove(&kb)
	theBot.SendMessage(answer)

	photoData := menu[update.Message.Text]
	photo, err := os.Open(photoData[1])

	if err != nil {
		fmt.Println("Error opening image", err)
	}
	defer photo.Close()

	answer2 := bot.NewPhotoAnswer(
		update.Message.Chat.ID,
		photo,
		photoData[1],
		photoData[0],
	)
	// answer2.AddPhoto(menu[update.Message.Text][1])
	theBot.SendPhoto(answer2)
}

func AddMenuHandlers() [][2]interface{} {
	handlers := [][2]interface{}{
		{CommandStartFilter, StartHandler},
		{ButtonMenuFilter, MenuHandler},
		{CommandMenuFilter, MenuHandler},
		{CategoryFilter, CategoryHandler},
	}

	return handlers
}
