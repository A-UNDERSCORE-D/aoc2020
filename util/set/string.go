package set

// Generated automatically. DO NOT EDIT
import "awesome-dragon.science/go/adventofcode2020/util"

type StringSet struct {
	internalMap map[string]struct{}
}

func NewStringSet() *StringSet {
	return &StringSet{make(map[string]struct{})}
}

func NewStringSetWithLength(length int) *StringSet {
	return &StringSet{make(map[string]struct{}, length)}
}

func (s *StringSet) Contains(v string) bool {
	_, ok := s.internalMap[v]
	return ok
}

func (s *StringSet) Insert(v string) {
	s.internalMap[v] = struct{}{}
}

func (s *StringSet) InsertMany(vs ...string) {
	for _, v := range vs {
		s.internalMap[v] = struct{}{}
	}
}

func (s *StringSet) Remove(value string) {
	delete(s.internalMap, value)
}

func (s *StringSet) Intersect(other *StringSet) *StringSet {
	out := NewStringSetWithLength(util.Min(s.Length(), other.Length()))
	for v := range s.internalMap {
		if other.Contains(v) {
			out.Insert(v)
		}
	}
	return out
}

func (s *StringSet) Union(other *StringSet) *StringSet {
	out := NewStringSetWithLength(s.Length() + other.Length())
	out.InsertMany(s.Values()...)
	out.InsertMany(other.Values()...)
	return out
}

func (s *StringSet) Difference(other *StringSet) *StringSet {
	out := NewStringSetWithLength(util.Min(s.Length(), other.Length()))
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

func (s *StringSet) Values() []string {
	out := make([]string, 0, len(s.internalMap))
	for v := range s.internalMap {
		out = append(out, v)
	}
	return out
}

func (s *StringSet) Length() int {
	return len(s.internalMap)
}
