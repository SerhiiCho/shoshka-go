package utils_test

import (
	"testing"

	"github.com/SerhiiCho/shoshka/utils"
)

func TestFilePutContents_saves_text_to_a_file(t *testing.T) {
	text := "Test text"
	filePath := "../test_files/some_text"
	utils.FilePutContent(filePath, text)

	result = utils.FileGetContents(filePath)

	if result != text {
		t.Error("FileGetContents hasn't saved the text to a file")
	}
}
