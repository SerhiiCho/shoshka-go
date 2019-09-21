package utils

import "log"

// HandleError checks error and logs it if there is one
func HandleError(err error, message string) {
	if err != nil {
		log.Fatal(message, err.Error())
	}
}
