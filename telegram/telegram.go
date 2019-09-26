package telegram

import (
	"fmt"
	"github.com/SerhiiCho/shoshka-go/config"
	"github.com/SerhiiCho/shoshka-go/utils"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// SendMessage sends message in telegram chat
func SendMessage(message string) {
	if len(message) == 0 {
		fmt.Println("There are not new Photo Reports")
		return
	}

	bot, err := tgbotapi.NewBotAPI(config.Config["botToken"])
	utils.HandleError(err, "Bot init error")

	msg := tgbotapi.NewMessage(getChatID(), message)

	_, _ = bot.Send(msg)
}
