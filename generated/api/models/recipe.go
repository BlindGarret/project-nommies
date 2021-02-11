// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Recipe recipe
//
// swagger:model Recipe
type Recipe struct {

	// List of categories this recipe belongs to.
	Category []string `json:"category"`

	// Unique recipe ID.
	// Required: true
	ID *int64 `json:"id"`

	// URL for display image of recipe.
	ImageURL string `json:"imageUrl,omitempty"`

	// Name of recipe.
	// Required: true
	Name *string `json:"name"`

	// Direct URL to recipe.
	RecipeURL string `json:"recipeUrl,omitempty"`
}

// Validate validates this recipe
func (m *Recipe) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Recipe) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *Recipe) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this recipe based on context it is used
func (m *Recipe) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Recipe) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Recipe) UnmarshalBinary(b []byte) error {
	var res Recipe
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
