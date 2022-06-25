package slice_test

import (
	"testing"

	"github.com/hlmerscher/oh-my-go-odness/slice"
)

func Test_slice_Reverse(t *testing.T) {
	list := []any{"abc", 123, true, false}
	reversed := slice.Reverse(list)

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

func Test_slice_Includes(t *testing.T) {
	s := []string{"aaa", "bbb", "ccc", "abc"}
	value := slice.Includes(s, "123")
	if value {
		t.Errorf("item %q expected not to be found on slice %q", "123", s)
	}
	value = slice.Includes(s, "aaa")
	if !value {
		t.Errorf("item %q expected to be found on slice %q", "123", s)
	}
}
