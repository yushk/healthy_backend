// Code generated by go-swagger; DO NOT EDIT.

package monitor

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	v1 "gitee.com/healthy/backend/apiserver/v1"
	"github.com/go-openapi/runtime/middleware"
)

// GetClassHandlerFunc turns a function with the right signature into a get class handler
type GetClassHandlerFunc func(GetClassParams, *v1.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GetClassHandlerFunc) Handle(params GetClassParams, principal *v1.Principal) middleware.Responder {
	return fn(params, principal)
}

// GetClassHandler interface for that can handle valid get class params
type GetClassHandler interface {
	Handle(GetClassParams, *v1.Principal) middleware.Responder
}

// NewGetClass creates a new http.Handler for the get class operation
func NewGetClass(ctx *middleware.Context, handler GetClassHandler) *GetClass {
	return &GetClass{Context: ctx, Handler: handler}
}

/* GetClass swagger:route GET /v1/class/{id} monitor getClass

获取班级信息

获取班级信息

*/
type GetClass struct {
	Context *middleware.Context
	Handler GetClassHandler
}

func (o *GetClass) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetClassParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *v1.Principal
	if uprinc != nil {
		principal = uprinc.(*v1.Principal) // this is really a v1.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
