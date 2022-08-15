// Code generated by go-swagger; DO NOT EDIT.

package monitor

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	v1 "github.com/yushk/healthy_backend/apiserver/v1"
)

// CreateMoveDataHandlerFunc turns a function with the right signature into a create move data handler
type CreateMoveDataHandlerFunc func(CreateMoveDataParams, *v1.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateMoveDataHandlerFunc) Handle(params CreateMoveDataParams, principal *v1.Principal) middleware.Responder {
	return fn(params, principal)
}

// CreateMoveDataHandler interface for that can handle valid create move data params
type CreateMoveDataHandler interface {
	Handle(CreateMoveDataParams, *v1.Principal) middleware.Responder
}

// NewCreateMoveData creates a new http.Handler for the create move data operation
func NewCreateMoveData(ctx *middleware.Context, handler CreateMoveDataHandler) *CreateMoveData {
	return &CreateMoveData{Context: ctx, Handler: handler}
}

/* CreateMoveData swagger:route POST /v1/moveData monitor createMoveData

创建运动数据信息

创建运动数据信息

*/
type CreateMoveData struct {
	Context *middleware.Context
	Handler CreateMoveDataHandler
}

func (o *CreateMoveData) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewCreateMoveDataParams()
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
