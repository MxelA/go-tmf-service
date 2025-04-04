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

// ServiceOrderRelationship Linked service order to the one containing this attribute
//
// swagger:model ServiceOrderRelationship
type ServiceOrderRelationship struct {

	// When sub-classing, this defines the super-class
	AtBaseType *string `json:"@baseType,omitempty" bson:"@baseType,omitempty"`

	// The entity type of the related order
	AtReferredType *string `json:"@referredType,omitempty" bson:"@referredType,omitempty"`

	// A URI to a JSON-Schema file that defines additional attributes and relationships
	// Format: uri
	AtSchemaLocation *strfmt.URI `json:"@schemaLocation,omitempty" bson:"@schemaLocation,omitempty"`

	// When sub-classing, this defines the sub-class Extensible name
	AtType *string `json:"@type,omitempty" bson:"@type,omitempty"`

	// A hyperlink to the related order
	Href *string `json:"href,omitempty" bson:"href,omitempty"`

	// The id of the related order
	// Required: true
	ID *string `json:"id" bson:"id"`

	// The type of related order, such as: [dependency] if the order needs to be [not started] until another order item is complete (a service order in this case) or [cross-ref] to keep track of the source order (a productOrder)
	// Required: true
	RelationshipType *string `json:"relationshipType" bson:"relationshipType,omitempty"`
}

// Validate validates this service order relationship
func (m *ServiceOrderRelationship) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAtSchemaLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelationshipType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceOrderRelationship) validateAtSchemaLocation(formats strfmt.Registry) error {
	if swag.IsZero(m.AtSchemaLocation) { // not required
		return nil
	}

	if err := validate.FormatOf("@schemaLocation", "body", "uri", m.AtSchemaLocation.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ServiceOrderRelationship) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *ServiceOrderRelationship) validateRelationshipType(formats strfmt.Registry) error {

	if err := validate.Required("relationshipType", "body", m.RelationshipType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this service order relationship based on context it is used
func (m *ServiceOrderRelationship) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServiceOrderRelationship) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceOrderRelationship) UnmarshalBinary(b []byte) error {
	var res ServiceOrderRelationship
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
