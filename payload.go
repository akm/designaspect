package designaspect

import (
	"goa.design/goa/v3/dsl"
)

type PayloadFunc func(*Payload)

func FilterPayloadFunc(s []interface{}) ([]PayloadFunc, []interface{}) {
	funcs := []PayloadFunc{}
	rest := []interface{}{}
	for _, i := range s {
		switch v := i.(type) {
		case PayloadFunc:
			funcs = append(funcs, v)
		case func(*Payload):
			funcs = append(funcs, v)
		default:
			rest = append(rest, i)
		}
	}
	return funcs, rest
}

type Payload struct {
	Method        *Method
	RequiredNames []string
}

func NewPayload(method *Method) *Payload {
	return &Payload{Method: method, RequiredNames: []string{}}
}

func (p *Payload) Attribute(name string, args ...interface{}) string {
	dsl.Attribute(name, args...)
	return name
}

func (p *Payload) Required(names ...string) {
	p.RequiredNames = append(p.RequiredNames, names...)
}

func (p *Payload) Done() {
	names := append([]string{}, p.RequiredNames...)
	elementsAll := p.Method.ElementsAll()
	for _, el := range elementsAll {
		if x, ok := el.(AttributeElement); ok {
			if x.IsRequired() {
				names = append(names, x.GetName())
			}
		}
	}
	if len(names) > 0 {
		dsl.Required(names...)
	}
}
