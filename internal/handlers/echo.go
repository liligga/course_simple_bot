package handlers

import (
	bot "github.com/liligga/hw_tg_bot/pkg/bot"
)

func EmptyFilter(update bot.Update, theBot *bot.Bot) bool {
	return true
}

func EchoHandler(update bot.Update, theBot *bot.Bot) {
	answer := bot.NewTextAnswer(
		update.Message.Chat.ID,
		update.Message.Text,
	)

	theBot.SendMessage(answer)
}

// func StartCommandFilter(update bot.Update, theBot *bot.Bot) bool {
// 	if !theBot.HasFeature(SimplePictureAndStartFeature) {
// 		return false
// 	}
// 	return update.Message.Text == "/start"
// }

// func StartCommandHandler(update bot.Update, theBot *bot.Bot) {
// 	answer := bot.NewTextAnswer(
// 		update.Message.Chat.ID,
// 		"Привет, "+update.Message.From.FirstName+"!",
// 	)
// 	theBot.SendMessage(answer)
// }

// func PictureCommandFilter(update bot.Update, theBot *bot.Bot) bool {
// 	if !theBot.HasFeature(SimplePictureAndStartFeature) {
// 		return false
// 	}
// 	return update.Message.Text == "/picture"
// }

// func PictureCommandHandler(update bot.Update, theBot *bot.Bot) {
// 	caption := "Кошка"
// 	image := "images/cat.jpeg"
// 	photo, err := os.Open(image)

// 	if err != nil {
// 		fmt.Println("Error opening image", err)
// 	}
// 	defer photo.Close()

// 	answer2 := bot.NewPhotoAnswer(
// 		update.Message.Chat.ID,
// 		photo,
// 		image,
// 		caption,
// 	)
// 	// answer2.AddPhoto(menu[update.Message.Text][1])
// 	theBot.SendPhoto(answer2)
// }

// func AddSimpleStartHandlers() [][2]interface{} {
// 	handlers := [][2]interface{}{
// 		{StartCommandFilter, StartCommandHandler},
// 		{PictureCommandFilter, PictureCommandHandler},
// 	}

// 	return handlers
// }
