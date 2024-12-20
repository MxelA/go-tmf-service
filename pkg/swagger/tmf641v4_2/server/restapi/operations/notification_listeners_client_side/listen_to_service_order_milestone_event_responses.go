// Code generated by go-swagger; DO NOT EDIT.

package notification_listeners_client_side

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/MxelA/tmf-service-go/pkg/swagger/tmf641v4_2/server/models"
)

// ListenToServiceOrderMilestoneEventCreatedCode is the HTTP code returned for type ListenToServiceOrderMilestoneEventCreated
const ListenToServiceOrderMilestoneEventCreatedCode int = 201

/*
ListenToServiceOrderMilestoneEventCreated Notified

swagger:response listenToServiceOrderMilestoneEventCreated
*/
type ListenToServiceOrderMilestoneEventCreated struct {

	/*
	  In: Body
	*/
	Payload *models.EventSubscription `json:"body,omitempty"`
}

// NewListenToServiceOrderMilestoneEventCreated creates ListenToServiceOrderMilestoneEventCreated with default headers values
func NewListenToServiceOrderMilestoneEventCreated() *ListenToServiceOrderMilestoneEventCreated {

	return &ListenToServiceOrderMilestoneEventCreated{}
}

// WithPayload adds the payload to the listen to service order milestone event created response
func (o *ListenToServiceOrderMilestoneEventCreated) WithPayload(payload *models.EventSubscription) *ListenToServiceOrderMilestoneEventCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the listen to service order milestone event created response
func (o *ListenToServiceOrderMilestoneEventCreated) SetPayload(payload *models.EventSubscription) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListenToServiceOrderMilestoneEventCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListenToServiceOrderMilestoneEventBadRequestCode is the HTTP code returned for type ListenToServiceOrderMilestoneEventBadRequest
const ListenToServiceOrderMilestoneEventBadRequestCode int = 400

/*
ListenToServiceOrderMilestoneEventBadRequest Bad Request

swagger:response listenToServiceOrderMilestoneEventBadRequest
*/
type ListenToServiceOrderMilestoneEventBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListenToServiceOrderMilestoneEventBadRequest creates ListenToServiceOrderMilestoneEventBadRequest with default headers values
func NewListenToServiceOrderMilestoneEventBadRequest() *ListenToServiceOrderMilestoneEventBadRequest {

	return &ListenToServiceOrderMilestoneEventBadRequest{}
}

