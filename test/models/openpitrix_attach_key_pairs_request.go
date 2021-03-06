// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixAttachKeyPairsRequest openpitrix attach key pairs request
// swagger:model openpitrixAttachKeyPairsRequest
type OpenpitrixAttachKeyPairsRequest struct {

	// ids of key pairs to attach
	KeyPairID []string `json:"key_pair_id"`

	// ids of node to attached
	NodeID []string `json:"node_id"`
}

// Validate validates this openpitrix attach key pairs request
func (m *OpenpitrixAttachKeyPairsRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKeyPairID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNodeID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenpitrixAttachKeyPairsRequest) validateKeyPairID(formats strfmt.Registry) error {

	if swag.IsZero(m.KeyPairID) { // not required
		return nil
	}

	return nil
}

func (m *OpenpitrixAttachKeyPairsRequest) validateNodeID(formats strfmt.Registry) error {

	if swag.IsZero(m.NodeID) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixAttachKeyPairsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixAttachKeyPairsRequest) UnmarshalBinary(b []byte) error {
	var res OpenpitrixAttachKeyPairsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
