package server

import (
	"net/http"

	"github.com/bismastr/mealprep-app/internal/handler"
	"github.com/bismastr/mealprep-app/internal/mealPrep"
)

func (s *Server) RegisterRoute() *http.ServeMux {
	mux := http.NewServeMux()
	mealPrepController := handler.NewMealPrepController(mealPrep.NewMealPrepRepository(s.db))

	//Recipe
	// mux.HandleFunc("POST /recipes", mealPrepController.CreateRecipe)
	// mux.HandleFunc("GET /recipes/{id}", mealPrepController.GetRecipeByID)
	mux.Handle("GET /recipes", handler.AppHandler(mealPrepController.GetRecipePaginated))
	//MealPrep
	mux.Handle("POST /mealprep", handler.AppHandler(mealPrepController.CreateMealPrep))
	mux.Handle("POST /mealprep/recipe", handler.AppHandler(mealPrepController.AddRecipeToMealPrep))
	mux.Handle("GET /mealprep/{id}", handler.AppHandler(mealPrepController.GetIngredientsForMealPrep))

	return mux
}
