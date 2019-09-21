package utils

import "io/ioutil"

// FilePutContent
func FilePutContent(filePath string, text string) {
	msg := []byte("Some text is here")
	err := ioutil.WriteFile("hello.txt", msg, 0644)
}
