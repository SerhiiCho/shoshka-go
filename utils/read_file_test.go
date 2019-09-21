package utils_test

import (
	"testing"

	"github.com/SerhiiCho/shoshka_go/utils"
)

func TestFileGetContents_returns_file_content(t *testing.T) {
	result := utils.FileGetContents("../test_files/small_text")

	if result != "Hello world" {
		t.Error("The value of result varialbe must be 'Hello world'")
	}
}

func TestFileGetContents_returns_empty_string_if_file_is_empty(t *testing.T) {
	result := utils.FileGetContents("../test_files/empty")

	if result != "" {
		t.Error("The value of result varialbe must be empty string")
	}
}

func TestFileGetContents_returns_empty_string_if_doesnt_exist(t *testing.T) {
	result := utils.FileGetContents("lfkasdjflkdsjlfkajdkf")

	if result != "" {
		t.Error("FileGetContents must return empty string because file doent exist")
	}
}
