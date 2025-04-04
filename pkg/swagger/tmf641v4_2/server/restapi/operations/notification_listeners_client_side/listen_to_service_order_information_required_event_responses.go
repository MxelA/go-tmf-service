// Code generated by go-swagger; DO NOT EDIT.

package notification_listeners_client_side

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/MxelA/tmf-service-go/pkg/swagger/tmf641v4_2/server/models"
)

// Alex Server response
// ListenToServiceOrderInformationRequiredEventCreatedCode is the HTTP code returned for type ListenToServiceOrderInformationRequiredEventCreated
const ListenToServiceOrderInformationRequiredEventCreatedCode int = 201

/*
ListenToServiceOrderInformationRequiredEventCreated Notified

swagger:response listenToServiceOrderInformationRequiredEventCreated
*/
type ListenToServiceOrderInformationRequiredEventCreated struct {

	/*
	  In: Body
	*/
	Payload *models.EventSubscription `json:"body,omitempty"`
}

type ListenToServiceOrderInformationRequiredEventCreatedRaw struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewListenToServiceOrderInformationRequiredEventCreated creates ListenToServiceOrderInformationRequiredEventCreated with default headers values
func NewListenToServiceOrderInformationRequiredEventCreated() *ListenToServiceOrderInformationRequiredEventCreated {

	return &ListenToServiceOrderInformationRequiredEventCreated{}
}

// NewListenToServiceOrderInformationRequiredEventCreated creates ListenToServiceOrderInformationRequiredEventCreatedRaw with default headers values
func NewListenToServiceOrderInformationRequiredEventCreatedRaw() *ListenToServiceOrderInformationRequiredEventCreatedRaw {

	return &ListenToServiceOrderInformationRequiredEventCreatedRaw{}
}

// WithPayload adds the payload to the listen to service order information required event created response
func (o *ListenToServiceOrderInformationRequiredEventCreated) WithPayload(payload *models.EventSubscription) *ListenToServiceOrderInformationRequiredEventCreated {
	o.Payload = payload
	return o
}

