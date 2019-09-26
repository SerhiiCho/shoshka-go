package telegram

import (
	"os"

	"github.com/SerhiiCho/shoshka-go/utils"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// SendMessage sends message in telegram chat
func SendMessage(message string) {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	utils.HandleError(err, "Bot init error")

	msg := tgbotapi.NewMessage(getChatID(), message)

	_, _ = bot.Send(msg)
}
