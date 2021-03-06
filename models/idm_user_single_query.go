// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// IdmUserSingleQuery idm user single query
// swagger:model idmUserSingleQuery
type IdmUserSingleQuery struct {

	// attribute any value
	AttributeAnyValue bool `json:"AttributeAnyValue,omitempty"`

	// Search on attribute
	AttributeName string `json:"AttributeName,omitempty"`

	// attribute value
	AttributeValue string `json:"AttributeValue,omitempty"`

	// Search a specific group by path
	FullPath string `json:"FullPath,omitempty"`

	// Search on group path, and if so, search recursively
	GroupPath string `json:"GroupPath,omitempty"`

	// Shortcut for pydio:profile attribute
	HasProfile string `json:"HasProfile,omitempty"`

	// Search on roles
	HasRole string `json:"HasRole,omitempty"`

	// login
	Login string `json:"Login,omitempty"`

	// node type
	NodeType IdmNodeType `json:"NodeType,omitempty"`

	// password
	Password string `json:"Password,omitempty"`

	// recursive
	Recursive bool `json:"Recursive,omitempty"`

	// Uuid
	UUID string `json:"Uuid,omitempty"`

	// not
	Not bool `json:"not,omitempty"`
}

// Validate validates this idm user single query
func (m *IdmUserSingleQuery) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNodeType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IdmUserSingleQuery) validateNodeType(formats strfmt.Registry) error {

	if swag.IsZero(m.NodeType) { // not required
		return nil
	}

	if err := m.NodeType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("NodeType")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IdmUserSingleQuery) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IdmUserSingleQuery) UnmarshalBinary(b []byte) error {
	var res IdmUserSingleQuery
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
