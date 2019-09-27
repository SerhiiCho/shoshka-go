package main

import (
	"github.com/SerhiiCho/shoshka-go/telegram"
	"github.com/SerhiiCho/shoshka-go/utils"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	utils.HandleError(err, "Error loading .env file")
}

func main() {
	var allMessages []string

	allMessages = append(allMessages, telegram.GetMessagesWithNewReports()...)
	allMessages = append(allMessages, telegram.GetMessagesWithNewErrors()...)
	allMessages = append(allMessages, telegram.GetMessageIfPingIsNotSuccessful()...)

	for _, msg := range allMessages {
		telegram.SendMessage(msg)
	}
}
