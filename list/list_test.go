package list

import "testing"

func Test_list_Reverse(t *testing.T) {
	list := []any{"abc", 123, true, false}
	reversed := Reverse(list)

	expectedList := []struct {
		idx   int
		value interface{}
	}{
		{0, false},
		{1, true},
		{2, 123},
		{3, "abc"},
	}
	for _, expected := range expectedList {
		item := reversed[expected.idx]
		if item != expected.value {
			t.Errorf("item on index %d expected to be %v, got %v", expected.idx, expected.value, item)
		}
	}
}
