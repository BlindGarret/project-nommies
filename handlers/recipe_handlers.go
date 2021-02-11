package handlers

import (
	"github.com/BlindGarret/project-nommies/generated/api/models"
	"github.com/BlindGarret/project-nommies/generated/api/restapi/operations/recipes"
	"github.com/go-openapi/runtime/middleware"
)

// NewListRecipesHandler constructs a ListRecipeHandlerFunc
func NewListRecipesHandler() recipes.ListRecipesHandlerFunc {
	return func(params recipes.ListRecipesParams) middleware.Responder {
		return recipes.NewListRecipesOK().WithPayload(&models.ListRecipesResponse{
			Next:    "something",
			Recipes: []models.Recipes{},
		})
	}
}
