// Code generated by go-swagger; DO NOT EDIT.

package notification_listeners_client_side

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/MxelA/tmf-service-go/pkg/swagger/tmf641v4_2/server/models"
)

// ListenToCancelServiceOrderInformationRequiredEventCreatedCode is the HTTP code returned for type ListenToCancelServiceOrderInformationRequiredEventCreated
const ListenToCancelServiceOrderInformationRequiredEventCreatedCode int = 201

/*
ListenToCancelServiceOrderInformationRequiredEventCreated Notified

swagger:response listenToCancelServiceOrderInformationRequiredEventCreated
*/
type ListenToCancelServiceOrderInformationRequiredEventCreated struct {

	/*
	  In: Body
	*/
	Payload *models.EventSubscription `json:"body,omitempty"`
}

// NewListenToCancelServiceOrderInformationRequiredEventCreated creates ListenToCancelServiceOrderInformationRequiredEventCreated with default headers values
func NewListenToCancelServiceOrderInformationRequiredEventCreated() *ListenToCancelServiceOrderInformationRequiredEventCreated {

	return &ListenToCancelServiceOrderInformationRequiredEventCreated{}
}

// WithPayload adds the payload to the listen to cancel service order information required event created response
func (o *ListenToCancelServiceOrderInformationRequiredEventCreated) WithPayload(payload *models.EventSubscription) *ListenToCancelServiceOrderInformationRequiredEventCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the listen to cancel service order information required event created response
func (o *ListenToCancelServiceOrderInformationRequiredEventCreated) SetPayload(payload *models.EventSubscription) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListenToCancelServiceOrderInformationRequiredEventCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListenToCancelServiceOrderInformationRequiredEventBadRequestCode is the HTTP code returned for type ListenToCancelServiceOrderInformationRequiredEventBadRequest
const ListenToCancelServiceOrderInformationRequiredEventBadRequestCode int = 400

/*
ListenToCancelServiceOrderInformationRequiredEventBadRequest Bad Request

swagger:response listenToCancelServiceOrderInformationRequiredEventBadRequest
*/
type ListenToCancelServiceOrderInformationRequiredEventBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListenToCancelServiceOrderInformationRequiredEventBadRequest creates ListenToCancelServiceOrderInformationRequiredEventBadRequest with default headers values
func NewListenToCancelServiceOrderInformationRequiredEventBadRequest() *ListenToCancelServiceOrderInformationRequiredEventBadRequest {

	return &ListenToCancelServiceOrderInformationRequiredEventBadRequest{}
}

