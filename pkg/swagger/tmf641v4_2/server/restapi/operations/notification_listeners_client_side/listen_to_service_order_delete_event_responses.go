// Code generated by go-swagger; DO NOT EDIT.

package notification_listeners_client_side

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/MxelA/tmf-service-go/pkg/swagger/tmf641v4_2/server/models"
)

// ListenToServiceOrderDeleteEventCreatedCode is the HTTP code returned for type ListenToServiceOrderDeleteEventCreated
const ListenToServiceOrderDeleteEventCreatedCode int = 201

/*
ListenToServiceOrderDeleteEventCreated Notified

swagger:response listenToServiceOrderDeleteEventCreated
*/
type ListenToServiceOrderDeleteEventCreated struct {

	/*
	  In: Body
	*/
	Payload *models.EventSubscription `json:"body,omitempty"`
}

// NewListenToServiceOrderDeleteEventCreated creates ListenToServiceOrderDeleteEventCreated with default headers values
func NewListenToServiceOrderDeleteEventCreated() *ListenToServiceOrderDeleteEventCreated {

	return &ListenToServiceOrderDeleteEventCreated{}
}

// WithPayload adds the payload to the listen to service order delete event created response
func (o *ListenToServiceOrderDeleteEventCreated) WithPayload(payload *models.EventSubscription) *ListenToServiceOrderDeleteEventCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the listen to service order delete event created response
func (o *ListenToServiceOrderDeleteEventCreated) SetPayload(payload *models.EventSubscription) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListenToServiceOrderDeleteEventCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListenToServiceOrderDeleteEventBadRequestCode is the HTTP code returned for type ListenToServiceOrderDeleteEventBadRequest
const ListenToServiceOrderDeleteEventBadRequestCode int = 400

/*
ListenToServiceOrderDeleteEventBadRequest Bad Request

swagger:response listenToServiceOrderDeleteEventBadRequest
*/
type ListenToServiceOrderDeleteEventBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListenToServiceOrderDeleteEventBadRequest creates ListenToServiceOrderDeleteEventBadRequest with default headers values
func NewListenToServiceOrderDeleteEventBadRequest() *ListenToServiceOrderDeleteEventBadRequest {

	return &ListenToServiceOrderDeleteEventBadRequest{}
}

