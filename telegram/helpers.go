package telegram

import (
	"fmt"
	"os"
	"strconv"

	"github.com/SerhiiCho/shoshka-go/models"
	"github.com/SerhiiCho/shoshka-go/utils"
)

func getChatID() int64 {
	chatID, err := strconv.ParseInt(os.Getenv("BOT_CHAT_ID"), 10, 64)
	utils.HandleError(err, "Cannot convert chat id to integer")
	return chatID
}

// GetMessagesWithNewReports generates telegram message for new photo report
func GetMessagesWithNewReports(messagesChan chan<- string) {
	for _, report := range GetTelegramMessageIfExists() {
		messagesChan <- fmt.Sprintf("Новый фотоотчет!\n\n%s\n\n%s", report.Title, report.URL)
	}
	close(messagesChan)
}

// GetTelegramMessageIfExists makes request gets data and searches for new Photo Reports
func GetTelegramMessageIfExists() []models.PhotoReport {
	html := utils.GetHTMLFromTargetURL(os.Getenv("BOT_TARGET_URL"))
	linksHTML := utils.GetLinksFromHTML(html)
	photoReports := utils.GetAllInformation(linksHTML)
	titles := utils.GetTitlesFromPhotoReports(photoReports)
	tgMessageData := utils.GenerateMapOfNewData(titles, photoReports)

	if len(titles) > 0 {
		defer utils.PutTitlesIntoCache(titles)
	}

	return tgMessageData
}
