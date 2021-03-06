// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixDescribeRuntimesResponse openpitrix describe runtimes response
// swagger:model openpitrixDescribeRuntimesResponse
type OpenpitrixDescribeRuntimesResponse struct {

	// runtime set
	RuntimeSet OpenpitrixDescribeRuntimesResponseRuntimeSet `json:"runtime_set"`

	// total count of runtime
	TotalCount int64 `json:"total_count,omitempty"`
}

// Validate validates this openpitrix describe runtimes response
func (m *OpenpitrixDescribeRuntimesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixDescribeRuntimesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixDescribeRuntimesResponse) UnmarshalBinary(b []byte) error {
	var res OpenpitrixDescribeRuntimesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
