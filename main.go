package main

import (
	"fmt"
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

	if len(os.Args) <= 1 {
		fmt.Print(`
			|-----------------------------------------------|
			| You should specify 1 of 3 possible parameters |
			|-----------------------------------------------|
		`)
		os.Exit(1)
	}

	command := os.Args[1]

	if !utils.Contains([]string{"errors", "titles", "ping"}, command) {
		fmt.Printf(`
			|-------------------------
			| Unknown parameter: %s
			|-------------------------\n\n
		`, command)
		os.Exit(1)
	}

	if command == "errors" {
		messages = telegram.GetMessagesWithNewErrors()
	} else if command == "titles" {
		messages = telegram.GetMessagesWithNewReports()
	} else if command == "ping" {
		messages = telegram.GetMessageIfPingIsNotSuccessful()
	}

	for _, msg := range messages {
		telegram.SendMessage(msg)
	}
}
