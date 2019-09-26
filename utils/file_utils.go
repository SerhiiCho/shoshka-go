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

// PutIntoCache insert given slice into a file in JSON
func PutIntoCache(items []string, fileName string) {
	file, _ := json.MarshalIndent(items, "", " ")
	err := ioutil.WriteFile("./storage/"+fileName, file, 0644)
	HandleError(err, "Can't write to a file storage/"+fileName)
}
