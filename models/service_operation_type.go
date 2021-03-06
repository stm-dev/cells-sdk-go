// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// ServiceOperationType service operation type
// swagger:model serviceOperationType
type ServiceOperationType string

const (

	// ServiceOperationTypeOR captures enum value "OR"
	ServiceOperationTypeOR ServiceOperationType = "OR"

	// ServiceOperationTypeAND captures enum value "AND"
	ServiceOperationTypeAND ServiceOperationType = "AND"
)

// for schema
var serviceOperationTypeEnum []interface{}

func init() {
	var res []ServiceOperationType
	if err := json.Unmarshal([]byte(`["OR","AND"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serviceOperationTypeEnum = append(serviceOperationTypeEnum, v)
	}
}

func (m ServiceOperationType) validateServiceOperationTypeEnum(path, location string, value ServiceOperationType) error {
	if err := validate.Enum(path, location, value, serviceOperationTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this service operation type
func (m ServiceOperationType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateServiceOperationTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
