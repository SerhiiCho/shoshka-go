package utils

import (
	"regexp"

	"github.com/SerhiiCho/shoshka-go/models"
)

// GetPostsFromHTML returns all html posts
func GetPostsFromHTML(html string) []string {
	linksRegex := "<article class(.|\n)*?<\\/article>"
	return regexp.MustCompile(linksRegex).FindAllString(html, -1)
}

// GetAllInformation parses anchor tags and takes image src,
// link url and title.
func GetAllInformation(html []string) []models.PhotoReport {
	var photoReport models.PhotoReport
	var result []models.PhotoReport

	for _, tag := range html {
		photoReport.Title = regexp.MustCompile("<h3 (.|\\n)*?[^ ]>(.*)<\\/a>").FindStringSubmatch(tag)[2]
		photoReport.Image = regexp.MustCompile(" src=\"([\\S]+)\" ").FindStringSubmatch(tag)[1]
		photoReport.URL = regexp.MustCompile(" href=[\"|']([\\S]+)(\"|') ").FindStringSubmatch(tag)[1]

		result = append(result, photoReport)
	}

	return result
}

// ParseErrors returns slice of errors
func ParseErrors(text string) []string {
	return regexp.MustCompile("\\[([A-z0-9 :-]+)\\] (.*)").FindAllString(text, -1)
}
