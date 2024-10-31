package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/bismastr/mealprep-app/internal/mealPrep"
	"github.com/bismastr/mealprep-app/internal/utils"
)

type Response struct {
	Message    string      `json:"message"`
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data"`
}

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

func (m *MealPrepController) CreateMealPrep(w http.ResponseWriter, r *http.Request) (*AppSucces, *AppError) {
	var mealPrepRequest mealPrep.CreateMealPrepRequest
	utils.UnmarshalJSON(r, &mealPrepRequest)

	fmt.Println(mealPrepRequest.UserID)

	res, err := m.MealPrepService.CreateMealPrep(&mealPrepRequest)
	if err != nil {
		return nil, &AppError{
			Code:    http.StatusBadRequest,
			Message: "unable to create meal prep",
			Error:   err,
		}
	}

	return &AppSucces{
		Code:    http.StatusOK,
		Message: "successfully created meal prep",
		Data:    res,
	}, nil
}

func (m *MealPrepController) AddRecipeToMealPrep(w http.ResponseWriter, r *http.Request) (*AppSucces, *AppError) {
	var recipeToMealPrepRequest mealPrep.MealPrepRecipe
	utils.UnmarshalJSON(r, &recipeToMealPrepRequest)

	err := m.MealPrepService.AddRecipeToMealprep(recipeToMealPrepRequest.MealPrepID, recipeToMealPrepRequest.RecipeID)
	if err != nil {
		return nil, &AppError{
			Code:    http.StatusBadRequest,
			Message: "unable to add recipe to meal prep",
			Error:   err,
		}
	}

	return &AppSucces{
		Code:    http.StatusOK,
		Message: "successfully added recipe to meal prep",
		Data:    nil,
	}, nil
}

func (m *MealPrepController) GetRecipePaginated(w http.ResponseWriter, r *http.Request) (*AppSucces, *AppError) {
	page, err := utils.GetIntFromValue(r, "page", 1)
	if err != nil {
		return nil, &AppError{
			Code:    http.StatusBadRequest,
			Message: "invalid page format",
			Error:   err,
		}
	}

	pageSize, err := utils.GetIntFromValue(r, "pageSize", 10)
	if err != nil {
		return nil, &AppError{
			Code:    http.StatusBadRequest,
			Message: "invalid pageSize format",
			Error:   err,
		}
	}

	recipes, err := m.MealPrepService.GetRecipePaginated(page, pageSize)
	if err != nil {
		return nil, &AppError{
			Code:    http.StatusBadRequest,
			Message: "invalid unable to get recipe",
			Error:   err,
		}
	}

	if len(*recipes) < 1 {
		return nil, &AppError{
			Code:    http.StatusNotFound,
			Message: "No data found in page " + strconv.Itoa(page),
			Error:   errors.New("no data found"),
		}
	}

	return &AppSucces{
		Code:    http.StatusOK,
		Message: "success",
		Data:    recipes,
	}, nil
}

func (m *MealPrepController) GetIngredientsForMealPrep(w http.ResponseWriter, r *http.Request) (*AppSucces, *AppError) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		return nil, &AppError{
			Code:    http.StatusBadRequest,
			Message: "invalid Id format",
			Error:   err,
		}
	}
	ingredient, err := m.MealPrepService.GetIngredientsForMealPrep(id)
	if err != nil {
		return nil, &AppError{
			Code:    http.StatusBadRequest,
			Message: "unable to get ingredients for meal prep",
			Error:   err,
		}
	}

	return &AppSucces{
		Code:    http.StatusOK,
		Message: "success",
		Data:    ingredient,
	}, nil
}

func (m *MealPrepController) GetMealPrepByUserId(w http.ResponseWriter, r *http.Request) (*AppSucces, *AppError) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		return nil, &AppError{
			Code:    http.StatusBadRequest,
			Message: "invalid Id format",
			Error:   err,
		}
	}

	mealPrep, err := m.MealPrepService.GetMealPrepByUserId(id)
	if err != nil {
		return nil, &AppError{
			Code:    http.StatusBadRequest,
			Message: "unable to get ingredients for meal prep",
			Error:   err,
		}
	}

	return &AppSucces{
		Code:    http.StatusOK,
		Message: "success",
		Data:    mealPrep,
	}, nil
}
