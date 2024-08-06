package bot

type GeneralKeyboard interface {
}

type KeyboardButton struct {
	Text string `json:"text"`
}

type ReplyKeyboardMarkup struct {
	Keyboard              [][]KeyboardButton `json:"keyboard"`
	InputFieldPlaceholder string             `json:"input_field_placeholder"`
	ResizeKeyboard        bool               `json:"resize_keyboard" default:"false"`
	// OneTimeKeyboard			bool				`default:"false"`
	// Selective 				bool				`default:"false"`
}

type ReplyKeyboardRemove struct {
	RemoveKeyboard bool `json:"remove_keyboard"`
}

type WebAppInfo struct {
	URL string `json:"url"`
}

type InlineKeyboardButton struct {
	Text         string `json:"text"`
	URL          string `json:"url"`
	CallbackData string `json:"callback_data"`
	// WebAppInfo WebAppInfo `json:"web_app"`
}

type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}
