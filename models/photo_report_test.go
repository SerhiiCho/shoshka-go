package models

import (
	"testing"
)

func TestPhotoReport_can_hold_information(t *testing.T) {
	photoReport := PhotoReport{
		Title: "Hello world",
		Image: "http://test.com",
		URL:   "http://url.com",
	}

	if photoReport.Title != "Hello world" {
		t.Error("PhotoReport model doesn't have value in its Title field")
	}

	if photoReport.Image != "http://test.com" {
		t.Error("PhotoReport model doesn't have value in its Image field")
	}

	if photoReport.URL != "http://url.com" {
		t.Error("PhotoReport model doesn't have value in its URL field")
	}
}
