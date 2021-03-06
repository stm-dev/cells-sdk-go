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

// ListSharedResourcesRequestListShareType list shared resources request list share type
// swagger:model ListSharedResourcesRequestListShareType
type ListSharedResourcesRequestListShareType string

const (

	// ListSharedResourcesRequestListShareTypeANY captures enum value "ANY"
	ListSharedResourcesRequestListShareTypeANY ListSharedResourcesRequestListShareType = "ANY"

	// ListSharedResourcesRequestListShareTypeLINKS captures enum value "LINKS"
	ListSharedResourcesRequestListShareTypeLINKS ListSharedResourcesRequestListShareType = "LINKS"

	// ListSharedResourcesRequestListShareTypeCELLS captures enum value "CELLS"
	ListSharedResourcesRequestListShareTypeCELLS ListSharedResourcesRequestListShareType = "CELLS"
)

// for schema
var listSharedResourcesRequestListShareTypeEnum []interface{}

func init() {
	var res []ListSharedResourcesRequestListShareType
	if err := json.Unmarshal([]byte(`["ANY","LINKS","CELLS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		listSharedResourcesRequestListShareTypeEnum = append(listSharedResourcesRequestListShareTypeEnum, v)
	}
}

func (m ListSharedResourcesRequestListShareType) validateListSharedResourcesRequestListShareTypeEnum(path, location string, value ListSharedResourcesRequestListShareType) error {
	if err := validate.Enum(path, location, value, listSharedResourcesRequestListShareTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this list shared resources request list share type
func (m ListSharedResourcesRequestListShareType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateListSharedResourcesRequestListShareTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
