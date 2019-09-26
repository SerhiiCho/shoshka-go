package telegram

import (
	"os"
	"strconv"

	"github.com/SerhiiCho/shoshka-go/utils"
)

func getChatID() int64 {
	chatID, err := strconv.ParseInt(os.Getenv("BOT_CHAT_ID"), 10, 64)
	utils.HandleError(err, "Cannot convert chat id to integer")
	return chatID
}

// GenerateMessagesWithNewPhotoReports generates telegram message for new photo report
func GenerateMessagesWithNewPhotoReports() string {
	msg := ""

	for _, report := range GetTelegramMessageIfExists() {
		msg += fmt.Sprintf("Новый фотоотчет!\n\n%s\n\n%s", report.Title, report.URL)
	}

	return msg
}

