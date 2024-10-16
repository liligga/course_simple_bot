package bot

import (
	"bytes"
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

func (dp *Dispatcher) DeleteWebhook(wg *sync.WaitGroup, client *http.Client) {
	defer wg.Done()
	url := dp.Bot.createRequestURL("deleteWebhook")
	fmt.Println(url)

	rq, err := http.NewRequest(
		http.MethodPost,
		url,
		nil,
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err := client.Do(rq)
	if err != nil {
		fmt.Println("Error while makeing DeleteWebhook request: ", err)
		return
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
	fmt.Println("")

}

func (dp *Dispatcher) SetMyCommands(wg *sync.WaitGroup, client *http.Client) {
	defer wg.Done()
	commands := dp.Bot.GetMyCommands()
	data, err := json.Marshal(commands)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))

	url := dp.Bot.createRequestURL("setMyCommands")

	rq, err := http.NewRequest(
		http.MethodPost,
		url,
		bytes.NewBuffer(data),
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err := client.Do(rq)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
}

func (dp *Dispatcher) LongPollingTgAPI(
	wg *sync.WaitGroup,
	client *http.Client,
	sleepRange time.Duration,
) {

	defer wg.Done()
	updateOffset := 0
	hadUpdates := false

	fmt.Println("LongPolling started")

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
					// check filter
					result := v[0].(func(Update, *Bot) bool)(update, &dp.Bot)
					if result {
						// if filter matches call associated handler
						v[1].(func(Update, *Bot))(update, &dp.Bot)
						break
					}
				}
			}
		} else {
			// No updates - no need to offset
			hadUpdates = false
		}

		time.Sleep(sleepRange)
	}

}
