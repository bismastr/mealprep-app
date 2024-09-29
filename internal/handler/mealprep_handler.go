package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/bismastr/mealprep-app/internal/mealPrep"
)

type MealPrepController struct {
	MealPrepService *mealPrep.MealPrepService
}

func NewMealPrepController(repository mealPrep.MealPrepRepository) *MealPrepController {
	return &MealPrepController{
		MealPrepService: mealPrep.NewMealPrepService(repository),
	}
}

func (m *MealPrepController) CreateRecipe(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	newRecipe := &mealPrep.Recipe{
		Name: r.FormValue("name"),
	}

	result, err := m.MealPrepService.CreateRecipe(newRecipe)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "unable to create recipe", "message": err.Error()})
	}

	json.NewEncoder(w).Encode(result)
}

func (m *MealPrepController) GetRecipeByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid Id format", "message": err.Error()})
	}

	recipe, err := m.MealPrepService.GetRecipeByID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "unable to create recipe", "message": err.Error()})
	}

	json.NewEncoder(w).Encode(recipe)
}
