// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ServiceInstanceSchemaObject service instance schema object
// swagger:model ServiceInstanceSchemaObject
type ServiceInstanceSchemaObject struct {

	// create
	Create *SchemaParameters `json:"create,omitempty"`

	// update
	Update *SchemaParameters `json:"update,omitempty"`
}

// Validate validates this service instance schema object
func (m *ServiceInstanceSchemaObject) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceInstanceSchemaObject) validateCreate(formats strfmt.Registry) error {

	if swag.IsZero(m.Create) { // not required
		return nil
	}

	if m.Create != nil {
		if err := m.Create.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("create")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceInstanceSchemaObject) validateUpdate(formats strfmt.Registry) error {

	if swag.IsZero(m.Update) { // not required
		return nil
	}

	if m.Update != nil {
		if err := m.Update.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("update")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceInstanceSchemaObject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceInstanceSchemaObject) UnmarshalBinary(b []byte) error {
	var res ServiceInstanceSchemaObject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
