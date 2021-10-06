package errors

import (
	"goa.design/goa/v3/dsl"
	"goa.design/goa/v3/expr"
)

type Mapping struct {
	Name       string
	StatusCode int
	ErrorMsg   string
	ErrorType  interface{}
}

func NewMapping(name string, statusCode int, errorMsg string) *Mapping {
	return &Mapping{
		Name:       name,
		StatusCode: statusCode,
		ErrorMsg:   errorMsg,
		ErrorType:  expr.ErrorResult,
	}
}

func (d *Mapping) Type(t interface{}) *Mapping {
	d.ErrorType = t
	return d
}

func (d *Mapping) Define() {
	dsl.Error(d.Name, d.ErrorType, d.ErrorMsg)
}

func (d *Mapping) Response() {
	dsl.Response(d.Name, d.StatusCode)
}

func (d *Mapping) InService() {
	d.Define()
}

func (d *Mapping) InMethod() {
	d.Define()
}

func (d *Mapping) InPayload() {
}

func (d *Mapping) InHTTP() {
	d.Response()
}
