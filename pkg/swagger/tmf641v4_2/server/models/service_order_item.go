// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ServiceOrderItem service order item
//
// swagger:model ServiceOrderItem
type ServiceOrderItem struct {

	// When sub-classing, this defines the super-class
	AtBaseType *string `json:"@baseType,omitempty" bson:"@baseType,omitempty"`

	// A URI to a JSON-Schema file that defines additional attributes and relationships
	// Format: uri
	AtSchemaLocation strfmt.URI `json:"@schemaLocation,omitempty" bson:"@schemaLocation,omitempty"`

	// When sub-classing, this defines the sub-class Extensible name
	AtType string `json:"@type,omitempty" bson:"@type,omitempty"`

	// The action to be carried out on the Service. Can be: add, modify, delete, noChange
	// Required: true
	// Enum: ["add","modify","delete","noChange"]
	Action *string `json:"action" bson:"action,omitempty"`

	// An appointment that was set up with a related party for this order item
	Appointment *AppointmentRef `json:"appointment,omitempty"`

	// the error(s) cause an order item status change
	ErrorMessage []*ServiceOrderItemErrorMessage `json:"errorMessage,omitempty" bson:"errorMessage,omitempty"`

	// Identifier of the individual line item
	// Required: true
	ID *string `json:"id" bson:"id,omitempty"`

	// A list of modification items provided for the service update when serviceOrderItem action is modify
	ModifyPath []*JSONPatch `json:"modifyPath,omitempty"`

	// Quantity ordered
	Quantity *int64 `json:"quantity,omitempty" bson:"quantity,omitempty"`

	// The Service to be acted on by the order item
	// Required: true
	Service *ServiceRefOrValue `json:"service"`

	// A list of order items embedded to this order item
	ServiceOrderItem []*ServiceOrderItem `json:"serviceOrderItem,omitempty" bson:"serviceOrderItem,omitempty"`

	// A list of order items related to this order item
	ServiceOrderItemRelationship []*ServiceOrderItemRelationship `json:"serviceOrderItemRelationship,omitempty" bson:"serviceOrderItemRelationship,omitempty"`

	// State of the order item: described in the state machine diagram. This is the requested state.
	// Enum: ["acknowledged","rejected","pending","held","inProgress","cancelled","completed","failed","assessingCancellation","pendingCancellation","partial"]
	State *string `json:"state,omitempty" bson:"state,omitempty"`
}

