package set

import (
	"regexp"
	"testing"
)

func Test_MutableSet_Contains(t *testing.T) {
	set := NewMutable[string]()
	input := []string{"abc", "efg"}
	for _, val := range input {
		set.Add(val)
	}
	if !set.Contains(input[0]) {
		t.Errorf("set expected to contain %v", input[0])
	}
	if !set.Contains(input[1]) {
		t.Errorf("set expected to contain %v", input[1])
	}
}

func Test_MutableSet_Len(t *testing.T) {
	set := NewMutable[string]()
	input := []string{"abc", "efg"}
	for _, val := range input {
		set.Add(val)
	}
	if set.Len() != 2 {
		t.Error("should contain only 2 items")
	}
}

func Test_MutableSet_Add(t *testing.T) {
	set := NewMutable[string]()
	input := []string{"abc", "efg"}

	// adding items twice
	for _, val := range input {
		set.Add(val)
	}
	for _, val := range input {
		set.Add(val)
	}

	if set.Len() != 2 {
		t.Error("should contain only 2 items")
	}
	if !set.Contains(input[0]) {
		t.Errorf("set expected to contain %v", input[0])
	}
	if !set.Contains(input[1]) {
		t.Errorf("set expected to contain %v", input[1])
	}
}

func Test_MutableSet_Del(t *testing.T) {
	set := NewMutable[string]()
	input := []string{"abc", "efg", "hijk"}

	for _, val := range input {
		set.Add(val)
	}
	set.Del("efg")

	if set.Len() != 2 {
		t.Error("should contain only 2 items")
	}
	if set.Contains(input[1]) {
		t.Errorf("set expected not to contain %v", input[1])
	}
}

func Test_MutableSet_String(t *testing.T) {
	input := []string{"abc", "efg", "hijk"}
	set := NewMutableWith(input...)

	expectedPattern := `MutableSet(.+, .+, .+)`
	matched, err := regexp.MatchString(expectedPattern, set.String())
	if err != nil {
		t.Fatal(err)
	}
	if !matched {
		t.Errorf("string expected to match the pattern %q, got %q", set.String(), expectedPattern)
	}
}

func Test_MutableSet_Keys(t *testing.T) {
	input := []string{"abc", "efg", "hijk"}
	set := NewMutableWith(input...)

	if len(set.Keys()) != len(input) {
		t.Errorf("keys expected %s, got %s", input, set.Keys())
	}
}

func Test_MutableSet_NewMutableWith(t *testing.T) {
	input := []string{"abc", "efg", "hijk"}
	set := NewMutableWith(input...)

	expectedPattern := `MutableSet(.+, .+, .+)`
	matched, err := regexp.MatchString(expectedPattern, set.String())
	if err != nil {
		t.Fatal(err)
	}
	if !matched {
		t.Errorf("string expected to match the pattern %q, got %q", set.String(), expectedPattern)
	}
}
