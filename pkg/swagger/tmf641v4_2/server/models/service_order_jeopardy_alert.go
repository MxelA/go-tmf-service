// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ServiceOrderJeopardyAlert A ServiceOrderJeopardyAlert represents a predicted exception during a service order processing that would brings risk to complete successfully the ordetr.
//
// swagger:model ServiceOrderJeopardyAlert
type ServiceOrderJeopardyAlert struct {

	// When sub-classing, this defines the super-class
	AtBaseType *string `json:"@baseType,omitempty" bson:"@baseType,omitempty"`

	// A URI to a JSON-Schema file that defines additional attributes and relationships
	// Format: uri
	AtSchemaLocation *strfmt.URI `json:"@schemaLocation,omitempty" bson:"@schemaLocation,omitempty"`

	// When sub-classing, this defines the sub-class Extensible name
	AtType *string `json:"@type,omitempty" bson:"@type,omitempty"`

	// A date time( DateTime). The date that the alert issued
	// Format: date-time
	AlertDate *strfmt.DateTime `json:"alertDate,omitempty" bson:"alertDate,omitempty"`

	//  The exception associated with this jeopardy alert
	Exception *string `json:"exception,omitempty" bson:"exception,omitempty"`

	// identifier of the JeopardyAlert
	ID *string `json:"id,omitempty" bson:"id,omitempty"`

	// A string represents the type of jeopardy/risk like Normal, Hazard, Critical, ...
	JeopardyType *string `json:"jeopardyType,omitempty" bson:"jeopardyType,omitempty"`

	// A string represents the message of the alert
	Message *string `json:"message,omitempty" bson:"message,omitempty"`

	// A string used to give a name to the jeopardy alert
	Name *string `json:"name,omitempty" bson:"name,omitempty"`

	// A list of order items corresponded to this alert
	ServiceOrderItem []*ServiceOrderItemRef `json:"serviceOrderItem,omitempty" bson:"serviceOrderItem"`
}

// Validate validates this service order jeopardy alert
func (m *ServiceOrderJeopardyAlert) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAtSchemaLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAlertDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceOrderItem(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceOrderJeopardyAlert) validateAtSchemaLocation(formats strfmt.Registry) error {
	if swag.IsZero(m.AtSchemaLocation) { // not required
		return nil
	}

	if err := validate.FormatOf("@schemaLocation", "body", "uri", m.AtSchemaLocation.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ServiceOrderJeopardyAlert) validateAlertDate(formats strfmt.Registry) error {
	if swag.IsZero(m.AlertDate) { // not required
		return nil
	}

	if err := validate.FormatOf("alertDate", "body", "date-time", m.AlertDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ServiceOrderJeopardyAlert) validateServiceOrderItem(formats strfmt.Registry) error {
	if swag.IsZero(m.ServiceOrderItem) { // not required
		return nil
	}

	for i := 0; i < len(m.ServiceOrderItem); i++ {
		if swag.IsZero(m.ServiceOrderItem[i]) { // not required
			continue
		}

		if m.ServiceOrderItem[i] != nil {
			if err := m.ServiceOrderItem[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("serviceOrderItem" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("serviceOrderItem" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this service order jeopardy alert based on the context it is used
func (m *ServiceOrderJeopardyAlert) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateServiceOrderItem(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceOrderJeopardyAlert) contextValidateServiceOrderItem(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ServiceOrderItem); i++ {

		if m.ServiceOrderItem[i] != nil {

			if swag.IsZero(m.ServiceOrderItem[i]) { // not required
				return nil
			}

			if err := m.ServiceOrderItem[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("serviceOrderItem" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("serviceOrderItem" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceOrderJeopardyAlert) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceOrderJeopardyAlert) UnmarshalBinary(b []byte) error {
	var res ServiceOrderJeopardyAlert
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
