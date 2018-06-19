// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// RestDeleteResponse Generic Message
// swagger:model restDeleteResponse
type RestDeleteResponse struct {

	// num rows
	NumRows int64 `json:"NumRows,string,omitempty"`

	// success
	Success bool `json:"Success,omitempty"`
}

// Validate validates this rest delete response
func (m *RestDeleteResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RestDeleteResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestDeleteResponse) UnmarshalBinary(b []byte) error {
	var res RestDeleteResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}