// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	v1 "gitee.com/healthy/backend/apiserver/v1"
	"github.com/go-openapi/runtime/middleware"
)

// GetPersonalHandlerFunc turns a function with the right signature into a get personal handler
type GetPersonalHandlerFunc func(GetPersonalParams, *v1.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GetPersonalHandlerFunc) Handle(params GetPersonalParams, principal *v1.Principal) middleware.Responder {
	return fn(params, principal)
}

// GetPersonalHandler interface for that can handle valid get personal params
type GetPersonalHandler interface {
	Handle(GetPersonalParams, *v1.Principal) middleware.Responder
}

// NewGetPersonal creates a new http.Handler for the get personal operation
func NewGetPersonal(ctx *middleware.Context, handler GetPersonalHandler) *GetPersonal {
	return &GetPersonal{Context: ctx, Handler: handler}
}

/* GetPersonal swagger:route GET /v1/personal/{id} user getPersonal

获取个人基础信息

获取个人基础信息

*/
type GetPersonal struct {
	Context *middleware.Context
	Handler GetPersonalHandler
}

func (o *GetPersonal) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetPersonalParams()
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
