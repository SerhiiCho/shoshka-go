package utils

import (
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