// WithPayload adds the payload to the listen to cancel service order information required event bad request response
func (o *ListenToCancelServiceOrderInformationRequiredEventBadRequest) WithPayload(payload *models.Error) *ListenToCancelServiceOrderInformationRequiredEventBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the listen to cancel service order information required event bad request response
func (o *ListenToCancelServiceOrderInformationRequiredEventBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListenToCancelServiceOrderInformationRequiredEventBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListenToCancelServiceOrderInformationRequiredEventUnauthorizedCode is the HTTP code returned for type ListenToCancelServiceOrderInformationRequiredEventUnauthorized
const ListenToCancelServiceOrderInformationRequiredEventUnauthorizedCode int = 401

/*
ListenToCancelServiceOrderInformationRequiredEventUnauthorized Unauthorized

swagger:response listenToCancelServiceOrderInformationRequiredEventUnauthorized
*/
type ListenToCancelServiceOrderInformationRequiredEventUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListenToCancelServiceOrderInformationRequiredEventUnauthorized creates ListenToCancelServiceOrderInformationRequiredEventUnauthorized with default headers values
func NewListenToCancelServiceOrderInformationRequiredEventUnauthorized() *ListenToCancelServiceOrderInformationRequiredEventUnauthorized {

	return &ListenToCancelServiceOrderInformationRequiredEventUnauthorized{}
}

// WithPayload adds the payload to the listen to cancel service order information required event unauthorized response
func (o *ListenToCancelServiceOrderInformationRequiredEventUnauthorized) WithPayload(payload *models.Error) *ListenToCancelServiceOrderInformationRequiredEventUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the listen to cancel service order information required event unauthorized response
func (o *ListenToCancelServiceOrderInformationRequiredEventUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListenToCancelServiceOrderInformationRequiredEventUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListenToCancelServiceOrderInformationRequiredEventForbiddenCode is the HTTP code returned for type ListenToCancelServiceOrderInformationRequiredEventForbidden
const ListenToCancelServiceOrderInformationRequiredEventForbiddenCode int = 403

/*
ListenToCancelServiceOrderInformationRequiredEventForbidden Forbidden

swagger:response listenToCancelServiceOrderInformationRequiredEventForbidden
*/
type ListenToCancelServiceOrderInformationRequiredEventForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListenToCancelServiceOrderInformationRequiredEventForbidden creates ListenToCancelServiceOrderInformationRequiredEventForbidden with default headers values
func NewListenToCancelServiceOrderInformationRequiredEventForbidden() *ListenToCancelServiceOrderInformationRequiredEventForbidden {

	return &ListenToCancelServiceOrderInformationRequiredEventForbidden{}
}

// WithPayload adds the payload to the listen to cancel service order information required event forbidden response
func (o *ListenToCancelServiceOrderInformationRequiredEventForbidden) WithPayload(payload *models.Error) *ListenToCancelServiceOrderInformationRequiredEventForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the listen to cancel service order information required event forbidden response
func (o *ListenToCancelServiceOrderInformationRequiredEventForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListenToCancelServiceOrderInformationRequiredEventForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListenToCancelServiceOrderInformationRequiredEventNotFoundCode is the HTTP code returned for type ListenToCancelServiceOrderInformationRequiredEventNotFound
const ListenToCancelServiceOrderInformationRequiredEventNotFoundCode int = 404

/*
ListenToCancelServiceOrderInformationRequiredEventNotFound Not Found

swagger:response listenToCancelServiceOrderInformationRequiredEventNotFound
*/
type ListenToCancelServiceOrderInformationRequiredEventNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListenToCancelServiceOrderInformationRequiredEventNotFound creates ListenToCancelServiceOrderInformationRequiredEventNotFound with default headers values
func NewListenToCancelServiceOrderInformationRequiredEventNotFound() *ListenToCancelServiceOrderInformationRequiredEventNotFound {

	return &ListenToCancelServiceOrderInformationRequiredEventNotFound{}
}

// WithPayload adds the payload to the listen to cancel service order information required event not found response
func (o *ListenToCancelServiceOrderInformationRequiredEventNotFound) WithPayload(payload *models.Error) *ListenToCancelServiceOrderInformationRequiredEventNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the listen to cancel service order information required event not found response
func (o *ListenToCancelServiceOrderInformationRequiredEventNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListenToCancelServiceOrderInformationRequiredEventNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListenToCancelServiceOrderInformationRequiredEventMethodNotAllowedCode is the HTTP code returned for type ListenToCancelServiceOrderInformationRequiredEventMethodNotAllowed
const ListenToCancelServiceOrderInformationRequiredEventMethodNotAllowedCode int = 405

/*
ListenToCancelServiceOrderInformationRequiredEventMethodNotAllowed Method Not allowed

swagger:response listenToCancelServiceOrderInformationRequiredEventMethodNotAllowed
*/
type ListenToCancelServiceOrderInformationRequiredEventMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListenToCancelServiceOrderInformationRequiredEventMethodNotAllowed creates ListenToCancelServiceOrderInformationRequiredEventMethodNotAllowed with default headers values
func NewListenToCancelServiceOrderInformationRequiredEventMethodNotAllowed() *ListenToCancelServiceOrderInformationRequiredEventMethodNotAllowed {

	return &ListenToCancelServiceOrderInformationRequiredEventMethodNotAllowed{}
}

// WithPayload adds the payload to the listen to cancel service order information required event method not allowed response
func (o *ListenToCancelServiceOrderInformationRequiredEventMethodNotAllowed) WithPayload(payload *models.Error) *ListenToCancelServiceOrderInformationRequiredEventMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the listen to cancel service order information required event method not allowed response
func (o *ListenToCancelServiceOrderInformationRequiredEventMethodNotAllowed) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListenToCancelServiceOrderInformationRequiredEventMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListenToCancelServiceOrderInformationRequiredEventConflictCode is the HTTP code returned for type ListenToCancelServiceOrderInformationRequiredEventConflict
const ListenToCancelServiceOrderInformationRequiredEventConflictCode int = 409

/*
ListenToCancelServiceOrderInformationRequiredEventConflict Conflict

swagger:response listenToCancelServiceOrderInformationRequiredEventConflict
*/
type ListenToCancelServiceOrderInformationRequiredEventConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListenToCancelServiceOrderInformationRequiredEventConflict creates ListenToCancelServiceOrderInformationRequiredEventConflict with default headers values
func NewListenToCancelServiceOrderInformationRequiredEventConflict() *ListenToCancelServiceOrderInformationRequiredEventConflict {

	return &ListenToCancelServiceOrderInformationRequiredEventConflict{}
}

// WithPayload adds the payload to the listen to cancel service order information required event conflict response
func (o *ListenToCancelServiceOrderInformationRequiredEventConflict) WithPayload(payload *models.Error) *ListenToCancelServiceOrderInformationRequiredEventConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the listen to cancel service order information required event conflict response
func (o *ListenToCancelServiceOrderInformationRequiredEventConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListenToCancelServiceOrderInformationRequiredEventConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListenToCancelServiceOrderInformationRequiredEventInternalServerErrorCode is the HTTP code returned for type ListenToCancelServiceOrderInformationRequiredEventInternalServerError
const ListenToCancelServiceOrderInformationRequiredEventInternalServerErrorCode int = 500

/*
ListenToCancelServiceOrderInformationRequiredEventInternalServerError Internal Server Error

swagger:response listenToCancelServiceOrderInformationRequiredEventInternalServerError
*/
type ListenToCancelServiceOrderInformationRequiredEventInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListenToCancelServiceOrderInformationRequiredEventInternalServerError creates ListenToCancelServiceOrderInformationRequiredEventInternalServerError with default headers values
func NewListenToCancelServiceOrderInformationRequiredEventInternalServerError() *ListenToCancelServiceOrderInformationRequiredEventInternalServerError {

	return &ListenToCancelServiceOrderInformationRequiredEventInternalServerError{}
}

// WithPayload adds the payload to the listen to cancel service order information required event internal server error response
func (o *ListenToCancelServiceOrderInformationRequiredEventInternalServerError) WithPayload(payload *models.Error) *ListenToCancelServiceOrderInformationRequiredEventInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the listen to cancel service order information required event internal server error response
func (o *ListenToCancelServiceOrderInformationRequiredEventInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListenToCancelServiceOrderInformationRequiredEventInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
