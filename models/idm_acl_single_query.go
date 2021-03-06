// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// IdmACLSingleQuery idm ACL single query
// swagger:model idmACLSingleQuery
type IdmACLSingleQuery struct {

	// actions
	Actions []*IdmACLAction `json:"Actions"`

	// node ids
	NodeIds []string `json:"NodeIDs"`

	// role ids
	RoleIds []string `json:"RoleIDs"`

	// workspace ids
	WorkspaceIds []string `json:"WorkspaceIDs"`

	// not
	Not bool `json:"not,omitempty"`
}

// Validate validates this idm ACL single query
func (m *IdmACLSingleQuery) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IdmACLSingleQuery) validateActions(formats strfmt.Registry) error {

	if swag.IsZero(m.Actions) { // not required
		return nil
	}

	for i := 0; i < len(m.Actions); i++ {
		if swag.IsZero(m.Actions[i]) { // not required
			continue
		}

		if m.Actions[i] != nil {
			if err := m.Actions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Actions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *IdmACLSingleQuery) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IdmACLSingleQuery) UnmarshalBinary(b []byte) error {
	var res IdmACLSingleQuery
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
