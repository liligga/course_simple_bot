package main

import (
	"net/http"
	"sync"
	"time"

	internal "github.com/liligga/hw_tg_bot/internal"
)

func main() {
	var wg sync.WaitGroup
	// ctx := context.Background()
	client := &http.Client{
		Timeout: time.Duration(10) * time.Second,
	}

	myDispatcher := internal.NewApp(client)

	wg.Add(1)
	go myDispatcher.GetMeHandler(&wg, client)

	wg.Add(1)
	go myDispatcher.SetMyCommands(&wg, client)

	wg.Add(1)
	go myDispatcher.DeleteWebhook(&wg, client)

	sleepRange := time.Duration(1500) * time.Millisecond
	wg.Add(1)
	go myDispatcher.LongPollingTgAPI(&wg, client, sleepRange)

	wg.Wait()
}
