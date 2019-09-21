package system

import (
	"io/ioutil"
	"net/http"

	"github.com/SerhiiCho/shoshka/utils"
)

/*
GetHTMLFromTargetURL makes request to a web page
and returns html string
*/
func GetHTMLFromTargetURL(url string) string {
	resp, err := http.Get(url)

	utils.HandleError(err, "Can't connect to "+url)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	utils.HandleError(err, "Can't read body")

	return string(body)
}
