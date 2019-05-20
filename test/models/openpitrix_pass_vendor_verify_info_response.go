// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixPassVendorVerifyInfoResponse openpitrix pass vendor verify info response
// swagger:model openpitrixPassVendorVerifyInfoResponse
type OpenpitrixPassVendorVerifyInfoResponse struct {

	// id of user passed
	UserID string `json:"user_id,omitempty"`
}

// Validate validates this openpitrix pass vendor verify info response
func (m *OpenpitrixPassVendorVerifyInfoResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixPassVendorVerifyInfoResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixPassVendorVerifyInfoResponse) UnmarshalBinary(b []byte) error {
	var res OpenpitrixPassVendorVerifyInfoResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
