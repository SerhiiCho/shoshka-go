package main

import (
	"log"
	"os"

	"github.com/SerhiiCho/shoshka-go/models"
	"github.com/SerhiiCho/shoshka-go/telegram"
	"github.com/SerhiiCho/shoshka-go/utils"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	utils.HandleError(err, "Error loading .env file")
}

func main() {
	html := utils.GetHTMLFromTargetURL(os.Getenv("BOT_TARGET_URL"))
	linksHTML := utils.GetLinksFromHTML(html)
	photoReports := utils.GetAllInformation(linksHTML)
	titles := getTitlesFromPhotoReports(photoReports)
	tgMessageData := utils.GenerateMapOfNewData(titles, photoReports)

	exitIfNoNewItems(tgMessageData)

	telegram.SendMessagesWithNewPhotoReports(tgMessageData)
	utils.PutTitlesIntoCache(titles)
}

func exitIfNoNewItems(tgMessageData []models.PhotoReport) {
	if tgMessageData == nil {
		log.Print("There are not new photo reports")
		os.Exit(1)
	}
}

func getTitlesFromPhotoReports(photoReports []models.PhotoReport) []string {
	var titles []string

	for _, photoReport := range photoReports {
		titles = append(titles, photoReport.Title)
	}

	return titles
}
