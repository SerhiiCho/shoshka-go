package telegram

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/SerhiiCho/shoshka-go/models"
	"github.com/SerhiiCho/shoshka-go/utils"
)

func getChatID() int64 {
	chatID, err := strconv.ParseInt(os.Getenv("BOT_CHAT_ID"), 10, 64)
	utils.HandleError(err, "Cannot convert chat id to integer")
	return chatID
}

// todayIsReportCheckDay returns true if current day is one of allowed for checking reports
func todayIsReportCheckDay(today string) bool {
	allowedDays := strings.Split(os.Getenv("BOT_DAYS_FOR_REPORT_CHECK"), ",")
	return utils.Contains(allowedDays, today)
}

// GetMessagesWithNewReports puts messages into a chanel
func GetMessagesWithNewReports(messagesChan chan<- string, doneChan chan<- int) {
	today := time.Now().Weekday().String()

	if todayIsReportCheckDay(today) {
		for _, report := range getReportsIfExist() {
			messagesChan <- fmt.Sprintf("Новый фотоотчет!\n\n%s\n\n%s", report.Title, report.URL)
		}
	}
	doneChan <- 1
}

// GetMessagesWithNewErrors puts messages into a chanel
func GetMessagesWithNewErrors(messagesChan chan<- string, doneChan chan<- int) {
	for _, errorMessage := range getErrorsIfExist() {
		messagesChan <- fmt.Sprintf("Ошибка на Шобаре!\nhttps://shobar.com.ua/\n\n%s", errorMessage)
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
