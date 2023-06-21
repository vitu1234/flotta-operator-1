// Code generated by go-swagger; DO NOT EDIT.

package yggdrasil

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PostControlMessageForDeviceHandlerFunc turns a function with the right signature into a post control message for device handler
type PostControlMessageForDeviceHandlerFunc func(PostControlMessageForDeviceParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostControlMessageForDeviceHandlerFunc) Handle(params PostControlMessageForDeviceParams) middleware.Responder {
	return fn(params)
}

// PostControlMessageForDeviceHandler interface for that can handle valid post control message for device params
type PostControlMessageForDeviceHandler interface {
	Handle(PostControlMessageForDeviceParams) middleware.Responder
}

// NewPostControlMessageForDevice creates a new http.Handler for the post control message for device operation
func NewPostControlMessageForDevice(ctx *middleware.Context, handler PostControlMessageForDeviceHandler) *PostControlMessageForDevice {
	return &PostControlMessageForDevice{Context: ctx, Handler: handler}
}

/* PostControlMessageForDevice swagger:route POST /control/{device_id}/out yggdrasil postControlMessageForDevice

Post control message for device API

*/
type PostControlMessageForDevice struct {
	Context *middleware.Context
	Handler PostControlMessageForDeviceHandler
}

func (o *PostControlMessageForDevice) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostControlMessageForDeviceParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
