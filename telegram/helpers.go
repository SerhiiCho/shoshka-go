package telegram

import (
	"fmt"
	"os"
	"os/exec"
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
	allowedDays := strings.Split(os.Getenv("DAYS_FOR_REPORT_CHECK"), ",")
	return utils.Contains(allowedDays, today)
}

// GetMessagesWithNewReports puts messages into a chanel
func GetMessagesWithNewReports() []string {
	var messages []string
	today := time.Now().Weekday().String()

	if !todayIsReportCheckDay(today) {
		fmt.Printf("Today is %s, but messages will be send only in %s days of the week\n", today, os.Getenv("DAYS_FOR_REPORT_CHECK"))
		return messages
	}

	for _, report := range getReportsIfExist() {
		messages = append(messages, fmt.Sprintf("New Photo Report!\n%s\n%s", report.Title, report.URL))
	}

	return messages
}

// GetMessageIfPingIsNotSuccessful returns error message if ping is not successful
func GetMessageIfPingIsNotSuccessful() []string {
	var messages []string
	host := os.Getenv("SITE_ADDRESS")
	out, err := exec.Command("ping", host, "-c2").Output()

	canPing := !strings.Contains(string(out), "Destination Host Unreachable")

	if canPing && err == nil {
		fmt.Printf("Host %s is reachable\n", host)
		return messages
	}

	fmt.Printf("Host %s is reachable\n", host)
	return append(messages, fmt.Sprintf("Host %s is not reachable", host))
}

// GetMessagesWithNewErrors puts messages into a chanel
func GetMessagesWithNewErrors() []string {
	var messages []string

	for _, errorMessage := range getErrorsIfExist() {
		messages = append(messages, fmt.Sprintf("Error has occurred on Shobar site https://shobar.com.ua\n\n%s", errorMessage))
	}

	if len(messages) == 0 {
		fmt.Printf("There are no new errors in log file: %s\n", os.Getenv("ERROR_LOG_PATH"))
		return messages
	}

	fmt.Printf("Found new errors in log file: %s\n", os.Getenv("ERROR_LOG_PATH"))
	return messages
}

func getErrorsIfExist() []string {
	errorsContext := utils.FileGetContents(os.Getenv("ERROR_LOG_PATH"))
	var emptyStrings []string

	if errorsContext == "" {
		utils.PutIntoCache(emptyStrings, "errors")
		return emptyStrings
	}

	newErrors := utils.ParseErrors(errorsContext)

	if len(newErrors) == 0 {
		utils.PutIntoCache(emptyStrings, "errors")
		return emptyStrings
	}

	oldErrors := utils.GetCached("errors")
	uniqueErrors := utils.GetUniqueItem(oldErrors, newErrors)

	if newErrors != nil && len(uniqueErrors) > 0 {
		utils.PutIntoCache(newErrors, "errors")
	}

	return uniqueErrors
}

// getReportsIfExist makes request gets data and searches for new Photo Reports
func getReportsIfExist() []models.PhotoReport {
	html := utils.GetHTMLFromTargetURL(os.Getenv("BOT_TARGET_URL"))
	posts := utils.GetPostsFromHTML(&html)
	photoReports := utils.GetAllInformation(&posts)
	titles := utils.GetTitlesFromPhotoReports(photoReports)

	return utils.GenerateMapOfNewData(titles, photoReports)
}
