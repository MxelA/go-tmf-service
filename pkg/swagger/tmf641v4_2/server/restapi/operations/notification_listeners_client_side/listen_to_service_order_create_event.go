// Code generated by go-swagger; DO NOT EDIT.

package notification_listeners_client_side

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListenToServiceOrderCreateEventHandlerFunc turns a function with the right signature into a listen to service order create event handler
type ListenToServiceOrderCreateEventHandlerFunc func(ListenToServiceOrderCreateEventParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListenToServiceOrderCreateEventHandlerFunc) Handle(params ListenToServiceOrderCreateEventParams) middleware.Responder {
	return fn(params)
}

// ListenToServiceOrderCreateEventHandler interface for that can handle valid listen to service order create event params
type ListenToServiceOrderCreateEventHandler interface {
	Handle(ListenToServiceOrderCreateEventParams) middleware.Responder
}

// NewListenToServiceOrderCreateEvent creates a new http.Handler for the listen to service order create event operation
func NewListenToServiceOrderCreateEvent(ctx *middleware.Context, handler ListenToServiceOrderCreateEventHandler) *ListenToServiceOrderCreateEvent {
	return &ListenToServiceOrderCreateEvent{Context: ctx, Handler: handler}
}

/*
	ListenToServiceOrderCreateEvent swagger:route POST /listener/serviceOrderCreateEvent notification listeners (client side) listenToServiceOrderCreateEvent

# Client listener for entity ServiceOrderCreateEvent

Example of a client listener for receiving the notification ServiceOrderCreateEvent
*/
type ListenToServiceOrderCreateEvent struct {
	Context *middleware.Context
	Handler ListenToServiceOrderCreateEventHandler
}

func (o *ListenToServiceOrderCreateEvent) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewListenToServiceOrderCreateEventParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
