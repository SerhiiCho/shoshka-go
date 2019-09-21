package main

import (
	"fmt"
	"github.com/SerhiiCho/shoshka_go/utils"
	"github.com/joho/godotenv"
	"log"
	"os"
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

	var titles []string

	telegram.SendMessagesWithNewPhotoReports(tgMessageData)
	utils.PutTitlesIntoCache(titles)
}

	doesntHaveNewItems, tgMessageData := utils.GenerateMapOfNewData(titles, photoReports)

	if doesntHaveNewItems {
		log.Print("There are not new photo reports")
		return
	}

	fmt.Printf("%#v\n", tgMessageData)

	// 2. send message with image, title and link
	// telegram.SendMessage("Hello man 2")

	//utils.PutTitlesIntoCache(titles)
}
