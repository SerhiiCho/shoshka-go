package utils

import (
	"testing"
)

func TestGetLinksFromHTML(t *testing.T) {
	t.Run("returns links", func(t *testing.T) {
		fileText := FileGetContents("../storage/test_files/example")
		result := GetLinksFromHTML(fileText)

		if len(result) != 8 {
			t.Error("GetLinksFromHTML must return slice with 8 items")
		}
	})
}

func TestGetAllInformation(t *testing.T) {
	t.Run("returns slice with PhotoReport models", func(t *testing.T) {
		links := []string{FileGetContents("../storage/test_files/link")}

		title := "Title here"
		img := "https://image.jpg"
		url := "https://test.com/hello/"

		result := GetAllInformation(links)

		if result[0].Title != title {
			t.Error("Returned Title from GetAllInformation func must be `" + title + "` but `" + result[0].Title + "` returned")
		}

		if result[0].Image != img {
			t.Error("Returned Image from GetAllInformation func must be `" + img + "` but `" + result[0].Image + "` returned")
		}

		if result[0].URL != url {
			t.Error("Returned URL from GetAllInformation func must be `" + url + "` but `" + result[0].URL + "` returned")
		}
	})
}

func TestParseErrors(t *testing.T) {
	errorsContext := FileGetContents("../storage/test_files/error_log")
	result := ParseErrors(errorsContext)

	if len(result) != 2 {
		t.Errorf("Result must returns slice with 2 items but %v returned", len(result))
	}
}
