package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/SerhiiCho/shoshka_go/utils"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	utils.HandleError(err, "Error loading .env file")
}

func saveTitlesIntoFile(titles []string) {
	file, _ := json.MarshalIndent(titles, "", " ")
	err := ioutil.WriteFile("./storage/titles", file, 0644)
	utils.HandleError(err, "Can't write to a file storage/titles")
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

	oldContext := utils.FileGetContents("./storage/titles")
	var oldSlice []string
	err := json.Unmarshal([]byte(oldContext), &oldSlice)

	utils.HandleError(err, "Unmarshal method returned error")

	diff := utils.GetUniqueItem(titles, oldSlice)

	if len(diff) == 0 {
		println("No diff")
		return
	}

	fmt.Printf("%#v\n", diff)
	//saveTitlesIntoFile(titles)
}
