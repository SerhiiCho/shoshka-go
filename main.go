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
	messagesChan := make(chan string)

	go telegram.GenerateMessagesWithNewPhotoReports(messagesChan)

	for msg := range messagesChan {
		telegram.SendMessage(msg)
	}
}
