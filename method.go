package designaspect

import (
	"goa.design/goa/v3/dsl"
)

type MethodFunc func(*Method)

func methodWithService(s *Srvc, name string, f MethodFunc) {
	m := newMethod(s)
	dsl.Method(name, func() {
		f(m)
		m.Elements.InMethod()
	})
}

type Method struct {
	Srvc     *Srvc
	Elements Elements
}

func newMethod(s *Srvc) *Method {
	return &Method{Srvc: s, Elements: Elements{}}
}

func (m *Method) Use(elements ...Element) {
	m.Elements = append(m.Elements, elements...)
}

func (m *Method) ElementsAll() Elements {
	return append(m.Srvc.Elements, m.Elements...)
}

func (m *Method) Payload(funcs ...PayloadFunc) {
	dsl.Payload(func() {
		m.ElementsAll().InPayload()
		p := NewPayload(m)
		for _, f := range funcs {
			f(p)
		}
		p.Done()
	})
}

func (m *Method) HTTP(funcs ...func()) {
	dsl.HTTP(func() {
		m.Elements.InHTTP()
		for _, f := range funcs {
			f()
		}
	})
}
