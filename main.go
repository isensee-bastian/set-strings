package main

import (
	"sort"
	"strings"
)

type Set struct {
	internalMap map[string]bool
}

func NewSet(elements ...string) Set {
	internalMap := map[string]bool{}

	for _, element := range elements {
		internalMap[element] = true
	}

	return Set{internalMap}
}

func (s Set) Size() int {
	return len(s.internalMap)
}

func (s *Set) Add(element string) {
	s.internalMap[element] = true
}

func (s *Set) Remove(element string) {
	delete(s.internalMap, element)
}

func (s Set) Contains(element string) bool {
	// Since we only store true values for keys that are existing,
	// we can just return the value. Missing key will result in default
	// value which is false for bool types.
	return s.internalMap[element]
}

func (s Set) Slice() []string {
	slice := make([]string, 0, len(s.internalMap))

	for element := range s.internalMap {
		slice = append(slice, element)
	}

	return slice
}

// Return elements ordered ascending.
func (s Set) String() string {
	slice := s.Slice()
	sort.Strings(slice)

	return strings.Join(slice, ", ")
}
