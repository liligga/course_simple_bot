package bot

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

func (bot *Bot) createRequestURL(method string) string {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/%s", bot.token, method)
	return url
}

func (dp *Dispatcher) GetMeHandler(wg *sync.WaitGroup, client *http.Client) {
	defer wg.Done()
	url := dp.Bot.createRequestURL("getMe")
	fmt.Println(url)

	rq, err := http.NewRequest(
		http.MethodGet,
		url,
		nil,
	)

	if err != nil {
		fmt.Println(err) 
		return
	}
	
	resp, err := client.Do(rq) 
	if err != nil { 
		fmt.Println("Error while makeing GetMe request: ", err)
		return
	} 
	defer resp.Body.Close() 

	io.Copy(os.Stdout, resp.Body)
	fmt.Println("")
}


func (dp *Dispatcher) LongPollingTgAPI(
	wg *sync.WaitGroup, 
	client *http.Client,
) {

	defer wg.Done()
	updateOffset := 0
	hadUpdates := false

	// loop to request updates
	for {

		url := dp.Bot.createRequestURL("getUpdates")
		if hadUpdates {
			updateOffset += 1
			url = fmt.Sprintf("%s?offset=%d", url, updateOffset)
			fmt.Println("Offset increased: ", updateOffset)
		}

		rq, err := http.NewRequest(
			http.MethodGet,
			url,
			nil,
		)

		if err != nil {
			fmt.Println(err)
			return
		}
	
		resp, err := client.Do(rq) 
		if err != nil { 
			fmt.Println("Error when executing request: ", err) 
			return
		}

		defer resp.Body.Close()

		// io.Copy(os.Stdout, resp.Body)
		// fmt.Println("")

		var apiResponse APIResponse //message.APIResponse
		if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
			fmt.Println("Error: Error reading json response body: ", err)
			return
		}

		if len(apiResponse.Results) > 0 {
			// to prevent same updates from being processed
			hadUpdates = true
			for _, update := range apiResponse.Results {
				updateOffset = update.UpdateID
				for _, v := range dp.Handlers {
					result := v[0].(func(Update) bool)(update)
					if result {
						v[1].(func(Update, Bot))(update, dp.Bot)
						break
					}
				}
			}
		} else {
			// No updates - no need to offset
			hadUpdates = false
		}
		
		time.Sleep(500 * time.Millisecond)
	}

}