// WithPayload adds the payload to the listen to service order delete event bad request response
func (o *ListenToServiceOrderDeleteEventBadRequest) WithPayload(payload *models.Error) *ListenToServiceOrderDeleteEventBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the listen to service order delete event bad request response
func (o *ListenToServiceOrderDeleteEventBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListenToServiceOrderDeleteEventBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListenToServiceOrderDeleteEventUnauthorizedCode is the HTTP code returned for type ListenToServiceOrderDeleteEventUnauthorized
const ListenToServiceOrderDeleteEventUnauthorizedCode int = 401

/*
ListenToServiceOrderDeleteEventUnauthorized Unauthorized

swagger:response listenToServiceOrderDeleteEventUnauthorized
*/
type ListenToServiceOrderDeleteEventUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListenToServiceOrderDeleteEventUnauthorized creates ListenToServiceOrderDeleteEventUnauthorized with default headers values
func NewListenToServiceOrderDeleteEventUnauthorized() *ListenToServiceOrderDeleteEventUnauthorized {

	return &ListenToServiceOrderDeleteEventUnauthorized{}
}

// WithPayload adds the payload to the listen to service order delete event unauthorized response
func (o *ListenToServiceOrderDeleteEventUnauthorized) WithPayload(payload *models.Error) *ListenToServiceOrderDeleteEventUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the listen to service order delete event unauthorized response
func (o *ListenToServiceOrderDeleteEventUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListenToServiceOrderDeleteEventUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListenToServiceOrderDeleteEventForbiddenCode is the HTTP code returned for type ListenToServiceOrderDeleteEventForbidden
const ListenToServiceOrderDeleteEventForbiddenCode int = 403

/*
ListenToServiceOrderDeleteEventForbidden Forbidden

swagger:response listenToServiceOrderDeleteEventForbidden
*/
type ListenToServiceOrderDeleteEventForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListenToServiceOrderDeleteEventForbidden creates ListenToServiceOrderDeleteEventForbidden with default headers values
func NewListenToServiceOrderDeleteEventForbidden() *ListenToServiceOrderDeleteEventForbidden {

	return &ListenToServiceOrderDeleteEventForbidden{}
}

// WithPayload adds the payload to the listen to service order delete event forbidden response
func (o *ListenToServiceOrderDeleteEventForbidden) WithPayload(payload *models.Error) *ListenToServiceOrderDeleteEventForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the listen to service order delete event forbidden response
func (o *ListenToServiceOrderDeleteEventForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListenToServiceOrderDeleteEventForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListenToServiceOrderDeleteEventNotFoundCode is the HTTP code returned for type ListenToServiceOrderDeleteEventNotFound
const ListenToServiceOrderDeleteEventNotFoundCode int = 404

/*
ListenToServiceOrderDeleteEventNotFound Not Found

swagger:response listenToServiceOrderDeleteEventNotFound
*/
type ListenToServiceOrderDeleteEventNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListenToServiceOrderDeleteEventNotFound creates ListenToServiceOrderDeleteEventNotFound with default headers values
func NewListenToServiceOrderDeleteEventNotFound() *ListenToServiceOrderDeleteEventNotFound {

	return &ListenToServiceOrderDeleteEventNotFound{}
}

// WithPayload adds the payload to the listen to service order delete event not found response
func (o *ListenToServiceOrderDeleteEventNotFound) WithPayload(payload *models.Error) *ListenToServiceOrderDeleteEventNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the listen to service order delete event not found response
func (o *ListenToServiceOrderDeleteEventNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListenToServiceOrderDeleteEventNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListenToServiceOrderDeleteEventMethodNotAllowedCode is the HTTP code returned for type ListenToServiceOrderDeleteEventMethodNotAllowed
const ListenToServiceOrderDeleteEventMethodNotAllowedCode int = 405

/*
ListenToServiceOrderDeleteEventMethodNotAllowed Method Not allowed

swagger:response listenToServiceOrderDeleteEventMethodNotAllowed
*/
type ListenToServiceOrderDeleteEventMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListenToServiceOrderDeleteEventMethodNotAllowed creates ListenToServiceOrderDeleteEventMethodNotAllowed with default headers values
func NewListenToServiceOrderDeleteEventMethodNotAllowed() *ListenToServiceOrderDeleteEventMethodNotAllowed {

	return &ListenToServiceOrderDeleteEventMethodNotAllowed{}
}

// WithPayload adds the payload to the listen to service order delete event method not allowed response
func (o *ListenToServiceOrderDeleteEventMethodNotAllowed) WithPayload(payload *models.Error) *ListenToServiceOrderDeleteEventMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the listen to service order delete event method not allowed response
func (o *ListenToServiceOrderDeleteEventMethodNotAllowed) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListenToServiceOrderDeleteEventMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListenToServiceOrderDeleteEventConflictCode is the HTTP code returned for type ListenToServiceOrderDeleteEventConflict
const ListenToServiceOrderDeleteEventConflictCode int = 409

/*
ListenToServiceOrderDeleteEventConflict Conflict

swagger:response listenToServiceOrderDeleteEventConflict
*/
type ListenToServiceOrderDeleteEventConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListenToServiceOrderDeleteEventConflict creates ListenToServiceOrderDeleteEventConflict with default headers values
func NewListenToServiceOrderDeleteEventConflict() *ListenToServiceOrderDeleteEventConflict {

	return &ListenToServiceOrderDeleteEventConflict{}
}

// WithPayload adds the payload to the listen to service order delete event conflict response
func (o *ListenToServiceOrderDeleteEventConflict) WithPayload(payload *models.Error) *ListenToServiceOrderDeleteEventConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the listen to service order delete event conflict response
func (o *ListenToServiceOrderDeleteEventConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListenToServiceOrderDeleteEventConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListenToServiceOrderDeleteEventInternalServerErrorCode is the HTTP code returned for type ListenToServiceOrderDeleteEventInternalServerError
const ListenToServiceOrderDeleteEventInternalServerErrorCode int = 500

/*
ListenToServiceOrderDeleteEventInternalServerError Internal Server Error

swagger:response listenToServiceOrderDeleteEventInternalServerError
*/
type ListenToServiceOrderDeleteEventInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListenToServiceOrderDeleteEventInternalServerError creates ListenToServiceOrderDeleteEventInternalServerError with default headers values
func NewListenToServiceOrderDeleteEventInternalServerError() *ListenToServiceOrderDeleteEventInternalServerError {

	return &ListenToServiceOrderDeleteEventInternalServerError{}
}

// WithPayload adds the payload to the listen to service order delete event internal server error response
func (o *ListenToServiceOrderDeleteEventInternalServerError) WithPayload(payload *models.Error) *ListenToServiceOrderDeleteEventInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the listen to service order delete event internal server error response
func (o *ListenToServiceOrderDeleteEventInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListenToServiceOrderDeleteEventInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}