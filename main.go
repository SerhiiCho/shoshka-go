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

	go telegram.GetMessagesWithNewReports(reportsMessagesChan, doneChan)
	go telegram.GetMessagesWithNewErrors(errorsMessagesChan, doneChan)

	for i := 2; i > 0; {
		select {
		case msg1 := <-reportsMessagesChan:
			fmt.Println(msg1)
			// telegram.SendMessage(msg1)
		case msg2 := <-errorsMessagesChan:
			fmt.Println(msg2)
			// telegram.SendMessage(msg2)
		case <-doneChan:
			i--
		}
	}
}
