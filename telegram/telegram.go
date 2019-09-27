package telegram

import (
	"os"

	"github.com/SerhiiCho/shoshka-go/utils"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// SendMessage sends message in telegram chat
func SendMessage(message string) {
	if len(message) == 0 {
		return
	}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	utils.HandleError(err, "Bot init error")

	msg := tgbotapi.NewMessage(getChatID(), message)
	msg.DisableWebPagePreview = true

	_, err = bot.Send(msg)
	utils.HandleError(err, "Error sending telegram message")
}
