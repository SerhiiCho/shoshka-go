package main

import (
	"github.com/SerhiiCho/shoshka-go/telegram"
)

func main() {
	messagesChan := make(chan string)

	go telegram.GenerateMessagesWithNewPhotoReports(messagesChan)

	for msg := range messagesChan {
		telegram.SendMessage(msg)
	}
}
