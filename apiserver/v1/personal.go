// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Personal 个人基础信息
//
// swagger:model Personal
type Personal struct {

	// 出生年月
	Birth string `json:"birth"`

	// 性别
	Gender string `json:"gender"`

	// id
	// Read Only: true
	ID string `json:"id"`

	// 姓名
	Name string `json:"name"`

	// 用户id
	Userid string `json:"userid"`
}

// Validate validates this personal
func (m *Personal) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this personal based on the context it is used
func (m *Personal) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Personal) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Personal) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Personal) UnmarshalBinary(b []byte) error {
	var res Personal
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
