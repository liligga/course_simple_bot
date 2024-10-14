package restaurantreview

import (
	// "slices"

	"slices"

	handlers "github.com/liligga/hw_tg_bot/internal/handlers"
	bot "github.com/liligga/hw_tg_bot/pkg/bot"
)

type RestaurantReview string

const (
	UsersName     RestaurantReview = "users_name"
	UsersPhone    RestaurantReview = "users_phone"
	VisitDate     RestaurantReview = "visit_date"
	VisitMonth    RestaurantReview = "visit_month"
	ServingRating RestaurantReview = "serving_rating"
	UserComment   RestaurantReview = "user_comment"
)

func CommandStartFilter(update bot.Update, theBot *bot.Bot) bool {
	if !theBot.HasFeature(handlers.ReviewFeature) {
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
	if !theBot.HasFeature(handlers.ReviewFeature) {
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
	if !theBot.HasFeature(handlers.ReviewFeature) {
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
	if !theBot.HasFeature(handlers.ReviewFeature) {
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
	if !theBot.HasFeature(handlers.ReviewFeature) {
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

func RestaurantReviewFilter(update bot.Update, theBot *bot.Bot) bool {
	if !theBot.HasFeature(handlers.ReviewFeature) {
		return false
	}
	return update.CallbackQuery.Data == "restaurant_review"
}

func StartRestaurantReviewHandler(update bot.Update, theBot *bot.Bot) {
	answer := bot.NewTextAnswer(
		update.CallbackQuery.Message.Chat.ID,
		"Представьтесь пожалуйста:",
	)
	theBot.SetContextState(
		update.CallbackQuery.From.ID,
		update.CallbackQuery.Message.Chat.ID,
		string(UsersName),
	)

	theBot.SendMessage(answer)
}

func StopReviewFilter(update bot.Update, theBot *bot.Bot) bool {
	if !theBot.HasFeature(handlers.ReviewFeature) {
		return false
	}
	return slices.Contains(
		[]string{"/stop", "/cancel", "стоп"},
		update.Message.Text,
	)
}

func StopReviewHandler(update bot.Update, theBot *bot.Bot) {
	answer := bot.NewTextAnswer(
		update.Message.Chat.ID,
		"Отзыв не будет сохранен",
	)

	theBot.DeleteUserContext(
		update.Message.From.ID,
		update.Message.Chat.ID,
	)

	theBot.SendMessage(answer)
}

func ProcessUsersNameFilter(update bot.Update, theBot *bot.Bot) bool {
	userContext, err := theBot.GetUserContext(
		update.Message.From.ID,
		update.Message.Chat.ID,
	)

	return err == nil && userContext.CurrentState == string(UsersName)
}

func ProcessUsersNameHandler(update bot.Update, theBot *bot.Bot) {
	answer := bot.NewTextAnswer(
		update.Message.Chat.ID,
		"Оставьте Ваш номер телефона:",
	)
	theBot.SetContextState(
		update.Message.From.ID,
		update.Message.Chat.ID,
		string(UsersPhone),
	)

	theBot.SendMessage(answer)
}

func ProcessUsersPhoneFilter(update bot.Update, theBot *bot.Bot) bool {
	userContext, err := theBot.GetUserContext(
		update.Message.From.ID,
		update.Message.Chat.ID,
	)

	return err == nil && userContext.CurrentState == string(UsersPhone)
}

func ProcessUsersPhoneHandler(update bot.Update, theBot *bot.Bot) {
	answer := bot.NewTextAnswer(
		update.Message.Chat.ID,
		"Давайте выясним, когда Вы посещали наш ресторан?\nДля начала введите число(от 1 до 31)?",
	)
	theBot.SetContextState(
		update.Message.From.ID,
		update.Message.Chat.ID,
		string(VisitDate),
	)

	theBot.SendMessage(answer)
}

func ProcessVisitDateFilter(update bot.Update, theBot *bot.Bot) bool {
	userId := update.Message.From.ID
	chatId := update.Message.Chat.ID

	userContext, err := theBot.GetUserContext(userId, chatId)

	return err == nil && userContext.CurrentState == string(VisitDate)
}

func ProcessVisitDateHandler(update bot.Update, theBot *bot.Bot) {
	// correctDate := slices.Contains(
	// 	[]string{
	// 		"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15",
	// 		"16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31",
	// 	},
	// 	update.Message.Text,
	// )

	// if !correctDate {
	// 	answer := bot.NewTextAnswer(
	// 		update.Message.Chat.ID,
	// 		"Давайте выясним, когда Вы посещали наш ресторан?\nДля начала введите число(от 1 до 31)?",
	// 	)
	// 	theBot.SendMessage(answer)
	// 	return
	// }

	answer := bot.NewTextAnswer(
		update.Message.Chat.ID,
		"В каком месяце это было(от 1 до 12 или название месяца)?",
	)

	reply_keyboard := bot.ReplyKeyboardMarkup{
		ResizeKeyboard: true,
		Keyboard: [][]bot.KeyboardButton{
			{
				{Text: "Январь"},
				{Text: "Февраль"},
			},
			{
				{Text: "Март"},
				{Text: "Апрель"},
			},
			{
				{Text: "Май"},
				{Text: "Июнь"},
			},
			{
				{Text: "Июль"},
				{Text: "Август"},
			},
			{
				{Text: "Сентябрь"},
				{Text: "Октябрь"},
			},
			{
				{Text: "Ноябрь"},
				{Text: "Декабрь"},
			},
		},
		InputFieldPlaceholder: "Месяц",
	}
	answer.AddReplyMarkup(&reply_keyboard)
	theBot.SetContextState(
		update.Message.From.ID,
		update.Message.Chat.ID,
		string(VisitMonth),
	)

	theBot.SendMessage(answer)
}

func ProcessVisitMonthFilter(update bot.Update, theBot *bot.Bot) bool {
	userId := update.Message.From.ID
	chatId := update.Message.Chat.ID

	userContext, err := theBot.GetUserContext(userId, chatId)

	return err == nil && userContext.CurrentState == string(VisitMonth)
}

func ProcessVisitMonthHandler(update bot.Update, theBot *bot.Bot) {
	// correctMonth := slices.Contains(
	// 	[]string{"Январь", "Февраль", "Март", "Апрель", "Май", "Июнь", "Июль", "Август", "Сентябрь", "Октябрь", "Ноябрь", "Декабрь"},
	// 	update.Message.Text,
	// ) || slices.Contains(
	// 	[]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"},
	// 	update.Message.Text,
	// )

	// if user entered correct month
	// if !correctMonth {
	// 	answer := bot.NewTextAnswer(
	// 		update.Message.Chat.ID,
	// 		"Пожалуйста, введите число от 1 до 12 или название месяца",
	// 	)

	// 	theBot.SendMessage(answer)
	// 	return
	// }

	answer := bot.NewTextAnswer(
		update.Message.Chat.ID,
		"Поставьте нам оценку от 1 до 5, где 1 - это очень плохо, а 5 - это отлично",
	)
	theBot.SetContextState(
		update.Message.From.ID,
		update.Message.Chat.ID,
		string(ServingRating),
	)

	kb := bot.ReplyKeyboardRemove{RemoveKeyboard: true}
	answer.AddReplyRemove(&kb)

	theBot.SendMessage(answer)
}

func ProcessRatingFilter(update bot.Update, theBot *bot.Bot) bool {
	userId := update.Message.From.ID
	chatId := update.Message.Chat.ID

	userContext, err := theBot.GetUserContext(userId, chatId)

	return err == nil && userContext.CurrentState == string(ServingRating)
}

func ProcessRatingHandler(update bot.Update, theBot *bot.Bot) {
	// correctRating := slices.Contains(
	// 	[]string{"1", "2", "3", "4", "5"},
	// 	update.Message.Text,
	// )

	// if !correctRating {
	// 	answer := bot.NewTextAnswer(
	// 		update.Message.Chat.ID,
	// 		"Пожалуйста, введите число от 1 до 5",
	// 	)

	// 	theBot.SendMessage(answer)
	// }

	answer := bot.NewTextAnswer(
		update.Message.Chat.ID,
		"Оставьте дополнительный комментарий пожалуйста. В нем вы можете оставить свои пожелания и предложения, жалобы или похвалу",
	)
	theBot.SetContextState(
		update.Message.From.ID,
		update.Message.Chat.ID,
		string(UserComment),
	)

	theBot.SendMessage(answer)
}

func ProcessCommentFilter(update bot.Update, theBot *bot.Bot) bool {
	userId := update.Message.From.ID
	chatId := update.Message.Chat.ID

	userContext, err := theBot.GetUserContext(userId, chatId)

	return err == nil && userContext.CurrentState == string(UserComment)
}

func ProcessCommentHandler(update bot.Update, theBot *bot.Bot) {
	answer := bot.NewTextAnswer(
		update.Message.Chat.ID,
		"Спасибо за ваш отзыв!",
	)
	theBot.DeleteUserContext(
		update.Message.From.ID,
		update.Message.Chat.ID,
	)

	theBot.SendMessage(answer)
}

func AddReviewFSMHandlers() [][2]interface{} {
	return [][2]interface{}{
		{CommandStartFilter, StartHandler},
		{AboutButtonFilter, AboutHandler},
		{ContactsButtonFilter, ContactsHandler},
		{ScheduleButtonFilter, ScheduleHandler},
		{VacanciesButtonFilter, VacanciesHandler},
		{RestaurantReviewFilter, StartRestaurantReviewHandler},
		{ProcessUsersNameFilter, ProcessUsersNameHandler},
		{ProcessUsersPhoneFilter, ProcessUsersPhoneHandler},
		{ProcessVisitDateFilter, ProcessVisitDateHandler},
		{ProcessVisitMonthFilter, ProcessVisitMonthHandler},
		{ProcessRatingFilter, ProcessRatingHandler},
		{ProcessCommentFilter, ProcessCommentHandler},
	}
}
