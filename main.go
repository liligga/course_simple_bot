package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	// "github.com/joho/godotenv"

	// "github.com/liligga/hw_tg_bot/bot"
	"github.com/liligga/hw_tg_bot/handlers"
)


func main() {
	var wg sync.WaitGroup
	// ctx := context.Background()
	client := &http.Client{
		Timeout: time.Duration(10) * time.Second,
	}

	myDispatcher := handlers.NewApp(client)


	fmt.Println("Hello, World!")

	wg.Add(1)
	go myDispatcher.GetMeHandler(&wg, client)

	wg.Add(1)
	go myDispatcher.LongPollingTgAPI(&wg, client)

	wg.Wait()
}
