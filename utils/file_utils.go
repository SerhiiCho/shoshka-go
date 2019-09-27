package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

// FileGetContents returns the content of the given file
func FileGetContents(filePath string) string {
	fileText, err := ioutil.ReadFile(filePath)

	if err != nil {
		print("File reading error\n")
		return ""
	}

	return string(fileText)
}

func getAbsPathToStorageFile(filename string) string {
	dirPath, err := filepath.Abs("./storage")
	HandleError(err, "Dir path error in getAbsPathToStorageFile function")
	return dirPath + "/" + filename
}

// GetCached returns cache
func GetCached(fileName string) []string {
	var oldSlice []string

	fmt.Println(getAbsPathToStorageFile(fileName))
	oldContext := FileGetContents(getAbsPathToStorageFile(fileName))
	err := json.Unmarshal([]byte(oldContext), &oldSlice)

	HandleError(err, "Unmarshal method returned error")

	return oldSlice
}

// PutIntoCache insert given slice into a file in JSON
func PutIntoCache(items []string, fileName string) {
	file, _ := json.MarshalIndent(items, "", " ")
	err := ioutil.WriteFile(getAbsPathToStorageFile(fileName), file, 0644)
	HandleError(err, "Can't write to a file storage/"+fileName)
}
