package set

import (
	"regexp"
	"testing"
)

func Test_ImmutableSet_Contains(t *testing.T) {
	originalSet := NewImmutable[string]()

	set := originalSet
	input := []string{"abc", "efg"}
	for _, val := range input {
		set = set.Add(val)
	}
	if !set.Contains(input[0]) {
		t.Errorf("set expected to contain %v", input[0])
	}
	if !set.Contains(input[1]) {
		t.Errorf("set expected to contain %v", input[1])
	}

	if originalSet.Contains(input[0]) {
		t.Errorf("original set not expected to contain %v", input[0])
	}
	if originalSet.Contains(input[1]) {
		t.Errorf("original set expected to contain %v", input[1])
	}
}

func Test_ImmutableSet_Len(t *testing.T) {
	originalSet := NewImmutable[string]()

	set := originalSet
	input := []string{"abc", "efg"}
	for _, val := range input {
		set = set.Add(val)
	}
	if set.Len() != 2 {
		t.Error("should contain only 2 items")
	}

	if originalSet.Len() != 0 {
		t.Error("original set expected to be empty")
	}
}

func Test_ImmutableSet_Add(t *testing.T) {
	originalSet := NewImmutable[string]()

	set := originalSet
	input := []string{"abc", "efg"}
	// adding items twice
	for _, val := range input {
		set = set.Add(val)
	}
	for _, val := range input {
		set = set.Add(val)
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

	if originalSet.Len() != 0 {
		t.Error("original set expected to be empty")
	}
}

func Test_ImmutableSet_Del(t *testing.T) {
	originalSet := NewImmutable[string]()

	set := originalSet
	input := []string{"abc", "efg", "hijk"}

	for _, val := range input {
		set = set.Add(val)
	}
	set = set.Del("efg")

	if set.Len() != 2 {
		t.Error("should contain only 2 items")
	}
	if set.Contains(input[1]) {
		t.Errorf("set expected not to contain %v", input[1])
	}

	if originalSet.Len() != 0 {
		t.Error("original set expected to be empty")
	}
}

func Test_ImmutableSet_String(t *testing.T) {
	set := NewImmutable[string]()
	input := []string{"abc", "efg", "hijk"}

	for _, val := range input {
		set = set.Add(val)
	}

	expectedPattern := `ImmutableSet(.+, .+, .+)`
	matched, err := regexp.MatchString(expectedPattern, set.String())
	if err != nil {
		t.Fatal(err)
	}
	if !matched {
		t.Errorf("string expected to match the pattern %q, got %q", set.String(), expectedPattern)
	}
}

func Test_ImmutableSet_Keys(t *testing.T) {
	input := []string{"abc", "efg", "hijk"}
	set := NewImmutableWith(input...)

	if len(set.Keys()) != len(input) {
		t.Errorf("keys expected %s, got %s", input, set.Keys())
	}
}

func Test_ImmutableSet_NewImmutableWith(t *testing.T) {
	input := []string{"abc", "efg", "hijk"}
	set := NewImmutableWith(input...)

	expectedPattern := `ImmutableSet(.+, .+, .+)`
	matched, err := regexp.MatchString(expectedPattern, set.String())
	if err != nil {
		t.Fatal(err)
	}
	if !matched {
		t.Errorf("string expected to match the pattern %q, got %q", set.String(), expectedPattern)
	}
}
