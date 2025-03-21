// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// ServiceOperatingStatusType Valid values for the Operating status of the service
//
// swagger:model ServiceOperatingStatusType
type ServiceOperatingStatusType string

func NewServiceOperatingStatusType(value ServiceOperatingStatusType) *ServiceOperatingStatusType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ServiceOperatingStatusType.
func (m ServiceOperatingStatusType) Pointer() *ServiceOperatingStatusType {
	return &m
}

const (

	// ServiceOperatingStatusTypePending captures enum value "pending"
	ServiceOperatingStatusTypePending ServiceOperatingStatusType = "pending"

	// ServiceOperatingStatusTypeConfigured captures enum value "configured"
	ServiceOperatingStatusTypeConfigured ServiceOperatingStatusType = "configured"

	// ServiceOperatingStatusTypeStarting captures enum value "starting"
	ServiceOperatingStatusTypeStarting ServiceOperatingStatusType = "starting"

	// ServiceOperatingStatusTypeRunning captures enum value "running"
	ServiceOperatingStatusTypeRunning ServiceOperatingStatusType = "running"

	// ServiceOperatingStatusTypeDegraded captures enum value "degraded"
	ServiceOperatingStatusTypeDegraded ServiceOperatingStatusType = "degraded"

	// ServiceOperatingStatusTypeFailed captures enum value "failed"
	ServiceOperatingStatusTypeFailed ServiceOperatingStatusType = "failed"

	// ServiceOperatingStatusTypeLimited captures enum value "limited"
	ServiceOperatingStatusTypeLimited ServiceOperatingStatusType = "limited"

	// ServiceOperatingStatusTypeStopping captures enum value "stopping"
	ServiceOperatingStatusTypeStopping ServiceOperatingStatusType = "stopping"

	// ServiceOperatingStatusTypeStopped captures enum value "stopped"
	ServiceOperatingStatusTypeStopped ServiceOperatingStatusType = "stopped"

	// ServiceOperatingStatusTypeUnknown captures enum value "unknown"
	ServiceOperatingStatusTypeUnknown ServiceOperatingStatusType = "unknown"
)

// for schema
var serviceOperatingStatusTypeEnum []interface{}

func init() {
	var res []ServiceOperatingStatusType
	if err := json.Unmarshal([]byte(`["pending","configured","starting","running","degraded","failed","limited","stopping","stopped","unknown"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serviceOperatingStatusTypeEnum = append(serviceOperatingStatusTypeEnum, v)
	}
}

func (m ServiceOperatingStatusType) validateServiceOperatingStatusTypeEnum(path, location string, value ServiceOperatingStatusType) error {
	if err := validate.EnumCase(path, location, value, serviceOperatingStatusTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this service operating status type
func (m ServiceOperatingStatusType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateServiceOperatingStatusTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this service operating status type based on context it is used
func (m ServiceOperatingStatusType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
