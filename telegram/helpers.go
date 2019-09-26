package telegram

import (
	"fmt"
	"github.com/SerhiiCho/shoshka-go/config"
	"strconv"

	"github.com/SerhiiCho/shoshka-go/models"
	"github.com/SerhiiCho/shoshka-go/utils"
)

func getChatID() int64 {
	chatID, err := strconv.ParseInt(config.Config["botChatID"], 10, 64)
	utils.HandleError(err, "Cannot convert chat id to integer")
	return chatID
}

// GenerateMessagesWithNewPhotoReports generates telegram message for new photo report
func GenerateMessagesWithNewPhotoReports(messagesChan chan<- string) {
	for _, report := range GetTelegramMessageIfExists() {
		messagesChan <- fmt.Sprintf("Новый фотоотчет!\n\n%s\n\n%s", report.Title, report.URL)
	}
	close(messagesChan)
}

// GetTelegramMessageIfExists makes request gets data and searches for new Photo Reports
func GetTelegramMessageIfExists() []models.PhotoReport {
	html := utils.GetHTMLFromTargetURL(config.Config["botTargetURL"])
	linksHTML := utils.GetLinksFromHTML(html)
	photoReports := utils.GetAllInformation(linksHTML)
	titles := utils.GetTitlesFromPhotoReports(photoReports)
	tgMessageData := utils.GenerateMapOfNewData(titles, photoReports)

	if len(titles) > 0 {
		defer utils.PutTitlesIntoCache(titles)
	}

	return tgMessageData
}
