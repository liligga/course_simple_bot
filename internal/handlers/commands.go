package handlers

import (
	bot "github.com/liligga/hw_tg_bot/pkg/bot"
)

func SetMyCommandsAdmin(theBot *bot.Bot) {
	adminCommands := bot.BotCommands{
		Commands: []bot.BotCommand{
			{
				Command:     "/admin",
				Description: "Показать меню администрирования",
			},
		},
		Scope: bot.BotCommandScopeChatStruct{
			Type:   bot.BotCommandScopeChat,
			ChatID: 243154734,
		},
	}
	theBot.SetMyCommands(adminCommands)
}

func SetMyCommandsProducts(theBot *bot.Bot) {
	productsCommands := bot.BotCommands{
		Commands: []bot.BotCommand{
			{
				Command:     "/start",
				Description: "Начало работы",
			},
			{
				Command:     "/random_recipe",
				Description: "Показать рецепт",
			},
			{
				Command:     "/myinfo",
				Description: "Мои данные",
			},
			{
				Command:     "/dishes",
				Description: "Показать список блюд",
			},
		},
		Scope: bot.BotCommandScopeAllPrivateChatsStruct{
			Type: bot.BotCommandScopeAllPrivateChats,
		},
	}
	theBot.SetMyCommands(productsCommands)
}
