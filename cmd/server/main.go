package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"log"
	"net/http"

	"sample-api/pkg/api"
)

func main() {
	r := chi.NewRouter()

	r.Get("/recipes", api.GetRecipes)
	r.Get("/recipes/{id}", api.GetRecipe)
	r.Post("/recipes", api.CreateRecipe)
	r.Put("/recipes/{id}", api.UpdateRecipe)
	r.Delete("/recipes/{id}", api.DeleteRecipe)

	port := 8080
	fmt.Printf("Server running on :%d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
