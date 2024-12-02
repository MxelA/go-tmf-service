// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AppointmentRef Refers an appointment, such as a Customer presentation or internal meeting or site visit
//
// swagger:model AppointmentRef
type AppointmentRef struct {

	// When sub-classing, this defines the super-class
	AtBaseType string `json:"@baseType,omitempty"`

	// The actual type of the target instance when needed for disambiguation
	AtReferredType string `json:"@referredType,omitempty"`

	// A URI to a JSON-Schema file that defines additional attributes and relationships
	// Format: uri
	AtSchemaLocation strfmt.URI `json:"@schemaLocation,omitempty"`

	// When sub-classing, this defines the sub-class Extensible name
	AtType string `json:"@type,omitempty"`

	// An explanatory text regarding the appointment made with a party
	Description string `json:"description,omitempty"`

	// The reference of the appointment
	Href string `json:"href,omitempty"`

	// The identifier of the referred appointment
	// Required: true
	ID *string `json:"id"`
}

// Validate validates this appointment ref
func (m *AppointmentRef) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAtSchemaLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppointmentRef) validateAtSchemaLocation(formats strfmt.Registry) error {
	if swag.IsZero(m.AtSchemaLocation) { // not required
		return nil
	}

	if err := validate.FormatOf("@schemaLocation", "body", "uri", m.AtSchemaLocation.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AppointmentRef) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this appointment ref based on context it is used
func (m *AppointmentRef) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AppointmentRef) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppointmentRef) UnmarshalBinary(b []byte) error {
	var res AppointmentRef
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}