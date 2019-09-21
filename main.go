package main

import (
	"encoding/json"
	"os"

	"github.com/SerhiiCho/shoshka_go/system"
	"github.com/SerhiiCho/shoshka_go/utils"
)

func init() {
	system.LoadEnvPackage()
}

func saveTitlesToFile(titles []string) {
	file, _ := json.MarshalIndent(titles, "", " ")
	utils.FilePutContent("", string(file))
}

func main() {
	// telegram.SendMessage("Hello man 2")
	html := system.GetHTMLFromTargetURL(os.Getenv("BOT_TARGET_URL"))

	linksHTML := system.GetLinksFromHTML(html)
	photoReports := system.GetAllInformation(linksHTML)

	var titles []string

	for _, photoReport := range photoReports {
		titles = append(titles, photoReport.Title)
		// log.Println(photoReport.Image)
		// log.Println(photoReport.Title)
	}

	saveTitlesToFile(titles)
}