// Validate validates this service order item
func (m *ServiceOrderItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAtSchemaLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAppointment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateErrorMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifyPath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateService(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceOrderItem(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceOrderItemRelationship(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceOrderItem) validateAtSchemaLocation(formats strfmt.Registry) error {
	if swag.IsZero(m.AtSchemaLocation) { // not required
		return nil
	}

	if err := validate.FormatOf("@schemaLocation", "body", "uri", m.AtSchemaLocation.String(), formats); err != nil {
		return err
	}

	return nil
}

var serviceOrderItemTypeActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["add","modify","delete","noChange"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serviceOrderItemTypeActionPropEnum = append(serviceOrderItemTypeActionPropEnum, v)
	}
}

const (

	// ServiceOrderItemActionAdd captures enum value "add"
	ServiceOrderItemActionAdd string = "add"

	// ServiceOrderItemActionModify captures enum value "modify"
	ServiceOrderItemActionModify string = "modify"

	// ServiceOrderItemActionDelete captures enum value "delete"
	ServiceOrderItemActionDelete string = "delete"

	// ServiceOrderItemActionNoChange captures enum value "noChange"
	ServiceOrderItemActionNoChange string = "noChange"
)

// prop value enum
func (m *ServiceOrderItem) validateActionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, serviceOrderItemTypeActionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ServiceOrderItem) validateAction(formats strfmt.Registry) error {

	if err := validate.Required("action", "body", m.Action); err != nil {
		return err
	}

	// value enum
	if err := m.validateActionEnum("action", "body", *m.Action); err != nil {
		return err
	}

	return nil
}

func (m *ServiceOrderItem) validateAppointment(formats strfmt.Registry) error {
	if swag.IsZero(m.Appointment) { // not required
		return nil
	}

	if m.Appointment != nil {
		if err := m.Appointment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("appointment")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("appointment")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceOrderItem) validateErrorMessage(formats strfmt.Registry) error {
	if swag.IsZero(m.ErrorMessage) { // not required
		return nil
	}

	for i := 0; i < len(m.ErrorMessage); i++ {
		if swag.IsZero(m.ErrorMessage[i]) { // not required
			continue
		}

		if m.ErrorMessage[i] != nil {
			if err := m.ErrorMessage[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("errorMessage" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("errorMessage" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ServiceOrderItem) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *ServiceOrderItem) validateModifyPath(formats strfmt.Registry) error {
	if swag.IsZero(m.ModifyPath) { // not required
		return nil
	}

	for i := 0; i < len(m.ModifyPath); i++ {
		if swag.IsZero(m.ModifyPath[i]) { // not required
			continue
		}

		if m.ModifyPath[i] != nil {
			if err := m.ModifyPath[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("modifyPath" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("modifyPath" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ServiceOrderItem) validateService(formats strfmt.Registry) error {

	if err := validate.Required("service", "body", m.Service); err != nil {
		return err
	}

	if m.Service != nil {
		if err := m.Service.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("service")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("service")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceOrderItem) validateServiceOrderItem(formats strfmt.Registry) error {
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

func (m *ServiceOrderItem) validateServiceOrderItemRelationship(formats strfmt.Registry) error {
	if swag.IsZero(m.ServiceOrderItemRelationship) { // not required
		return nil
	}

	for i := 0; i < len(m.ServiceOrderItemRelationship); i++ {
		if swag.IsZero(m.ServiceOrderItemRelationship[i]) { // not required
			continue
		}

		if m.ServiceOrderItemRelationship[i] != nil {
			if err := m.ServiceOrderItemRelationship[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("serviceOrderItemRelationship" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("serviceOrderItemRelationship" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var serviceOrderItemTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["acknowledged","rejected","pending","held","inProgress","cancelled","completed","failed","assessingCancellation","pendingCancellation","partial"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serviceOrderItemTypeStatePropEnum = append(serviceOrderItemTypeStatePropEnum, v)
	}
}

const (

	// ServiceOrderItemStateAcknowledged captures enum value "acknowledged"
	ServiceOrderItemStateAcknowledged string = "acknowledged"

	// ServiceOrderItemStateRejected captures enum value "rejected"
	ServiceOrderItemStateRejected string = "rejected"

	// ServiceOrderItemStatePending captures enum value "pending"
	ServiceOrderItemStatePending string = "pending"

	// ServiceOrderItemStateHeld captures enum value "held"
	ServiceOrderItemStateHeld string = "held"

	// ServiceOrderItemStateInProgress captures enum value "inProgress"
	ServiceOrderItemStateInProgress string = "inProgress"

	// ServiceOrderItemStateCancelled captures enum value "cancelled"
	ServiceOrderItemStateCancelled string = "cancelled"

	// ServiceOrderItemStateCompleted captures enum value "completed"
	ServiceOrderItemStateCompleted string = "completed"

	// ServiceOrderItemStateFailed captures enum value "failed"
	ServiceOrderItemStateFailed string = "failed"

	// ServiceOrderItemStateAssessingCancellation captures enum value "assessingCancellation"
	ServiceOrderItemStateAssessingCancellation string = "assessingCancellation"

	// ServiceOrderItemStatePendingCancellation captures enum value "pendingCancellation"
	ServiceOrderItemStatePendingCancellation string = "pendingCancellation"

	// ServiceOrderItemStatePartial captures enum value "partial"
	ServiceOrderItemStatePartial string = "partial"
)

// prop value enum
func (m *ServiceOrderItem) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, serviceOrderItemTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ServiceOrderItem) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", *m.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this service order item based on the context it is used
func (m *ServiceOrderItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAppointment(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateErrorMessage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateModifyPath(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateService(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateServiceOrderItem(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateServiceOrderItemRelationship(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceOrderItem) contextValidateAppointment(ctx context.Context, formats strfmt.Registry) error {

	if m.Appointment != nil {

		if swag.IsZero(m.Appointment) { // not required
			return nil
		}

		if err := m.Appointment.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("appointment")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("appointment")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceOrderItem) contextValidateErrorMessage(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ErrorMessage); i++ {

		if m.ErrorMessage[i] != nil {

			if swag.IsZero(m.ErrorMessage[i]) { // not required
				return nil
			}

			if err := m.ErrorMessage[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("errorMessage" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("errorMessage" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ServiceOrderItem) contextValidateModifyPath(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ModifyPath); i++ {

		if m.ModifyPath[i] != nil {

			if swag.IsZero(m.ModifyPath[i]) { // not required
				return nil
			}

			if err := m.ModifyPath[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("modifyPath" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("modifyPath" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ServiceOrderItem) contextValidateService(ctx context.Context, formats strfmt.Registry) error {

	if m.Service != nil {

		if err := m.Service.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("service")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("service")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceOrderItem) contextValidateServiceOrderItem(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ServiceOrderItem) contextValidateServiceOrderItemRelationship(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ServiceOrderItemRelationship); i++ {

		if m.ServiceOrderItemRelationship[i] != nil {

			if swag.IsZero(m.ServiceOrderItemRelationship[i]) { // not required
				return nil
			}

			if err := m.ServiceOrderItemRelationship[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("serviceOrderItemRelationship" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("serviceOrderItemRelationship" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceOrderItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceOrderItem) UnmarshalBinary(b []byte) error {
	var res ServiceOrderItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
