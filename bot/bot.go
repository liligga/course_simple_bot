package bot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)


type Bot struct {
	token 	string
	client 	*http.Client
}


type Dispatcher struct {
	Bot 		Bot
	Updates 	[]Update
	Handlers 	[][2]interface{}
}


func NewDispatcher(token string, client *http.Client) Dispatcher {
	return Dispatcher{
		Bot: Bot{
			token: token,
			client: client,
		},
		Handlers: make([][2]interface{}, 0),
		// Updates: make([]Update, 0),
	}
}

type KeyboardButton struct {
	Text string		`json:"text"`
}

type ReplyKeyboardMarkup struct {
	Keyboard 				[][]KeyboardButton	`json:"keyboard"`
	InputFieldPlaceholder 	string				`json:"input_field_placeholder"`
	ResizeKeyboard 			bool				`json:"resize_keyboard" default:"false"`
	// OneTimeKeyboard			bool				`default:"false"`
	// Selective 				bool				`default:"false"`
}

type ReplyKeyboardRemove struct {
	RemoveKeyboard bool	`json:"remove_keyboard"`
}

func createRequestBody(
	chatID 			int,
	text 			string,
	replyMarkup 	*ReplyKeyboardMarkup,
	keyBoardRemove 	*ReplyKeyboardRemove,
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

	return body
}

func (bot *Bot) SendMessage(
	chatID 			int,
	text 			string,
	replyMarkup 	*ReplyKeyboardMarkup,
	keyBoardRemove 	*ReplyKeyboardRemove,
	) {
	url := bot.createRequestURL("sendMessage")
	// var data string

	// data = fmt.Sprintf("chat_id=%d&text=%s", chatID, text)
	// if replyMarkup != nil {
	// 	replyMarkupJSON, err := json.Marshal(replyMarkup)
	// 	if err != nil {
	// 		fmt.Println("Error while marshaling replyMarkup: ", err)
	// 		return
	// 	}
	// 	data = fmt.Sprintf("chat_id=%d&text=%s&reply_markup=%s", chatID, text, string(replyMarkupJSON))
	// 	fmt.Println("URL: ", data)
	// }
	
	// finalURL := fmt.Sprintf("%s?%s", url, data)

	body := createRequestBody(chatID, text, replyMarkup, keyBoardRemove)
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
	io.Copy(os.Stdout, response.Body)
	fmt.Println("")
}