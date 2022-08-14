// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	v1 "gitee.com/healthy/backend/apiserver/v1"
	"github.com/go-openapi/runtime/middleware"
)

// GetUserHandlerFunc turns a function with the right signature into a get user handler
type GetUserHandlerFunc func(GetUserParams, *v1.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GetUserHandlerFunc) Handle(params GetUserParams, principal *v1.Principal) middleware.Responder {
	return fn(params, principal)
}

// GetUserHandler interface for that can handle valid get user params
type GetUserHandler interface {
	Handle(GetUserParams, *v1.Principal) middleware.Responder
}

// NewGetUser creates a new http.Handler for the get user operation
func NewGetUser(ctx *middleware.Context, handler GetUserHandler) *GetUser {
	return &GetUser{Context: ctx, Handler: handler}
}

/* GetUser swagger:route GET /v1/user/{id} user getUser

获取用户信息

获取单个用户的详细信息

*/
type GetUser struct {
	Context *middleware.Context
	Handler GetUserHandler
}

func (o *GetUser) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetUserParams()
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
