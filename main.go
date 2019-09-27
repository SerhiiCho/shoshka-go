package main

import (
	"fmt"

	"github.com/SerhiiCho/shoshka-go/telegram"
	"github.com/SerhiiCho/shoshka-go/utils"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	utils.HandleError(err, "Error loading .env file")
}

func main() {
	doneChan := make(chan int)
	reportsMessagesChan := make(chan string)
	errorsMessagesChan := make(chan string)
	pingMessagesChan := make(chan string)

	go telegram.GetMessagesWithNewReports(reportsMessagesChan, doneChan)
	go telegram.GetMessagesWithNewErrors(errorsMessagesChan, doneChan)
	go telegram.GetMessagesIfPingIsNotSuccessfull(pingMessagesChan, doneChan)

	for i := 3; i > 0; {
		select {
		case msg1 := <-reportsMessagesChan:
			telegram.SendMessage(msg1)
		case msg2 := <-errorsMessagesChan:
			telegram.SendMessage(msg2)
		case <-doneChan:
			i--
		}
	}
}
