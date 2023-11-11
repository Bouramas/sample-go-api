package api

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
)

// Recipe struct represents a recipe entity.
type Recipe struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Ingredients []string `json:"ingredients"`
}

var Recipes = []Recipe{
	{ID: 1, Name: "Pasta Carbonara", Ingredients: []string{"Spaghetti", "Eggs", "Bacon", "Parmesan"}},
	{ID: 2, Name: "Chicken Curry", Ingredients: []string{"Chicken", "Curry Sauce", "Rice"}},
}

// GetRecipes returns the list of all recipes.
func GetRecipes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Recipes)
}

// GetRecipe returns a specific recipe by ID.
func GetRecipe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := chi.URLParam(r, "id")

	for _, recipe := range Recipes {
		if fmt.Sprint(recipe.ID) == id {
			json.NewEncoder(w).Encode(recipe)
			return
		}
	}

	http.NotFound(w, r)
}

// CreateRecipe creates a new recipe.
func CreateRecipe(w http.ResponseWriter, r *http.Request) {
	var newRecipe Recipe
	err := json.NewDecoder(r.Body).Decode(&newRecipe)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	Recipes = append(Recipes, newRecipe)
	w.WriteHeader(http.StatusCreated)
}

// UpdateRecipe updates an existing recipe by ID.
func UpdateRecipe(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	for i, recipe := range Recipes {
		if fmt.Sprint(recipe.ID) == id {
			var updatedRecipe Recipe
			err := json.NewDecoder(r.Body).Decode(&updatedRecipe)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			Recipes[i] = updatedRecipe
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	http.NotFound(w, r)
}

// DeleteRecipe deletes a recipe by ID.
func DeleteRecipe(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	for i, recipe := range Recipes {
		if fmt.Sprint(recipe.ID) == id {
			Recipes = append(Recipes[:i], Recipes[i+1:]...)
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	http.NotFound(w, r)
}
