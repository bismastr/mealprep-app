package handler

import (
	"encoding/json"
	"errors"
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

func (m *MealPrepController) GetRecipePaginated(w http.ResponseWriter, r *http.Request) *AppError {
	page, err := getIntFormValue(r, "page", 1)
	if err != nil {
		return &AppError{
			Code:    http.StatusBadRequest,
			Message: "invalid page format",
			Error:   err,
		}
	}

	pageSize, err := getIntFormValue(r, "pageSize", 10)
	if err != nil {
		return &AppError{
			Code:    http.StatusBadRequest,
			Message: "invalid pageSize format",
			Error:   err,
		}
	}

	recipes, err := m.MealPrepService.GetRecipePaginated(page, pageSize)
	if err != nil {
		return &AppError{
			Code:    http.StatusBadRequest,
			Message: "invalid unable to get recipe",
			Error:   err,
		}
	}

	if len(*recipes) < 1 {
		return &AppError{
			Code:    http.StatusNotFound,
			Message: "No data found in page " + strconv.Itoa(page),
			Error:   errors.New("no data found"),
		}

	}

	json.NewEncoder(w).Encode(recipes)

	return nil
}

func getIntFormValue(r *http.Request, key string, defaultValue int) (int, error) {
	valueStr := r.FormValue(key)

	if valueStr == "" {
		return defaultValue, nil
	}
	return strconv.Atoi(valueStr)
}
