// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RestSchedulerActionsResponse rest scheduler actions response
// swagger:model restSchedulerActionsResponse
type RestSchedulerActionsResponse struct {

	// List of all registered actions
	Actions map[string]RestActionDescription `json:"Actions,omitempty"`
}

// Validate validates this rest scheduler actions response
func (m *RestSchedulerActionsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestSchedulerActionsResponse) validateActions(formats strfmt.Registry) error {

	if swag.IsZero(m.Actions) { // not required
		return nil
	}

	for k := range m.Actions {

		if err := validate.Required("Actions"+"."+k, "body", m.Actions[k]); err != nil {
			return err
		}
		if val, ok := m.Actions[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RestSchedulerActionsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestSchedulerActionsResponse) UnmarshalBinary(b []byte) error {
	var res RestSchedulerActionsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
