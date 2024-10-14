package startinlinebuttons

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"

	handlers "github.com/liligga/hw_tg_bot/internal/handlers"
	bot "github.com/liligga/hw_tg_bot/pkg/bot"
)

var recipies = [4]string{
	"Омлет: Для приготовления омлета разбейте 3 куринных яйца и добавьте пол стакана молока, все взбейте и запеките",
	"Котлета: Для приготовления котлеты смешайте фарш с яйцом, тестом, добавьте соль и запекайте",
	"Борщ: Для приготовления борща поставьте мясо вариться на 40 минут, сделайте зажарку. Добавьте капусту и картофель в бульон. Добавьте зажарку",
	"Ризотто: В сотейнике растопите масло и обжарьте на нём измельчённый лук. Выложите рис и обжаривайте пару минут.Постепенно влейте в рис вино и горячий бульон. Затем добавьте рыбу, нарезанную небольшими кусочками, рубленый зелёный лук и шпинат и перемешайте.  Готовьте ещё немного, пока шпинат не размягчится. При необходимости добавьте в ризотто соль.",
}

func CommandStartFilter(update bot.Update, theBot *bot.Bot) bool {
	if !theBot.HasFeature(handlers.InlineButtonsMenuFeature) {
		return false
	}
	return update.Message.Text == "/start"
}

func StartHandler(update bot.Update, theBot *bot.Bot) {
	kb := bot.InlineKeyboardMarkup{
		InlineKeyboard: [][]bot.InlineKeyboardButton{
			{
				{Text: "Наш сайт", URL: "https://github.com/liligga"},
				{Text: "Расписание", CallbackData: "schedule"},
				{Text: "Контакты", CallbackData: "contacts"},
			},
			{
				{Text: "Вакансии", CallbackData: "vacancies"},
				{Text: "О нас", CallbackData: "about"},
				{Text: "Оставить отзыв", CallbackData: "restaurant_review"},
			},
		},
	}

	text := "Привет, я бот ресторана 'Ресторан'. Я помогу тебе выбрать что заказать\n" +
		"Доступные категории блюд: \n" +
		"Супы, Салаты, Закуски, Охладительные напитки, Горячие напитки, Десерты\n" // +
		// "Также продолжают работать команды:\n" +
		// "/start - начать работу с ботом\n" +
		// "/random_recipe - рандомный рецепт\n" +
		// "/myinfo - информация о тебе\n"
	answer := bot.NewTextAnswer(
		update.Message.Chat.ID,
		text,
	)
	answer.AddInlineMarkup(&kb)
	theBot.SendMessage(answer)
}

func ContactsButtonFilter(update bot.Update, theBot *bot.Bot) bool {
	if !theBot.HasFeature(handlers.InlineButtonsMenuFeature) {
		return false
	}
	data := update.CallbackQuery.Data
	return data == "contacts"
}

func ContactsHandler(update bot.Update, theBot *bot.Bot) {
	answer := bot.NewTextAnswer(
		update.CallbackQuery.Message.Chat.ID,
		"Контакты:\n"+"Телефон: +996 555 555 555\n"+"Адрес: г. Бишкек, ул. Ленина, д. 1",
	)

	theBot.SendMessage(answer)
}

func ScheduleButtonFilter(update bot.Update, theBot *bot.Bot) bool {
	if !theBot.HasFeature(handlers.InlineButtonsMenuFeature) {
		return false
	}
	data := update.CallbackQuery.Data
	return data == "schedule"
}

func ScheduleHandler(update bot.Update, theBot *bot.Bot) {
	answer := bot.NewTextAnswer(
		update.CallbackQuery.Message.Chat.ID,
		"Расписание работы ресторана:\n"+"Понедельник - Пятница: 10:00 - 22:00\n"+"Суббота: 10:00 - 19:00\n"+"Воскресенье: выходной",
	)

	theBot.SendMessage(answer)
}

