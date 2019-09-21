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
	utils.HandleError(err, "Can't write to a file")
}

func contains(slice []string, needle string) bool {
	for _, sliceItem := range slice {
		if sliceItem == needle {
			return true
		}
	}
	return false
}

func getUniqueItem(slice1 []string, slice2 []string) []string {
	result := []string{}

	for _, item1 := range slice1 {
		if !contains(slice2, item1) {
			result = append(result, item1)
		}
	}

	return result
}

func main() {
	// telegram.SendMessage("Hello man 2")
	html := utils.GetHTMLFromTargetURL(os.Getenv("BOT_TARGET_URL"))

	linksHTML := utils.GetLinksFromHTML(html)
	photoReports := utils.GetAllInformation(linksHTML)

	titles := []string{}

	for _, photoReport := range photoReports {
		titles = append(titles, photoReport.Title)
		// log.Println(photoReport.Image)
		// log.Println(photoReport.Title)
	}

	oldContext := utils.FileGetContents("./storage/titles")
	oldSlice := []string{}
	json.Unmarshal([]byte(oldContext), &oldSlice)
	// saveTitlesIntoFile(titles)
	diff := getUniqueItem(titles, oldSlice)

	if len(diff) == 0 {
		println("No diff")
		return
	}

	fmt.Printf("%#v\n", diff)
}
