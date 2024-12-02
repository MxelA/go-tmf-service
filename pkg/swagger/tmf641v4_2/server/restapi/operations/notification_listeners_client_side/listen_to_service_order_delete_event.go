// Code generated by go-swagger; DO NOT EDIT.

package notification_listeners_client_side

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListenToServiceOrderDeleteEventHandlerFunc turns a function with the right signature into a listen to service order delete event handler
type ListenToServiceOrderDeleteEventHandlerFunc func(ListenToServiceOrderDeleteEventParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListenToServiceOrderDeleteEventHandlerFunc) Handle(params ListenToServiceOrderDeleteEventParams) middleware.Responder {
	return fn(params)
}

// ListenToServiceOrderDeleteEventHandler interface for that can handle valid listen to service order delete event params
type ListenToServiceOrderDeleteEventHandler interface {
	Handle(ListenToServiceOrderDeleteEventParams) middleware.Responder
}

// NewListenToServiceOrderDeleteEvent creates a new http.Handler for the listen to service order delete event operation
func NewListenToServiceOrderDeleteEvent(ctx *middleware.Context, handler ListenToServiceOrderDeleteEventHandler) *ListenToServiceOrderDeleteEvent {
	return &ListenToServiceOrderDeleteEvent{Context: ctx, Handler: handler}
}

/*
	ListenToServiceOrderDeleteEvent swagger:route POST /listener/serviceOrderDeleteEvent notification listeners (client side) listenToServiceOrderDeleteEvent

# Client listener for entity ServiceOrderDeleteEvent

Example of a client listener for receiving the notification ServiceOrderDeleteEvent
*/
type ListenToServiceOrderDeleteEvent struct {
	Context *middleware.Context
	Handler ListenToServiceOrderDeleteEventHandler
}

func (o *ListenToServiceOrderDeleteEvent) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewListenToServiceOrderDeleteEventParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}