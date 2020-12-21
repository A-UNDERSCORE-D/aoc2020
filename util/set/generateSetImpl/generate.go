package main

import (
	"os"
	"strings"
	"text/template"
)

func main() {
	targetType := os.Args[1]
	setName := strings.Title(targetType) + "Set"

	templ := template.Must(template.New("set").Parse(baseImpl))
	templ.Execute(os.Stdout, struct{ Name, Type string }{setName, targetType})
}

const baseImpl = `
package set
// Generated automatically. DO NOT EDIT
import "awesome-dragon.science/go/adventofcode2020/util"

type {{.Name}} struct {
	internalMap map[{{.Type}}]struct{}
}

func New{{.Name}}() *{{.Name}} {
	return &{{.Name}}{make(map[{{.Type}}]struct{})}
}

func New{{.Name}}WithLength(length int) *{{.Name}} {
	return &{{.Name}}{make(map[{{.Type}}]struct{}, length)}
}

func (s *{{.Name}}) Contains(v {{.Type}}) bool {
	_, ok := s.internalMap[v]
	return ok
}

func (s *{{.Name}}) Insert(v {{.Type}}) {
	s.internalMap[v] = struct{}{}
}

func (s *{{.Name}}) InsertMany(vs ...{{.Type}}) {
	for _, v := range vs {
		s.internalMap[v] = struct{}{}
	}
}

func (s *{{.Name}}) Remove(value {{.Type}}) {
	delete(s.internalMap, value)
}

func (s *{{.Name}}) Intersect(other *{{.Name}}) *{{.Name}} {
	out := New{{.Name}}WithLength(util.Min(s.Length(), other.Length()))
	for v := range s.internalMap {
		if other.Contains(v) {
			out.Insert(v)
		}
	}
	return out
}

func (s *{{.Name}}) Union(other *{{.Name}}) *{{.Name}} {
	out := New{{.Name}}WithLength(s.Length() + other.Length())
	out.InsertMany(s.Values()...)
	out.InsertMany(other.Values()...)
	return out
}

func (s *{{.Name}}) Difference(other *{{.Name}}) *{{.Name}} {
	out := New{{.Name}}WithLength(util.Min(s.Length(), other.Length()))
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

func (s *{{.Name}}) Values() []{{.Type}} {
	out := make([]{{.Type}}, 0, len(s.internalMap))
	for v := range s.internalMap {
		out = append(out, v)
	}
	return out
}

func (s *{{.Name}}) Length() int {
	return len(s.internalMap)
}
`
