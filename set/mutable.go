package set

import (
	"fmt"
	"strings"
)

type MutableSet[T comparable] map[T]struct{}

func (s MutableSet[T]) Contains(item T) bool {
	_, ok := s[item]
	return ok
}

func (s MutableSet[T]) Add(item T) {
	s[item] = struct{}{}
}

func (s MutableSet[T]) Del(item T) {
	for k := range s {
		if k == item {
			delete(s, k)
		}
	}
}

func (s MutableSet[T]) Len() int {
	return len(s)
}

func (s MutableSet[T]) String() string {
	items := make([]string, len(s))
	i := 0
	for k := range s {
		items[i] = fmt.Sprint(k)
		i++
	}
	return fmt.Sprintf("MutableSet(%s)", strings.Join(items, ", "))
}

func NewMutable[T comparable]() MutableSet[T] {
	return MutableSet[T]{}
}

func NewMutableWith[T comparable](values ...T) MutableSet[T] {
	set := NewMutable[T]()
	for _, val := range values {
		set.Add(val)
	}
	return set
}
