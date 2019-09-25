package utilstests

import (
	"testing"

	"github.com/SerhiiCho/shoshka-go/utils"
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

func TestGetIndexOfSliceItem(t *testing.T) {
	slice := []string{"first", "second", "third"}
	cases := []struct {
		Title  string
		Expect int
		Input  string
	}{
		{"returns -1 if value is not in a slice", -1, "forth"},
		{"returns 1 if value has index 1", 0, "first"},
		{"returns 2 if value has index 2", 1, "second"},
		{"returns 3 if value has index 3", 2, "third"},
	}

	for _, singleCase := range cases {
		t.Run(singleCase.Title, func(t *testing.T) {
			result := utils.GetIndexOfSliceItem(slice, singleCase.Input)

			if result != singleCase.Expect {
				t.Errorf("resulted index must be %d but %d returned", singleCase.Expect, result)
			}
		})
	}
}
