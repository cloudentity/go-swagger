// Code generated by go-swagger; DO NOT EDIT.

package customers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/cloudentity/go-swagger/examples/authentication/models"
)

// CreateHandlerFunc turns a function with the right signature into a create handler
type CreateHandlerFunc func(CreateParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateHandlerFunc) Handle(params CreateParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// CreateHandler interface for that can handle valid create params
type CreateHandler interface {
	Handle(CreateParams, *models.Principal) middleware.Responder
}

// NewCreate creates a new http.Handler for the create operation
func NewCreate(ctx *middleware.Context, handler CreateHandler) *Create {
	return &Create{Context: ctx, Handler: handler}
}

/*
	Create swagger:route POST /customers customers create

Create a new customer to track
*/
type Create struct {
	Context *middleware.Context
	Handler CreateHandler
}

func (o *Create) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewCreateParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