// WithPayload adds the payload to the listen to service order information required event created response
func (o *ListenToServiceOrderInformationRequiredEventCreatedRaw) WithPayload(payload interface{}) *ListenToServiceOrderInformationRequiredEventCreatedRaw {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the listen to service order information required event created response
func (o *ListenToServiceOrderInformationRequiredEventCreated) SetPayload(payload *models.EventSubscription) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListenToServiceOrderInformationRequiredEventCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WriteResponseRaw to the client
func (o *ListenToServiceOrderInformationRequiredEventCreatedRaw) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// Alex Server response
// ListenToServiceOrderInformationRequiredEventBadRequestCode is the HTTP code returned for type ListenToServiceOrderInformationRequiredEventBadRequest
const ListenToServiceOrderInformationRequiredEventBadRequestCode int = 400

/*
ListenToServiceOrderInformationRequiredEventBadRequest Bad Request

swagger:response listenToServiceOrderInformationRequiredEventBadRequest
*/
type ListenToServiceOrderInformationRequiredEventBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

type ListenToServiceOrderInformationRequiredEventBadRequestRaw struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewListenToServiceOrderInformationRequiredEventBadRequest creates ListenToServiceOrderInformationRequiredEventBadRequest with default headers values
func NewListenToServiceOrderInformationRequiredEventBadRequest() *ListenToServiceOrderInformationRequiredEventBadRequest {

	return &ListenToServiceOrderInformationRequiredEventBadRequest{}
}

// NewListenToServiceOrderInformationRequiredEventBadRequest creates ListenToServiceOrderInformationRequiredEventBadRequestRaw with default headers values
func NewListenToServiceOrderInformationRequiredEventBadRequestRaw() *ListenToServiceOrderInformationRequiredEventBadRequestRaw {

	return &ListenToServiceOrderInformationRequiredEventBadRequestRaw{}
}

// WithPayload adds the payload to the listen to service order information required event bad request response
func (o *ListenToServiceOrderInformationRequiredEventBadRequest) WithPayload(payload *models.Error) *ListenToServiceOrderInformationRequiredEventBadRequest {
	o.Payload = payload
	return o
}

// WithPayload adds the payload to the listen to service order information required event bad request response
func (o *ListenToServiceOrderInformationRequiredEventBadRequestRaw) WithPayload(payload interface{}) *ListenToServiceOrderInformationRequiredEventBadRequestRaw {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the listen to service order information required event bad request response
func (o *ListenToServiceOrderInformationRequiredEventBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListenToServiceOrderInformationRequiredEventBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WriteResponseRaw to the client
func (o *ListenToServiceOrderInformationRequiredEventBadRequestRaw) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// Alex Server response
// ListenToServiceOrderInformationRequiredEventUnauthorizedCode is the HTTP code returned for type ListenToServiceOrderInformationRequiredEventUnauthorized
const ListenToServiceOrderInformationRequiredEventUnauthorizedCode int = 401

/*
ListenToServiceOrderInformationRequiredEventUnauthorized Unauthorized

swagger:response listenToServiceOrderInformationRequiredEventUnauthorized
*/
type ListenToServiceOrderInformationRequiredEventUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

type ListenToServiceOrderInformationRequiredEventUnauthorizedRaw struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewListenToServiceOrderInformationRequiredEventUnauthorized creates ListenToServiceOrderInformationRequiredEventUnauthorized with default headers values
func NewListenToServiceOrderInformationRequiredEventUnauthorized() *ListenToServiceOrderInformationRequiredEventUnauthorized {

	return &ListenToServiceOrderInformationRequiredEventUnauthorized{}
}

// NewListenToServiceOrderInformationRequiredEventUnauthorized creates ListenToServiceOrderInformationRequiredEventUnauthorizedRaw with default headers values
func NewListenToServiceOrderInformationRequiredEventUnauthorizedRaw() *ListenToServiceOrderInformationRequiredEventUnauthorizedRaw {

	return &ListenToServiceOrderInformationRequiredEventUnauthorizedRaw{}
}

// WithPayload adds the payload to the listen to service order information required event unauthorized response
func (o *ListenToServiceOrderInformationRequiredEventUnauthorized) WithPayload(payload *models.Error) *ListenToServiceOrderInformationRequiredEventUnauthorized {
	o.Payload = payload
	return o
}

// WithPayload adds the payload to the listen to service order information required event unauthorized response
func (o *ListenToServiceOrderInformationRequiredEventUnauthorizedRaw) WithPayload(payload interface{}) *ListenToServiceOrderInformationRequiredEventUnauthorizedRaw {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the listen to service order information required event unauthorized response
func (o *ListenToServiceOrderInformationRequiredEventUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListenToServiceOrderInformationRequiredEventUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WriteResponseRaw to the client
func (o *ListenToServiceOrderInformationRequiredEventUnauthorizedRaw) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// Alex Server response
// ListenToServiceOrderInformationRequiredEventForbiddenCode is the HTTP code returned for type ListenToServiceOrderInformationRequiredEventForbidden
const ListenToServiceOrderInformationRequiredEventForbiddenCode int = 403

/*
ListenToServiceOrderInformationRequiredEventForbidden Forbidden

swagger:response listenToServiceOrderInformationRequiredEventForbidden
*/
type ListenToServiceOrderInformationRequiredEventForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

type ListenToServiceOrderInformationRequiredEventForbiddenRaw struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewListenToServiceOrderInformationRequiredEventForbidden creates ListenToServiceOrderInformationRequiredEventForbidden with default headers values
func NewListenToServiceOrderInformationRequiredEventForbidden() *ListenToServiceOrderInformationRequiredEventForbidden {

	return &ListenToServiceOrderInformationRequiredEventForbidden{}
}

// NewListenToServiceOrderInformationRequiredEventForbidden creates ListenToServiceOrderInformationRequiredEventForbiddenRaw with default headers values
func NewListenToServiceOrderInformationRequiredEventForbiddenRaw() *ListenToServiceOrderInformationRequiredEventForbiddenRaw {

	return &ListenToServiceOrderInformationRequiredEventForbiddenRaw{}
}

// WithPayload adds the payload to the listen to service order information required event forbidden response
func (o *ListenToServiceOrderInformationRequiredEventForbidden) WithPayload(payload *models.Error) *ListenToServiceOrderInformationRequiredEventForbidden {
	o.Payload = payload
	return o
}

// WithPayload adds the payload to the listen to service order information required event forbidden response
func (o *ListenToServiceOrderInformationRequiredEventForbiddenRaw) WithPayload(payload interface{}) *ListenToServiceOrderInformationRequiredEventForbiddenRaw {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the listen to service order information required event forbidden response
func (o *ListenToServiceOrderInformationRequiredEventForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListenToServiceOrderInformationRequiredEventForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WriteResponseRaw to the client
func (o *ListenToServiceOrderInformationRequiredEventForbiddenRaw) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// Alex Server response
// ListenToServiceOrderInformationRequiredEventNotFoundCode is the HTTP code returned for type ListenToServiceOrderInformationRequiredEventNotFound
const ListenToServiceOrderInformationRequiredEventNotFoundCode int = 404

/*
ListenToServiceOrderInformationRequiredEventNotFound Not Found

swagger:response listenToServiceOrderInformationRequiredEventNotFound
*/
type ListenToServiceOrderInformationRequiredEventNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

type ListenToServiceOrderInformationRequiredEventNotFoundRaw struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewListenToServiceOrderInformationRequiredEventNotFound creates ListenToServiceOrderInformationRequiredEventNotFound with default headers values
func NewListenToServiceOrderInformationRequiredEventNotFound() *ListenToServiceOrderInformationRequiredEventNotFound {

	return &ListenToServiceOrderInformationRequiredEventNotFound{}
}

// NewListenToServiceOrderInformationRequiredEventNotFound creates ListenToServiceOrderInformationRequiredEventNotFoundRaw with default headers values
func NewListenToServiceOrderInformationRequiredEventNotFoundRaw() *ListenToServiceOrderInformationRequiredEventNotFoundRaw {

	return &ListenToServiceOrderInformationRequiredEventNotFoundRaw{}
}

// WithPayload adds the payload to the listen to service order information required event not found response
func (o *ListenToServiceOrderInformationRequiredEventNotFound) WithPayload(payload *models.Error) *ListenToServiceOrderInformationRequiredEventNotFound {
	o.Payload = payload
	return o
}

// WithPayload adds the payload to the listen to service order information required event not found response
func (o *ListenToServiceOrderInformationRequiredEventNotFoundRaw) WithPayload(payload interface{}) *ListenToServiceOrderInformationRequiredEventNotFoundRaw {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the listen to service order information required event not found response
func (o *ListenToServiceOrderInformationRequiredEventNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListenToServiceOrderInformationRequiredEventNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WriteResponseRaw to the client
func (o *ListenToServiceOrderInformationRequiredEventNotFoundRaw) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// Alex Server response
// ListenToServiceOrderInformationRequiredEventMethodNotAllowedCode is the HTTP code returned for type ListenToServiceOrderInformationRequiredEventMethodNotAllowed
const ListenToServiceOrderInformationRequiredEventMethodNotAllowedCode int = 405

/*
ListenToServiceOrderInformationRequiredEventMethodNotAllowed Method Not allowed

swagger:response listenToServiceOrderInformationRequiredEventMethodNotAllowed
*/
type ListenToServiceOrderInformationRequiredEventMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

type ListenToServiceOrderInformationRequiredEventMethodNotAllowedRaw struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewListenToServiceOrderInformationRequiredEventMethodNotAllowed creates ListenToServiceOrderInformationRequiredEventMethodNotAllowed with default headers values
func NewListenToServiceOrderInformationRequiredEventMethodNotAllowed() *ListenToServiceOrderInformationRequiredEventMethodNotAllowed {

	return &ListenToServiceOrderInformationRequiredEventMethodNotAllowed{}
}

// NewListenToServiceOrderInformationRequiredEventMethodNotAllowed creates ListenToServiceOrderInformationRequiredEventMethodNotAllowedRaw with default headers values
func NewListenToServiceOrderInformationRequiredEventMethodNotAllowedRaw() *ListenToServiceOrderInformationRequiredEventMethodNotAllowedRaw {

	return &ListenToServiceOrderInformationRequiredEventMethodNotAllowedRaw{}
}

// WithPayload adds the payload to the listen to service order information required event method not allowed response
func (o *ListenToServiceOrderInformationRequiredEventMethodNotAllowed) WithPayload(payload *models.Error) *ListenToServiceOrderInformationRequiredEventMethodNotAllowed {
	o.Payload = payload
	return o
}

// WithPayload adds the payload to the listen to service order information required event method not allowed response
func (o *ListenToServiceOrderInformationRequiredEventMethodNotAllowedRaw) WithPayload(payload interface{}) *ListenToServiceOrderInformationRequiredEventMethodNotAllowedRaw {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the listen to service order information required event method not allowed response
func (o *ListenToServiceOrderInformationRequiredEventMethodNotAllowed) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListenToServiceOrderInformationRequiredEventMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WriteResponseRaw to the client
func (o *ListenToServiceOrderInformationRequiredEventMethodNotAllowedRaw) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// Alex Server response
// ListenToServiceOrderInformationRequiredEventConflictCode is the HTTP code returned for type ListenToServiceOrderInformationRequiredEventConflict
const ListenToServiceOrderInformationRequiredEventConflictCode int = 409

/*
ListenToServiceOrderInformationRequiredEventConflict Conflict

swagger:response listenToServiceOrderInformationRequiredEventConflict
*/
type ListenToServiceOrderInformationRequiredEventConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

type ListenToServiceOrderInformationRequiredEventConflictRaw struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewListenToServiceOrderInformationRequiredEventConflict creates ListenToServiceOrderInformationRequiredEventConflict with default headers values
func NewListenToServiceOrderInformationRequiredEventConflict() *ListenToServiceOrderInformationRequiredEventConflict {

	return &ListenToServiceOrderInformationRequiredEventConflict{}
}

// NewListenToServiceOrderInformationRequiredEventConflict creates ListenToServiceOrderInformationRequiredEventConflictRaw with default headers values
func NewListenToServiceOrderInformationRequiredEventConflictRaw() *ListenToServiceOrderInformationRequiredEventConflictRaw {

	return &ListenToServiceOrderInformationRequiredEventConflictRaw{}
}

// WithPayload adds the payload to the listen to service order information required event conflict response
func (o *ListenToServiceOrderInformationRequiredEventConflict) WithPayload(payload *models.Error) *ListenToServiceOrderInformationRequiredEventConflict {
	o.Payload = payload
	return o
}

// WithPayload adds the payload to the listen to service order information required event conflict response
func (o *ListenToServiceOrderInformationRequiredEventConflictRaw) WithPayload(payload interface{}) *ListenToServiceOrderInformationRequiredEventConflictRaw {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the listen to service order information required event conflict response
func (o *ListenToServiceOrderInformationRequiredEventConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListenToServiceOrderInformationRequiredEventConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WriteResponseRaw to the client
func (o *ListenToServiceOrderInformationRequiredEventConflictRaw) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// Alex Server response
// ListenToServiceOrderInformationRequiredEventInternalServerErrorCode is the HTTP code returned for type ListenToServiceOrderInformationRequiredEventInternalServerError
const ListenToServiceOrderInformationRequiredEventInternalServerErrorCode int = 500

/*
ListenToServiceOrderInformationRequiredEventInternalServerError Internal Server Error

swagger:response listenToServiceOrderInformationRequiredEventInternalServerError
*/
type ListenToServiceOrderInformationRequiredEventInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

type ListenToServiceOrderInformationRequiredEventInternalServerErrorRaw struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewListenToServiceOrderInformationRequiredEventInternalServerError creates ListenToServiceOrderInformationRequiredEventInternalServerError with default headers values
func NewListenToServiceOrderInformationRequiredEventInternalServerError() *ListenToServiceOrderInformationRequiredEventInternalServerError {

	return &ListenToServiceOrderInformationRequiredEventInternalServerError{}
}

// NewListenToServiceOrderInformationRequiredEventInternalServerError creates ListenToServiceOrderInformationRequiredEventInternalServerErrorRaw with default headers values
func NewListenToServiceOrderInformationRequiredEventInternalServerErrorRaw() *ListenToServiceOrderInformationRequiredEventInternalServerErrorRaw {

	return &ListenToServiceOrderInformationRequiredEventInternalServerErrorRaw{}
}

// WithPayload adds the payload to the listen to service order information required event internal server error response
func (o *ListenToServiceOrderInformationRequiredEventInternalServerError) WithPayload(payload *models.Error) *ListenToServiceOrderInformationRequiredEventInternalServerError {
	o.Payload = payload
	return o
}

// WithPayload adds the payload to the listen to service order information required event internal server error response
func (o *ListenToServiceOrderInformationRequiredEventInternalServerErrorRaw) WithPayload(payload interface{}) *ListenToServiceOrderInformationRequiredEventInternalServerErrorRaw {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the listen to service order information required event internal server error response
func (o *ListenToServiceOrderInformationRequiredEventInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListenToServiceOrderInformationRequiredEventInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WriteResponseRaw to the client
func (o *ListenToServiceOrderInformationRequiredEventInternalServerErrorRaw) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
