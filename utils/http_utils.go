package utils

import (
	"io/ioutil"
	"net/http"
)

// GetHTMLFromTargetURL makes request to a web page
// and returns html string
func GetHTMLFromTargetURL(url string) string {
	resp, err := http.Get(url)
	HandleError(err, "Can't connect to "+url)

	body, err := ioutil.ReadAll(resp.Body)
	HandleError(err, "Can't read body")

	err = resp.Body.Close()
	HandleError(err, "Error while trying to close body")

	return string(body)
}
