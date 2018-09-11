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

// RestSearchWorkspaceRequest Rest request for searching workspaces
// swagger:model restSearchWorkspaceRequest
type RestSearchWorkspaceRequest struct {

	// count only
	CountOnly bool `json:"CountOnly,omitempty"`

	// group by
	GroupBy int32 `json:"GroupBy,omitempty"`

	// limit
	Limit string `json:"Limit,omitempty"`

	// offset
	Offset string `json:"Offset,omitempty"`

	// operation
	Operation ServiceOperationType `json:"Operation,omitempty"`

	// queries
	Queries []*IdmWorkspaceSingleQuery `json:"Queries"`

	// resource policy query
	ResourcePolicyQuery *RestResourcePolicyQuery `json:"ResourcePolicyQuery,omitempty"`
}

// Validate validates this rest search workspace request
func (m *RestSearchWorkspaceRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOperation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQueries(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourcePolicyQuery(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestSearchWorkspaceRequest) validateOperation(formats strfmt.Registry) error {

	if swag.IsZero(m.Operation) { // not required
		return nil
	}

	if err := m.Operation.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Operation")
		}
		return err
	}

	return nil
}

func (m *RestSearchWorkspaceRequest) validateQueries(formats strfmt.Registry) error {

	if swag.IsZero(m.Queries) { // not required
		return nil
	}

	for i := 0; i < len(m.Queries); i++ {
		if swag.IsZero(m.Queries[i]) { // not required
			continue
		}

		if m.Queries[i] != nil {
			if err := m.Queries[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Queries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RestSearchWorkspaceRequest) validateResourcePolicyQuery(formats strfmt.Registry) error {

	if swag.IsZero(m.ResourcePolicyQuery) { // not required
		return nil
	}

	if m.ResourcePolicyQuery != nil {
		if err := m.ResourcePolicyQuery.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ResourcePolicyQuery")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RestSearchWorkspaceRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestSearchWorkspaceRequest) UnmarshalBinary(b []byte) error {
	var res RestSearchWorkspaceRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
