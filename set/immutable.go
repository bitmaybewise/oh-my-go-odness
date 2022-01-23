package set

import (
	"fmt"
	"strings"
)

type ImmutableSet[T comparable] map[T]struct{}

func (s ImmutableSet[T]) Contains(item T) bool {
	_, ok := s[item]
	return ok
}

func (s ImmutableSet[T]) Len() int {
	return len(s)
}

func (s ImmutableSet[T]) Add(item T) ImmutableSet[T] {
	if _, ok := s[item]; ok {
		return s
	}
	newset := NewImmutable[T]()
	for k := range s {
		newset[k] = struct{}{}
	}
	newset[item] = struct{}{}
	return newset
}

func (s ImmutableSet[T]) Del(item T) ImmutableSet[T] {
	newset := NewImmutable[T]()
	for k := range s {
		if k == item {
			continue
		}
		newset[k] = struct{}{}
	}
	return newset
}

func (s ImmutableSet[T]) String() string {
	items := make([]string, len(s))
	i := 0
	for k := range s {
		items[i] = fmt.Sprint(k)
		i++
	}
	return fmt.Sprintf("ImmutableSet(%s)", strings.Join(items, ", "))
}

func (s ImmutableSet[T]) Keys() []T {
	keys := make([]T, len(s))
	i := 0
	for k := range s {
		keys[i] = k
		i++
	}
	return keys
}

func NewImmutable[T comparable]() ImmutableSet[T] {
	return ImmutableSet[T]{}
}

func NewImmutableWith[T comparable](values ...T) ImmutableSet[T] {
	set := NewImmutable[T]()
	for _, val := range values {
		set = set.Add(val)
	}
	return set
}