// WithPayload adds the payload to the listen to service order milestone event bad request response
func (o *ListenToServiceOrderMilestoneEventBadRequest) WithPayload(payload *models.Error) *ListenToServiceOrderMilestoneEventBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the listen to service order milestone event bad request response
func (o *ListenToServiceOrderMilestoneEventBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListenToServiceOrderMilestoneEventBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListenToServiceOrderMilestoneEventUnauthorizedCode is the HTTP code returned for type ListenToServiceOrderMilestoneEventUnauthorized
const ListenToServiceOrderMilestoneEventUnauthorizedCode int = 401

/*
ListenToServiceOrderMilestoneEventUnauthorized Unauthorized

swagger:response listenToServiceOrderMilestoneEventUnauthorized
*/
type ListenToServiceOrderMilestoneEventUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListenToServiceOrderMilestoneEventUnauthorized creates ListenToServiceOrderMilestoneEventUnauthorized with default headers values
func NewListenToServiceOrderMilestoneEventUnauthorized() *ListenToServiceOrderMilestoneEventUnauthorized {

	return &ListenToServiceOrderMilestoneEventUnauthorized{}
}

// WithPayload adds the payload to the listen to service order milestone event unauthorized response
func (o *ListenToServiceOrderMilestoneEventUnauthorized) WithPayload(payload *models.Error) *ListenToServiceOrderMilestoneEventUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the listen to service order milestone event unauthorized response
func (o *ListenToServiceOrderMilestoneEventUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListenToServiceOrderMilestoneEventUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListenToServiceOrderMilestoneEventForbiddenCode is the HTTP code returned for type ListenToServiceOrderMilestoneEventForbidden
const ListenToServiceOrderMilestoneEventForbiddenCode int = 403

/*
ListenToServiceOrderMilestoneEventForbidden Forbidden

swagger:response listenToServiceOrderMilestoneEventForbidden
*/
type ListenToServiceOrderMilestoneEventForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListenToServiceOrderMilestoneEventForbidden creates ListenToServiceOrderMilestoneEventForbidden with default headers values
func NewListenToServiceOrderMilestoneEventForbidden() *ListenToServiceOrderMilestoneEventForbidden {

	return &ListenToServiceOrderMilestoneEventForbidden{}
}

// WithPayload adds the payload to the listen to service order milestone event forbidden response
func (o *ListenToServiceOrderMilestoneEventForbidden) WithPayload(payload *models.Error) *ListenToServiceOrderMilestoneEventForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the listen to service order milestone event forbidden response
func (o *ListenToServiceOrderMilestoneEventForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListenToServiceOrderMilestoneEventForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListenToServiceOrderMilestoneEventNotFoundCode is the HTTP code returned for type ListenToServiceOrderMilestoneEventNotFound
const ListenToServiceOrderMilestoneEventNotFoundCode int = 404

/*
ListenToServiceOrderMilestoneEventNotFound Not Found

swagger:response listenToServiceOrderMilestoneEventNotFound
*/
type ListenToServiceOrderMilestoneEventNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListenToServiceOrderMilestoneEventNotFound creates ListenToServiceOrderMilestoneEventNotFound with default headers values
func NewListenToServiceOrderMilestoneEventNotFound() *ListenToServiceOrderMilestoneEventNotFound {

	return &ListenToServiceOrderMilestoneEventNotFound{}
}

// WithPayload adds the payload to the listen to service order milestone event not found response
func (o *ListenToServiceOrderMilestoneEventNotFound) WithPayload(payload *models.Error) *ListenToServiceOrderMilestoneEventNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the listen to service order milestone event not found response
func (o *ListenToServiceOrderMilestoneEventNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListenToServiceOrderMilestoneEventNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListenToServiceOrderMilestoneEventMethodNotAllowedCode is the HTTP code returned for type ListenToServiceOrderMilestoneEventMethodNotAllowed
const ListenToServiceOrderMilestoneEventMethodNotAllowedCode int = 405

/*
ListenToServiceOrderMilestoneEventMethodNotAllowed Method Not allowed

swagger:response listenToServiceOrderMilestoneEventMethodNotAllowed
*/
type ListenToServiceOrderMilestoneEventMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListenToServiceOrderMilestoneEventMethodNotAllowed creates ListenToServiceOrderMilestoneEventMethodNotAllowed with default headers values
func NewListenToServiceOrderMilestoneEventMethodNotAllowed() *ListenToServiceOrderMilestoneEventMethodNotAllowed {

	return &ListenToServiceOrderMilestoneEventMethodNotAllowed{}
}

// WithPayload adds the payload to the listen to service order milestone event method not allowed response
func (o *ListenToServiceOrderMilestoneEventMethodNotAllowed) WithPayload(payload *models.Error) *ListenToServiceOrderMilestoneEventMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the listen to service order milestone event method not allowed response
func (o *ListenToServiceOrderMilestoneEventMethodNotAllowed) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListenToServiceOrderMilestoneEventMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListenToServiceOrderMilestoneEventConflictCode is the HTTP code returned for type ListenToServiceOrderMilestoneEventConflict
const ListenToServiceOrderMilestoneEventConflictCode int = 409

/*
ListenToServiceOrderMilestoneEventConflict Conflict

swagger:response listenToServiceOrderMilestoneEventConflict
*/
type ListenToServiceOrderMilestoneEventConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListenToServiceOrderMilestoneEventConflict creates ListenToServiceOrderMilestoneEventConflict with default headers values
func NewListenToServiceOrderMilestoneEventConflict() *ListenToServiceOrderMilestoneEventConflict {

	return &ListenToServiceOrderMilestoneEventConflict{}
}

// WithPayload adds the payload to the listen to service order milestone event conflict response
func (o *ListenToServiceOrderMilestoneEventConflict) WithPayload(payload *models.Error) *ListenToServiceOrderMilestoneEventConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the listen to service order milestone event conflict response
func (o *ListenToServiceOrderMilestoneEventConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListenToServiceOrderMilestoneEventConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListenToServiceOrderMilestoneEventInternalServerErrorCode is the HTTP code returned for type ListenToServiceOrderMilestoneEventInternalServerError
const ListenToServiceOrderMilestoneEventInternalServerErrorCode int = 500

/*
ListenToServiceOrderMilestoneEventInternalServerError Internal Server Error

swagger:response listenToServiceOrderMilestoneEventInternalServerError
*/
type ListenToServiceOrderMilestoneEventInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListenToServiceOrderMilestoneEventInternalServerError creates ListenToServiceOrderMilestoneEventInternalServerError with default headers values
func NewListenToServiceOrderMilestoneEventInternalServerError() *ListenToServiceOrderMilestoneEventInternalServerError {

	return &ListenToServiceOrderMilestoneEventInternalServerError{}
}

// WithPayload adds the payload to the listen to service order milestone event internal server error response
func (o *ListenToServiceOrderMilestoneEventInternalServerError) WithPayload(payload *models.Error) *ListenToServiceOrderMilestoneEventInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the listen to service order milestone event internal server error response
func (o *ListenToServiceOrderMilestoneEventInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListenToServiceOrderMilestoneEventInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
