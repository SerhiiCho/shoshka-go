package utils

import (
	"encoding/json"
	"io/ioutil"
)

// FileGetContents returns the content of the given file
func FileGetContents(filePath string) string {
	fileText, err := ioutil.ReadFile(filePath)

	if err != nil {
		print("File reading error")
		return ""
	}

	return string(fileText)
}

// GetCached returns cache
func GetCached(fileName string) []string {
	var oldSlice []string

	oldContext := FileGetContents("./storage/" + fileName)
	err := json.Unmarshal([]byte(oldContext), &oldSlice)
	HandleError(err, "Unmarshal method returned error")

	return oldSlice
}

// PutTitlesIntoCache insert titles into a file in JSON
func PutTitlesIntoCache(titles []string) {
	file, _ := json.MarshalIndent(titles, "", " ")
	err := ioutil.WriteFile("./storage/titles", file, 0644)
	HandleError(err, "Can't write to a file storage/titles")
}
