package errors

import (
	"github.com/akm/designaspect"
)

type Mappings []*Mapping

func (s Mappings) Select(f func(*Mapping) bool) Mappings {
	r := Mappings{}
	for _, i := range s {
		if f(i) {
			r = append(r, i)
		}
	}
	return r
}

func containString(s []string, t string) bool {
	for _, i := range s {
		if i == t {
			return true
		}
	}
	return false
}

func (s Mappings) SelectByNames(args ...string) Mappings {
	return s.Select(func(i *Mapping) bool {
		return containString(args, i.Name)
	})
}

func (s Mappings) Elements() designaspect.Elements {
	r := make(designaspect.Elements, len(s))
	for idx, i := range s {
		r[idx] = i
	}
	return r
}

func (s Mappings) Define() {
	for _, i := range s {
		i.Define()
	}
}
