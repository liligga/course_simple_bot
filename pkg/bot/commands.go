package bot

type BotCommand struct {
	Command     string `json:"command"`
	Description string `json:"description"`
}

const (
	BotCommandScopeDefault               = "default"
	BotCommandScopeAllPrivateChats       = "all_private_chats"
	BotCommandScopeAllGroupChats         = "all_group_chats"
	BotCommandScopeAllChatAdministration = "all_chat_administrators"
	BotCommandScopeChat                  = "chat"
	BotCommandScopeChatAdministrators    = "chat_administrators"
	BotCommandScopeChatMember            = "chat_member"
)

type BotCommandScopeDefaultStruct struct {
	Type string `json:"type"`
}

type BotCommandScopeAllPrivateChatsStruct struct {
	Type string `json:"type"`
}

type BotCommandScopeAllGroupChatsStruct struct {
	Type string `json:"type"`
}

type BotCommandScopeAllChatAdministratorsStruct struct {
	Type string `json:"type"`
}

type BotCommandScopeChatStruct struct {
	Type   string `json:"type"`
	ChatID int    `json:"chat_id"`
}

type BotCommandScopeChatAdministratorsStruct struct {
	Type   string `json:"type"`
	ChatID int    `json:"chat_id"`
}

type BotCommandScopeChatMemberStruct struct {
	Type   string `json:"type"`
	ChatID int    `json:"chat_id"`
	UserId int    `json:"user_id"`
}

type BotCommands struct {
	Commands []BotCommand `json:"commands"`
	Scope    interface{}  `json:"scope"`
}
