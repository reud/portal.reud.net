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

// Book book
// swagger:model Book
type Book struct {

	// ID
	ID int64 `json:"ID,omitempty"`

	// href
	// Required: true
	Href *string `json:"href"`

	// irjp image source
	// Required: true
	IrjpImageSource *string `json:"irjpImageSource"`

	// tag1
	// Required: true
	Tag1 *string `json:"tag1"`

	// tag2
	Tag2 string `json:"tag2,omitempty"`

	// tag3
	Tag3 string `json:"tag3,omitempty"`

	// title
	// Required: true
	Title *string `json:"title"`

	// wsfe image source
	// Required: true
	WsfeImageSource *string `json:"wsfeImageSource"`
}

// Validate validates this book
func (m *Book) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHref(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIrjpImageSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTag1(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWsfeImageSource(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Book) validateHref(formats strfmt.Registry) error {

	if err := validate.Required("href", "body", m.Href); err != nil {
		return err
	}

	return nil
}

func (m *Book) validateIrjpImageSource(formats strfmt.Registry) error {

	if err := validate.Required("irjpImageSource", "body", m.IrjpImageSource); err != nil {
		return err
	}

	return nil
}

func (m *Book) validateTag1(formats strfmt.Registry) error {

	if err := validate.Required("tag1", "body", m.Tag1); err != nil {
		return err
	}

	return nil
}

func (m *Book) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

func (m *Book) validateWsfeImageSource(formats strfmt.Registry) error {

	if err := validate.Required("wsfeImageSource", "body", m.WsfeImageSource); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Book) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Book) UnmarshalBinary(b []byte) error {
	var res Book
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}