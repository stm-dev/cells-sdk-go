// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// RestActionDescription rest action description
// swagger:model restActionDescription
type RestActionDescription struct {

	// User-defined category to organize actions list
	Category string `json:"Category,omitempty"`

	// Long description and help text
	Description string `json:"Description,omitempty"`

	// Whether this action has a form or not
	HasForm bool `json:"HasForm,omitempty"`

	// Mdi reference name for displaying icon
	Icon string `json:"Icon,omitempty"`

	// Additional description for expected inputs
	InputDescription string `json:"InputDescription,omitempty"`

	// Human-readable label
	Label string `json:"Label,omitempty"`

	// Unique name of the action
	Name string `json:"Name,omitempty"`

	// Additional description describing the action output
	OutputDescription string `json:"OutputDescription,omitempty"`

	// Template for building a short summary of the action configuration
	SummaryTemplate string `json:"SummaryTemplate,omitempty"`

	// User-defined hexa or rgb color
	Tint string `json:"Tint,omitempty"`
}

// Validate validates this rest action description
func (m *RestActionDescription) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RestActionDescription) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestActionDescription) UnmarshalBinary(b []byte) error {
	var res RestActionDescription
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
