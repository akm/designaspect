package designaspect

import (
	"goa.design/goa/v3/dsl"
	"goa.design/goa/v3/expr"
)

type HttpMethod func(string) *expr.RouteExpr

type Method struct {
	Name           string
	Path           string
	BuildPayload   func(*Method) interface{}
	ResultType     *expr.ResultTypeExpr
	HttpMethod     HttpMethod
	HttpStatusCode int
	Elements       Elements
}

func (x *Method) Define() {
	dsl.Method(x.Name, func() {
		x.Elements.InMethod()
		dsl.Payload(x.BuildPayload(x))
		if x.ResultType != nil {
			dsl.Result(x.ResultType)
		}
		dsl.HTTP(func() {
			x.HttpMethod(x.Path)
			dsl.Response(x.HttpStatusCode)
			x.Elements.InHTTP()
		})
	})
}