func VacanciesButtonFilter(update bot.Update, theBot *bot.Bot) bool {
	if !theBot.HasFeature(handlers.InlineButtonsMenuFeature) {
		return false
	}
	data := update.CallbackQuery.Data
	return data == "vacancies"
}

func VacanciesHandler(update bot.Update, theBot *bot.Bot) {
	answer := bot.NewTextAnswer(
		update.CallbackQuery.Message.Chat.ID,
		"В данный момент нет открытых вакансий",
	)

	theBot.SendMessage(answer)
}

func AboutButtonFilter(update bot.Update, theBot *bot.Bot) bool {
	if !theBot.HasFeature(handlers.InlineButtonsMenuFeature) {
		return false
	}
	data := update.CallbackQuery.Data
	return data == "about"
}

func AboutHandler(update bot.Update, theBot *bot.Bot) {
	answer := bot.NewTextAnswer(
		update.CallbackQuery.Message.Chat.ID,
		"О нас:\n"+"ООО 'Ресторан', 2023",
	)

	theBot.SendMessage(answer)
}

func FoodCategoryFilter(update bot.Update, theBot *bot.Bot) bool {
	if !theBot.HasFeature(handlers.InlineButtonsMenuFeature) {
		return false
	}
	msgText := update.Message.Text

	if msgText == "Супы" || msgText == "Салаты" || msgText == "Закуски" || msgText == "Охладительные напитки" || msgText == "Горячие напитки" || msgText == "Десерты" {
		return true
	}

	return false
}

func FoodCategoryHandler(update bot.Update, theBot *bot.Bot) {
	kb := bot.ReplyKeyboardRemove{
		RemoveKeyboard: true,
	}
	menu := make(map[string][2]string)
	menu["Супы"] = [2]string{
		"Луковый суп",
		"internal/images/soups1.jpg",
	}
	menu["Салаты"] = [2]string{
		"Цезарь",
		"internal/images/salads1.webp",
	}
	menu["Закуски"] = [2]string{
		"Салат из краба",
		"internal/images/snacks1.jpg",
	}
	menu["Охладительные напитки"] = [2]string{
		"Кола",
		"internal/images/beverages1.webp",
	}
	menu["Горячие напитки"] = [2]string{
		"Кофе",
		"internal/images/beverages2.jpg",
	}
	menu["Десерты"] = [2]string{
		"Капкейки",
		"internal/images/desserts1.jpg",
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

func RandomRecipeFilter2(update bot.Update, theBot *bot.Bot) bool {
	if !theBot.HasFeature(handlers.InlineButtonsMenuFeature) {
		return false
	}
	return update.Message.Text == "/random_recipe"
}

func RandomRecipeHandler2(update bot.Update, theBot *bot.Bot) {
	selectedRecipe := recipies[rand.Intn(len(recipies))]
	answer := bot.NewTextAnswer(
		update.Message.Chat.ID,
		"Вот ваш рецепт:\n"+selectedRecipe,
	)
	theBot.SendMessage(answer)
}

func InfoCommandFilter2(update bot.Update, theBot *bot.Bot) bool {
	if !theBot.HasFeature(handlers.InlineButtonsMenuFeature) {
		return false
	}
	return update.Message.Text == "/myinfo"
}

func InfoCommandHandler2(update bot.Update, theBot *bot.Bot) {
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

func AddCategoriesHandlers() [][2]interface{} {
	handlers := [][2]interface{}{
		{FoodCategoryFilter, FoodCategoryHandler},
		{ContactsButtonFilter, ContactsHandler},
		{ScheduleButtonFilter, ScheduleHandler},
		{VacanciesButtonFilter, VacanciesHandler},
		{AboutButtonFilter, AboutHandler},
		// {RandomRecipeFilter2, RandomRecipeHandler2},
		// {InfoCommandFilter2, InfoCommandHandler2},
		{CommandStartFilter, StartHandler},
	}

	return handlers
}
