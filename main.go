package main

import (
	"fmt"
	"github.com/SerhiiCho/shoshka_go/utils"
	"github.com/joho/godotenv"
	"os"
)

func init() {
	err := godotenv.Load()
	utils.HandleError(err, "Error loading .env file")
}

func main() {
	// telegram.SendMessage("Hello man 2")
	html := utils.GetHTMLFromTargetURL(os.Getenv("BOT_TARGET_URL"))

	linksHTML := utils.GetLinksFromHTML(html)
	photoReports := utils.GetAllInformation(linksHTML)

	var titles []string

	for _, photoReport := range photoReports {
		titles = append(titles, photoReport.Title)
		// log.Println(photoReport.Image)
		// log.Println(photoReport.Title)
	}

	oldTitles := utils.GetCachedTitles()
	diff := utils.GetUniqueItem(titles, oldTitles)

	if len(diff) == 0 {
		println("No diff")
		return
	}

	fmt.Printf("%#v\n", diff)
	utils.PutTitlesIntoCache(titles)
}
