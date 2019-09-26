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

// GetMessagesWithNewReports puts messages into a chanel
func GetMessagesWithNewReports(messagesChan chan<- string, doneChan chan<- int) {
	for _, report := range getReportsIfExist() {
		messagesChan <- fmt.Sprintf("Новый фотоотчет!\n\n%s\n\n%s", report.Title, report.URL)
	}

// GetMessagesWithNewErrors puts messages into a chanel
func GetMessagesWithNewErrors(messagesChan chan<- string, doneChan chan<- int) {
	for _, errorMessage := range getErrorsIfExist() {
		messagesChan <- fmt.Sprintf("Ошибка на Шобаре!\n\n%s", errorMessage)
	}
	doneChan <- 1
}

func getErrorsIfExist() []string {
	errorsContext := utils.FileGetContents(os.Getenv("BOT_ERROR_LOG_PATH"))

	newErrors := utils.ParseErrors(errorsContext)
	oldErrors := utils.GetCached("errors")
	uniqueErrors := utils.GetUniqueItem(oldErrors, newErrors)

	if len(uniqueErrors) > 0 {
		defer utils.PutIntoCache(newErrors, "errors")
	}

	return uniqueErrors
}

// getReportsIfExist makes request gets data and searches for new Photo Reports
func getReportsIfExist() []models.PhotoReport {
	html := utils.GetHTMLFromTargetURL(os.Getenv("BOT_TARGET_URL"))
	linksHTML := utils.GetLinksFromHTML(html)
	photoReports := utils.GetAllInformation(linksHTML)
	titles := utils.GetTitlesFromPhotoReports(photoReports)
	tgMessageData := utils.GenerateMapOfNewData(titles, photoReports)

	if len(titles) > 0 {
		defer utils.PutIntoCache(titles, "titles")
	}

	return tgMessageData
}
