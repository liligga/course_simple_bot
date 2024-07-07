package bot

import "fmt"

type Message struct {
	MessageID			int		`json:"message_id"`
	From struct {
		ID				int		`json:"id"`
		IsBot 			bool	`json:"is_bot"`
		FirstName 		string	`json:"first_name"`
		Username 		string	`json:"username"`
		LanguageCode 	string	`json:"language_code"`
		IsPremium 		bool	`json:"is_premium"`
	}							`json:"from"`
	Chat struct  { 
		ID				int		`json:"id"`
		FirstName 		string	`json:"first_name"`
		Username 		string	`json:"username"`
		Type 			string	`json:"type"`
	}							`json:"chat"`
	Text				string	`json:"text"`
}

func (m *Message) Answer(text string) {
	fmt.Printf("%d: %s", m.MessageID, m.Text)
}

type Update struct {
	UpdateID 	int 	`json:"update_id"`
	Message 	Message	`json:"message"`
}

type APIResponse struct {
	Ok 			bool		`json:"ok"`
	Results 	[]Update	`json:"result"`
}