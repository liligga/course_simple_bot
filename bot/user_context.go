package bot

import "time"

type UserContext struct {
	ChatID       int
	UserID       int
	LastAccess   time.Time
	Data         map[string]interface{}
	CurrentState string
}

func (bot *Bot) GetUserContext(userID int, chatID int) *UserContext {
	bot.mu.RLock()
	defer bot.mu.RUnlock()
	if ctx, ok := bot.userContexts[userID]; ok {
		if ctx.ChatID == chatID {
			ctx.LastAccess = time.Now()
			return ctx
		}
	}

	ctx := &UserContext{
		ChatID:     chatID,
		UserID:     int(userID),
		Data:       make(map[string]interface{}),
		LastAccess: time.Now(),
	}
	bot.userContexts[userID] = ctx
	return ctx
}

func (bot *Bot) DeleteUserContext(userID int, chatID int) {
	bot.mu.Lock()
	defer bot.mu.Unlock()
	if ctx, ok := bot.userContexts[userID]; ok {
		if ctx.ChatID == chatID {
			delete(bot.userContexts, userID)
		}
	}
}
