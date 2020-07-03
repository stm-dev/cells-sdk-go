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

// JobsContextMetaFilterType jobs context meta filter type
// swagger:model jobsContextMetaFilterType
type JobsContextMetaFilterType string

const (

	// JobsContextMetaFilterTypeRequestMeta captures enum value "RequestMeta"
	JobsContextMetaFilterTypeRequestMeta JobsContextMetaFilterType = "RequestMeta"

	// JobsContextMetaFilterTypeContextUser captures enum value "ContextUser"
	JobsContextMetaFilterTypeContextUser JobsContextMetaFilterType = "ContextUser"
)

// for schema
var jobsContextMetaFilterTypeEnum []interface{}

func init() {
	var res []JobsContextMetaFilterType
	if err := json.Unmarshal([]byte(`["RequestMeta","ContextUser"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		jobsContextMetaFilterTypeEnum = append(jobsContextMetaFilterTypeEnum, v)
	}
}

func (m JobsContextMetaFilterType) validateJobsContextMetaFilterTypeEnum(path, location string, value JobsContextMetaFilterType) error {
	if err := validate.Enum(path, location, value, jobsContextMetaFilterTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this jobs context meta filter type
func (m JobsContextMetaFilterType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateJobsContextMetaFilterTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}