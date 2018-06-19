// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// AuthLdapMemberOfMapping auth ldap member of mapping
// swagger:model authLdapMemberOfMapping
type AuthLdapMemberOfMapping struct {

	// group filter
	GroupFilter *AuthLdapSearchFilter `json:"GroupFilter,omitempty"`

	// mapping
	Mapping *AuthLdapMapping `json:"Mapping,omitempty"`

	// pydio member of attribute
	PydioMemberOfAttribute string `json:"PydioMemberOfAttribute,omitempty"`

	// pydio member of value format
	PydioMemberOfValueFormat string `json:"PydioMemberOfValueFormat,omitempty"`

	// real member of
	RealMemberOf bool `json:"RealMemberOf,omitempty"`

	// real member of attribute
	RealMemberOfAttribute string `json:"RealMemberOfAttribute,omitempty"`

	// real member of value format
	RealMemberOfValueFormat string `json:"RealMemberOfValueFormat,omitempty"`

	// support nested group
	SupportNestedGroup bool `json:"SupportNestedGroup,omitempty"`
}

// Validate validates this auth ldap member of mapping
func (m *AuthLdapMemberOfMapping) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGroupFilter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMapping(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuthLdapMemberOfMapping) validateGroupFilter(formats strfmt.Registry) error {

	if swag.IsZero(m.GroupFilter) { // not required
		return nil
	}

	if m.GroupFilter != nil {
		if err := m.GroupFilter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("GroupFilter")
			}
			return err
		}
	}

	return nil
}

func (m *AuthLdapMemberOfMapping) validateMapping(formats strfmt.Registry) error {

	if swag.IsZero(m.Mapping) { // not required
		return nil
	}

	if m.Mapping != nil {
		if err := m.Mapping.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Mapping")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AuthLdapMemberOfMapping) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuthLdapMemberOfMapping) UnmarshalBinary(b []byte) error {
	var res AuthLdapMemberOfMapping
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}