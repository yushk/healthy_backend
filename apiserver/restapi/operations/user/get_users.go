// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	v1 "github.com/yushk/healthy_backend/apiserver/v1"
)

// GetUsersHandlerFunc turns a function with the right signature into a get users handler
type GetUsersHandlerFunc func(GetUsersParams, *v1.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GetUsersHandlerFunc) Handle(params GetUsersParams, principal *v1.Principal) middleware.Responder {
	return fn(params, principal)
}

// GetUsersHandler interface for that can handle valid get users params
type GetUsersHandler interface {
	Handle(GetUsersParams, *v1.Principal) middleware.Responder
}

// NewGetUsers creates a new http.Handler for the get users operation
func NewGetUsers(ctx *middleware.Context, handler GetUsersHandler) *GetUsers {
	return &GetUsers{Context: ctx, Handler: handler}
}

/* GetUsers swagger:route GET /v1/users user getUsers

获取用户列表

获取用户列表

*/
type GetUsers struct {
	Context *middleware.Context
	Handler GetUsersHandler
}

func (o *GetUsers) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetUsersParams()
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
