package set_test

import (
	"testing"

	"github.com/hlmerscher/oh-my-go-odness/set"
)

func Test_Set_Empty(t *testing.T) {
	empty := set.Empty[string, struct{}]()
	if len(empty) > 0 {
		t.Errorf("set expected to be empty, got %s", empty)
	}
}

func Test_Set_Keys(t *testing.T) {
	s := map[string]struct{}{
		"a":   {},
		"abc": {},
		"b":   {},
		"123": {},
	}
	keys := set.Keys(s)
	if len(keys) != 4 {
		t.Errorf("keys has an unexpected size, got %s", keys)
	}
}

func Test_Set_Union(t *testing.T) {
	t.Run("returns empty union with empty params", func(t *testing.T) {
		union := set.Union[string, bool]()
		if len(union) != 0 {
			t.Errorf("union expected to be empty, got %s", union)
		}
	})
	t.Run("returns a new map with keys from sets", func(t *testing.T) {
		a := map[string]struct{}{"abc": {}, "efg": {}, "123": {}}
		b := map[string]struct{}{"abc": {}, "456": {}, "xyz": {}}
		union := set.Union(a, b)
		if len(union) != 5 {
			t.Errorf("union expected to have 5 items, got %s", union)
		}
	})
}

func Test_Set_Intersection(t *testing.T) {
	t.Run("intersection of nothing should be empty", func(t *testing.T) {
		intersection := set.Intersection[int, struct{}]()
		if len(intersection) != 0 {
			t.Errorf("intersection expected to be empty, got %v", intersection)
		}
	})

	t.Run("intersection of one set should be a copy of itself", func(t *testing.T) {
		a := map[string]struct{}{"abc": {}, "efg": {}, "123": {}}
		intersection := set.Intersection(a)
		if len(intersection) != len(a) {
			t.Errorf("intersection expected to be equal %v, got %v", a, intersection)
		}
		if &intersection == &a {
			t.Errorf("intersection expected to different than %p, got %p", a, intersection)
		}
	})

	t.Run("intersection among 3 sets", func(t *testing.T) {
		a := map[string]struct{}{"abc": {}, "efg": {}, "123": {}, "xyz": {}}
		b := map[string]struct{}{"abc": {}, "456": {}, "xyz": {}}
		c := map[string]struct{}{"abc": {}}
		intersection := set.Intersection(a, b, c)
		if len(intersection) != 2 {
			t.Errorf("intersection expected to have 2 items, got %s", intersection)
		}
	})
}
