package utilstests

import (
	"github.com/SerhiiCho/shoshka_go/utils"
	"testing"
)

func TestContains(t *testing.T) {
	t.Run("must return true", func(t *testing.T) {
		slice := []string{"hello", "many", "also", "hook"}

		if !utils.Contains(slice, "hello") {
			t.Error("slice must return true instead of false")
		}
	})

	t.Run("must return false", func(t *testing.T) {
		slice := []string{"hello", "many", "also", "hook"}

		if utils.Contains(slice, "anton") {
			t.Error("slice must return false instead of true")
		}
	})
}
