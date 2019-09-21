package main

import (
	"io/ioutil"
	"log"

	"github.com/SerhiiCho/shoshka/system"
	// "github.com/SerhiiCho/shoshka/telegram"
)

func init() {
	system.LoadEnvPackage()
}

func main() {
	// telegram.SendMessage("Hello man 2")
	// html := system.GetHTMLFromTargetURL(os.Getenv("BOT_TARGET_URL"))

	// linksHTML := system.GetLinksFromHTML(html)
	// photoReports := system.GetAllInformation(linksHTML)
	msg := []byte("Some text is here")
	err := ioutil.WriteFile("hello.txt", msg, 0644)

	if err != nil {
		log.Fatal(err)
	}
	// // Print all titles
	// for _, photoReport := range photoReports {
	// 	log.Println(photoReport.Title)
	// 	log.Println(photoReport.Image)
	// 	log.Println(photoReport.Title)
	// }
}
