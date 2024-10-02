package bot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"sync"
)

type Bot struct {
	token        string
	client       *http.Client
	userContexts map[int]*UserContext // context.Context
	features     map[Feature]interface{}
	commands     map[string]string
	mu           sync.RWMutex
}

type Dispatcher struct {
	Bot      Bot
	Handlers [][2]interface{}
}

func NewDispatcher(token string, client *http.Client) Dispatcher {
	return Dispatcher{
		Bot: Bot{
			token:        token,
			client:       client,
			userContexts: make(map[int]*UserContext),
			features:     make(map[Feature]interface{}),
			commands:     make(map[string]string),
		},
		Handlers: make([][2]interface{}, 0),
	}
}

func (dp *Dispatcher) OnStartup(f func(dpp *Dispatcher)) {
	f(dp)
}

func (dp *Dispatcher) AddHandler(handler [2]interface{}) {
	dp.Handlers = append(dp.Handlers, handler)
}

func (dp *Dispatcher) AddHandlers(handlers ...[2]interface{}) {
	dp.Handlers = append(dp.Handlers, handlers...)
}

func (theBot *Bot) AddCommand(command BotCommand) {
	theBot.commands[command.Command] = command.Description
}

func (theBot *Bot) GetMyCommands() BotCommands {
	var commands []BotCommand
	for k, v := range theBot.commands {
		commands = append(commands, BotCommand{
			Command:     k,
			Description: v,
		})
	}
	return BotCommands{
		Commands: commands,
		Scope:    BotCommandScopeDefaultStruct{Type: BotCommandScopeDefault},
	}
}

type AttachedKeyboard struct {
	// ChatID 			int						`json:"chat_id"`
	// Text 			string					`json:"text"`
	ReplyMarkup    *ReplyKeyboardMarkup  `json:"reply_markup"`
	KeyBoardRemove *ReplyKeyboardRemove  `json:"reply_markup"`
	InlineMarkup   *InlineKeyboardMarkup `json:"inline_markup"`
}

type SimpleAnswer interface {
	AddKeyboard()
}

type TextAnswer struct {
	ChatID           int               `json:"chat_id"`
	Text             string            `json:"text"`
	attachedKeyboard *AttachedKeyboard `json:"-"`
}

func NewTextAnswer(chatID int, text string) *TextAnswer {
	attachedKeyboard := AttachedKeyboard{
		// ChatID: 		chatID,
		// Text: 			text,
		ReplyMarkup:    nil,
		KeyBoardRemove: nil,
		InlineMarkup:   nil,
	}

	textAnswer := TextAnswer{
		ChatID:           chatID,
		Text:             text,
		attachedKeyboard: &attachedKeyboard,
	}
	return &textAnswer
}

type PhotoAnswer struct {
	ChatID           int               `json:"chat_id"`
	Photo            *os.File          `json:"photo"`
	PhotoAddress     string            `json:"-"`
	Caption          string            `json:"caption"`
	attachedKeyboard *AttachedKeyboard `json:"-"`
}

func NewPhotoAnswer(
	chatID int,
	photo *os.File,
	photoAddress string,
	caption string,
) *PhotoAnswer {
	attachedKeyboard := AttachedKeyboard{
		ReplyMarkup:    nil,
		KeyBoardRemove: nil,
		InlineMarkup:   nil,
	}

	photoAnswer := PhotoAnswer{
		ChatID:           chatID,
		Photo:            photo,
		PhotoAddress:     photoAddress,
		Caption:          caption,
		attachedKeyboard: &attachedKeyboard,
	}
	return &photoAnswer
}

func (textAnswer *TextAnswer) AddReplyMarkup(replyMarkup *ReplyKeyboardMarkup) {
	textAnswer.attachedKeyboard.ReplyMarkup = replyMarkup
}

func (textAnswer *TextAnswer) AddReplyRemove(keyBoardRemove *ReplyKeyboardRemove) {
	textAnswer.attachedKeyboard.KeyBoardRemove = keyBoardRemove
}

func (textAnswer *TextAnswer) AddInlineMarkup(inlineMarkup *InlineKeyboardMarkup) {
	textAnswer.attachedKeyboard.InlineMarkup = inlineMarkup
}

func createRequestBody(
	chatID int,
	text string,
	replyMarkup *ReplyKeyboardMarkup,
	keyBoardRemove *ReplyKeyboardRemove,
	inlineMarkup *InlineKeyboardMarkup,
) []byte {

	body := []byte(`{
		"chat_id": ` + fmt.Sprintf("%d", chatID) + `,
		"text": "` + text + `",
	}`)

	if replyMarkup != nil {
		jsonMarshal, err := json.Marshal(replyMarkup)
		if err != nil {
			fmt.Println("Error while marshaling replyMarkup: ", err)
			return nil
		}
		body = []byte(`{
			"chat_id": ` + fmt.Sprintf("%d", chatID) + `,
			"text": "` + text + `",
			"reply_markup": ` + string(jsonMarshal) + `,
		}`)
	}

	if keyBoardRemove != nil {
		jsonMarshal, err := json.Marshal(keyBoardRemove)
		if err != nil {
			fmt.Println("Error while marshaling keyBoardRemove: ", err)
			return nil
		}
		body = []byte(`{
			"chat_id": ` + fmt.Sprintf("%d", chatID) + `,
			"text": "` + text + `",
			"reply_markup": ` + string(jsonMarshal) + `,
		}`)
	}

	if inlineMarkup != nil {
		jsonMarshal, err := json.Marshal(inlineMarkup)
		if err != nil {
			fmt.Println("Error while marshaling inlineMarkup: ", err)
			return nil
		}
		body = []byte(`{
			"chat_id": ` + fmt.Sprintf("%d", chatID) + `,
			"text": "` + text + `",
			"reply_markup": ` + string(jsonMarshal) + `,
		}`)
	}

	return body
}

