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

func TestGetUniqueItem(t *testing.T) {
	slice1 := []string{"hello", "another", "nice", "cool"}
	slice2 := []string{"hello", "another", "nice", "cool", "unique"}
	expect := []string{"unique"}

	t.Run("returns unique item", func(t *testing.T) {
		result := utils.GetUniqueItem(slice2, slice1)

		if result[0] != expect[0] {
			t.Errorf("result must be %v", expect)
		}
	})

	t.Run("returns unique item with different order of args", func(t *testing.T) {
		result := utils.GetUniqueItem(slice1, slice2)

		if result[0] != expect[0] {
			t.Errorf("result must be %v", expect)
		}
	})

	t.Run("returns nil if no diff", func(t *testing.T) {
		sameSlice := []string{"hello", "another", "nice", "cool"}
		result := utils.GetUniqueItem(slice1, sameSlice)

		if result != nil {
			t.Error("result must return nil")
		}
	})
}
