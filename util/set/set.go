// Package set contains a map based set implementation
package set

import "awesome-dragon.science/go/adventofcode2020/util"

type Set struct {
	internalMap map[interface{}]struct{}
}

func NewSet() *Set {
	return &Set{make(map[interface{}]struct{})}
}

func NewSetWithLength(length int) *Set {
	return &Set{make(map[interface{}]struct{}, length)}
}

func (s *Set) Contains(v interface{}) bool {
	_, ok := s.internalMap[v]
	return ok
}

func (s *Set) Insert(v interface{}) {
	s.internalMap[v] = struct{}{}
}

func (s *Set) InsertMany(vs ...interface{}) {
	for _, v := range vs {
		s.internalMap[v] = struct{}{}
	}
}

func (s *Set) Intersect(other *Set) *Set {
	out := NewSetWithLength(util.Min(s.Length(), other.Length()))
	for v := range s.internalMap {
		if other.Contains(v) {
			out.Insert(v)
		}
	}
	return out
}

func (s *Set) Union(other *Set) *Set {
	out := NewSetWithLength(s.Length() + other.Length())
	out.InsertMany(s.Values()...)
	out.InsertMany(other.Values()...)
	return out
}

func (s *Set) Difference(other *Set) *Set {
	out := NewSetWithLength(util.Min(s.Length(), other.Length()))
	for v := range s.internalMap {
		if !other.Contains(v) {
			out.Insert(v)
		}
	}

	for v := range other.internalMap {
		if !s.Contains(v) {
			out.Insert(v)
		}
	}
	return out
}

func (s *Set) Values() []interface{} {
	out := make([]interface{}, 0, len(s.internalMap))
	for v := range s.internalMap {
		out = append(out, v)
	}
	return out
}

func (s *Set) Length() int {
	return len(s.internalMap)
}
