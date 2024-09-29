package server

import (
	"net/http"

	"github.com/bismastr/mealprep-app/internal/handler"
	"github.com/bismastr/mealprep-app/internal/mealPrep"
)

func (s *Server) RegisterRoute() *http.ServeMux {
	mux := http.NewServeMux()
	mealPrepController := handler.NewMealPrepController(mealPrep.NewMealPrepRepository(s.db))

	//MealPrep
	mux.HandleFunc("POST /recipes", mealPrepController.CreateRecipe)
	mux.HandleFunc("GET /recipes/{id}", mealPrepController.GetRecipeByID)

	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	return mux
}
