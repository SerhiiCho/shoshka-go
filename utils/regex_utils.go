package utils

import (
	"regexp"

	"github.com/SerhiiCho/shoshka-go/models"
)

// GetLinksFromHTML returns all html links
func GetLinksFromHTML(html string) []string {
	linksRegex := "<a href=[\"'](.*?)[\"']\\sdata-rel=[\"']slide-.*?[\"']\\sclass=[\"']slide-image[\"']\\s.*?>.*?<\\/a>"
	return regexp.MustCompile(linksRegex).FindAllString(html, -1)
}

// GetAllInformation parses anchor tags and takes image src,
// link url and title.
func GetAllInformation(html []string) []models.PhotoReport {
	var photoReport models.PhotoReport
	var result []models.PhotoReport

	for _, tag := range html {
		photoReport.Title = regexp.MustCompile("alt=\"([\\d\\W\\D\\s\\S]+)\"").FindStringSubmatch(tag)[1]
		photoReport.Image = regexp.MustCompile("\\ssrc=\"([\\S]+)\"\\s").FindStringSubmatch(tag)[1]
		photoReport.URL = regexp.MustCompile("\\shref=(\"|')([\\S]+)(\"|')\\s").FindStringSubmatch(tag)[2]

		result = append(result, photoReport)
	}

	return result
}
