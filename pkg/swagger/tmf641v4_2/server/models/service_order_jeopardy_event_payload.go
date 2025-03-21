// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ServiceOrderJeopardyEventPayload The event data structure
//
// swagger:model ServiceOrderJeopardyEventPayload
type ServiceOrderJeopardyEventPayload struct {

	// The involved resource data for the event
	ServiceOrder *ServiceOrder `json:"serviceOrder,omitempty"`
}

// Validate validates this service order jeopardy event payload
func (m *ServiceOrderJeopardyEventPayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateServiceOrder(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceOrderJeopardyEventPayload) validateServiceOrder(formats strfmt.Registry) error {
	if swag.IsZero(m.ServiceOrder) { // not required
		return nil
	}

	if m.ServiceOrder != nil {
		if err := m.ServiceOrder.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("serviceOrder")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("serviceOrder")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this service order jeopardy event payload based on the context it is used
func (m *ServiceOrderJeopardyEventPayload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateServiceOrder(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceOrderJeopardyEventPayload) contextValidateServiceOrder(ctx context.Context, formats strfmt.Registry) error {

	if m.ServiceOrder != nil {

		if swag.IsZero(m.ServiceOrder) { // not required
			return nil
		}

		if err := m.ServiceOrder.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("serviceOrder")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("serviceOrder")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceOrderJeopardyEventPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceOrderJeopardyEventPayload) UnmarshalBinary(b []byte) error {
	var res ServiceOrderJeopardyEventPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