func (bot *Bot) SendMessage(
	answer *TextAnswer,
) {
	url := bot.createRequestURL("sendMessage")

	body := createRequestBody(
		answer.ChatID,
		answer.Text,
		answer.attachedKeyboard.ReplyMarkup,
		answer.attachedKeyboard.KeyBoardRemove,
		answer.attachedKeyboard.InlineMarkup,
	)
	if body == nil {
		fmt.Println("Error while creating request body")
		return
	}

	rq, err := http.NewRequest(
		http.MethodPost,
		url,
		bytes.NewBuffer(body),
	)

	if err != nil {
		fmt.Println("Error constructing request: ", err)
		return
	}
	rq.Header.Add("Content-Type", "application/json")

	response, err := bot.client.Do(rq)
	if err != nil {
		fmt.Println("Error when executing request: ", err)
		return
	}

	defer response.Body.Close()
	// io.Copy(os.Stdout, response.Body)
	// fmt.Println("")
}

func (bot *Bot) SendPhoto(
	answer *PhotoAnswer,
) {
	url := bot.createRequestURL("sendPhoto")

	body := bytes.Buffer{}
	writer := multipart.NewWriter(&body)

	writer.WriteField("chat_id", fmt.Sprintf("%d", answer.ChatID))
	writer.WriteField("caption", answer.Caption)

	part, err := writer.CreateFormFile("photo", answer.PhotoAddress)
	if err != nil {
		fmt.Println("Error while creating form file: ", err)
		return
	}

	_, err = io.Copy(part, answer.Photo)
	if err != nil {
		fmt.Println("Error while copying photo: ", err)
		return
	}

	err = writer.Close()
	if err != nil {
		fmt.Println("Error while closing writer: ", err)
		return
	}

	// fmt.Println("Photo: ", string(answer.Photo)	)

	// body := []byte(`{
	// 	"chat_id": ` + fmt.Sprintf("%d", answer.ChatID) + `,
	// 	"photo": "` + string(answer.Photo) + `",
	// 	"caption": "` + answer.Caption + `",
	// }`)
	// if body == nil {
	// 	fmt.Println("Error while creating request body")
	// 	return
	// }

	// fmt.Println("Body: ", string(body))

	// fmt.Println("Body: ", &body)

	rq, err := http.NewRequest(
		http.MethodPost,
		url,
		&body,
	)
	if err != nil {
		fmt.Println("Error constructing request: ", err)
		return
	}

	// rq.Header.Set("Content-Type", "multipart/form-data")
	rq.Header.Set("Content-Type", writer.FormDataContentType())
	response, err := bot.client.Do(rq)
	if err != nil {
		fmt.Println("Error when executing request: ", err)
		return
	}

	defer response.Body.Close()
	// io.Copy(os.Stdout, response.Body)
	// fmt.Println("")
}

type PollAnswer struct {
	ChatID      int      `json:"chat_id"`
	Question    string   `json:"question"`
	Options     []string `json:"options"`
	IsAnonymous bool     `json:"is_anonymous"`
}

func NewPollAnswer(
	chatID int,
	question string,
	options []string,
) *PollAnswer {
	return &PollAnswer{
		ChatID:      chatID,
		Question:    question,
		Options:     options,
		IsAnonymous: false,
	}
}

func (bot *Bot) SendPoll(
	answer *PollAnswer,
) {
	url := bot.createRequestURL("sendPoll")

	// body := []byte(`{
	// 	"chat_id": ` + fmt.Sprintf("%d", answer.ChatID) + `,
	// 	"question": "` + answer.Question + `",
	// 	"options": ` + fmt.Sprintf("%v", answer.Options) + `,
	// }`)

	body, err := json.Marshal(answer)
	if err != nil {
		fmt.Println("Error while marshaling body: ", err)
		return
	}

	rq, err := http.NewRequest(
		http.MethodPost,
		url,
		bytes.NewBuffer(body),
	)

	if err != nil {
		fmt.Println("Error constructing request: ", err)
		return
	}

	rq.Header.Add("Content-Type", "application/json")

	response, err := bot.client.Do(rq)
	if err != nil {
		fmt.Println("Error when executing request: ", err)
		return
	}

	defer response.Body.Close()
	io.Copy(os.Stdout, response.Body)
	fmt.Println("")
}
