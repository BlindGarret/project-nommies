// Code generated by go-swagger; DO NOT EDIT.

package ingredients

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/BlindGarret/project-nommies/generated/api/models"
)

// ListIngredientCategoriesHandlerFunc turns a function with the right signature into a list ingredient categories handler
type ListIngredientCategoriesHandlerFunc func(ListIngredientCategoriesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListIngredientCategoriesHandlerFunc) Handle(params ListIngredientCategoriesParams) middleware.Responder {
	return fn(params)
}

// ListIngredientCategoriesHandler interface for that can handle valid list ingredient categories params
type ListIngredientCategoriesHandler interface {
	Handle(ListIngredientCategoriesParams) middleware.Responder
}

// NewListIngredientCategories creates a new http.Handler for the list ingredient categories operation
func NewListIngredientCategories(ctx *middleware.Context, handler ListIngredientCategoriesHandler) *ListIngredientCategories {
	return &ListIngredientCategories{Context: ctx, Handler: handler}
}

/* ListIngredientCategories swagger:route GET /ingredientCategories ingredients listIngredientCategories

List all ingredient categories

*/
type ListIngredientCategories struct {
	Context *middleware.Context
	Handler ListIngredientCategoriesHandler
}

func (o *ListIngredientCategories) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListIngredientCategoriesParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// ListIngredientCategoriesOKBody list ingredient categories o k body
//
// swagger:model ListIngredientCategoriesOKBody
type ListIngredientCategoriesOKBody struct {

	// Url to next page. Empty if no additional pages.
	Next string `json:"next,omitempty"`

	// recipes
	Recipes []*models.IngredientCategory `json:"recipes"`
}

// Validate validates this list ingredient categories o k body
func (o *ListIngredientCategoriesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRecipes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListIngredientCategoriesOKBody) validateRecipes(formats strfmt.Registry) error {
	if swag.IsZero(o.Recipes) { // not required
		return nil
	}

	for i := 0; i < len(o.Recipes); i++ {
		if swag.IsZero(o.Recipes[i]) { // not required
			continue
		}

		if o.Recipes[i] != nil {
			if err := o.Recipes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listIngredientCategoriesOK" + "." + "recipes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this list ingredient categories o k body based on the context it is used
func (o *ListIngredientCategoriesOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateRecipes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListIngredientCategoriesOKBody) contextValidateRecipes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Recipes); i++ {

		if o.Recipes[i] != nil {
			if err := o.Recipes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listIngredientCategoriesOK" + "." + "recipes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListIngredientCategoriesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListIngredientCategoriesOKBody) UnmarshalBinary(b []byte) error {
	var res ListIngredientCategoriesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
