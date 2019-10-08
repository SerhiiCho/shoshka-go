package utils

import (
	"testing"
)

func TestFileGetContents(t *testing.T) {
	t.Run("returns file content", func(t *testing.T) {
		result := FileGetContents("../storage/test_files/small_text")

		if result != "Hello world" {
			t.Errorf("The value of result variable must be `Hello world` but %v returned", result)
		}
	})

	t.Run("returns empty string if file is empty", func(t *testing.T) {
		result := FileGetContents("../storage/test_files/empty")

		if result != "" {
			t.Error("The value of result variable must be empty string")
		}
	})

	t.Run("returns empty string if does not exist", func(t *testing.T) {
		result := FileGetContents("nostradamus")

		if result != "" {
			t.Error("FileGetContents must return empty string because file doesn't exist")
		}
	})
}
