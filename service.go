package designaspect

import (
	"goa.design/goa/v3/dsl"
	"goa.design/goa/v3/expr"
)

type ServiceFunc func(*Srvc)

func Service(name string, f ServiceFunc) *expr.ServiceExpr {
	s := newSrvc()
	return dsl.Service(name, func() {
		f(s)
		s.Elements.InService()
	})
}

type Srvc struct {
	Elements Elements
}

func newSrvc() *Srvc {
	return &Srvc{Elements: Elements{}}
}

func (s *Srvc) Use(elements ...Element) {
	s.Elements = append(s.Elements, elements...)
}

func (s *Srvc) Method(name string, f MethodFunc) {
	methodWithService(s, name, f)
}

func (s *Srvc) HTTP(funcs ...func()) {
	dsl.HTTP(func() {
		s.Elements.InHTTP()
		for _, f := range funcs {
			f()
		}
	})
}
