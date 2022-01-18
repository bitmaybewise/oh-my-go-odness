package set

import (
	"fmt"
	"strings"
)

type ImmutableSet[T comparable] map[T]bool

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
		newset[k] = true
	}
	newset[item] = true
	return newset
}

func (s ImmutableSet[T]) Del(item T) ImmutableSet[T] {
	newset := NewImmutable[T]()
	for k := range s {
		if k == item {
			continue
		}
		newset[k] = true
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

func NewImmutable[T comparable]() ImmutableSet[T] {
	return ImmutableSet[T]{}
}
