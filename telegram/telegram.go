package telegram

import (
	"fmt"
	"github.com/SerhiiCho/shoshka_go/models"
	"os"

	"github.com/SerhiiCho/shoshka_go/utils"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// SendMessage sends message in telegram chat
func SendMessage(message string) {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	utils.HandleError(err, "Bot init error")

	msg := tgbotapi.NewMessage(getChatID(), message)

	_, _ = bot.Send(msg)
}

func SendMessagesWithNewPhotoReports(newReports []models.PhotoReport) {
	for _, report := range newReports {
		SendMessage(fmt.Sprintf("Новый фотоотчет!\n\n%s\n\n%s", report.Title, report.URL))
	}
}
