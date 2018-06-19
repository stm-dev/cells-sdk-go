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

// IdmPolicyEffect idm policy effect
// swagger:model idmPolicyEffect
type IdmPolicyEffect string

const (

	// IdmPolicyEffectUnknown captures enum value "unknown"
	IdmPolicyEffectUnknown IdmPolicyEffect = "unknown"

	// IdmPolicyEffectDeny captures enum value "deny"
	IdmPolicyEffectDeny IdmPolicyEffect = "deny"

	// IdmPolicyEffectAllow captures enum value "allow"
	IdmPolicyEffectAllow IdmPolicyEffect = "allow"
)

// for schema
var idmPolicyEffectEnum []interface{}

func init() {
	var res []IdmPolicyEffect
	if err := json.Unmarshal([]byte(`["unknown","deny","allow"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		idmPolicyEffectEnum = append(idmPolicyEffectEnum, v)
	}
}

func (m IdmPolicyEffect) validateIdmPolicyEffectEnum(path, location string, value IdmPolicyEffect) error {
	if err := validate.Enum(path, location, value, idmPolicyEffectEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this idm policy effect
func (m IdmPolicyEffect) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateIdmPolicyEffectEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}