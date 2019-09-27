package main

import (
	"os"

	"github.com/SerhiiCho/shoshka-go/telegram"
	"github.com/SerhiiCho/shoshka-go/utils"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	utils.HandleError(err, "Error loading .env file")
}

func main() {
	var messages []string

	if os.Args[1] == "errors" {
		messages = telegram.GetMessagesWithNewReports()
	} else if os.Args[1] == "titles" {
		messages = telegram.GetMessagesWithNewErrors()
	} else if os.Args[1] == "ping" {
		messages = telegram.GetMessageIfPingIsNotSuccessful()
	}

	for _, msg := range messages {
		telegram.SendMessage(msg)
	}
}
