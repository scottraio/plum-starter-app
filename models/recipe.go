package models

import (
	"context"

	models "github.com/scottraio/plum/models"
	store "github.com/scottraio/plum/vectorstores"
)

type RecipeModel struct {
	models.Model
	models.Markdown
}

func Recipe() *models.Model {
	var recipe *RecipeModel

	// create the model
	recipe = &RecipeModel{
		// Model is the base model that you want to use
		Model: models.Model{
			Name: "Recipe",

			// How to understand the data
			HowTo: `
				You are given excerpts from many markdown files. 
				Summarize the data and return a synopsis of the query. 
			`,

			// Train is a function that returns the data to be used for training
			Train: func(ctx context.Context) []store.Vector {
				// Hook up a vector store to use this feature.
				return []store.Vector{}
			},

			// Return is a function that returns the result that you want to use in your prompt
			Return: func(query string) string {
				return "Recipe data"
			},
		},
	}

	return &recipe.Model
}
