package set

import (
	"regexp"
	"testing"
)

func Test_Mutable_Contains(t *testing.T) {
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

func Test_Mutable_Len(t *testing.T) {
	set := NewMutable[string]()
	input := []string{"abc", "efg"}
	for _, val := range input {
		set.Add(val)
	}
	if set.Len() != 2 {
		t.Error("should contain only 2 items")
	}
}

func Test_Mutable_Add(t *testing.T) {
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

func Test_Mutable_Del(t *testing.T) {
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

func Test_Mutable_String(t *testing.T) {
	set := NewMutable[string]()
	input := []string{"abc", "efg", "hijk"}

	for _, val := range input {
		set.Add(val)
	}

	expectedPattern := `MutableSet(.+, .+, .+)`
	matched, err := regexp.MatchString(expectedPattern, set.String())
	if err != nil {
		t.Fatal(err)
	}
	if !matched {
		t.Errorf("string expected to match the pattern %q, got %q", set.String(), expectedPattern)
	}
}
