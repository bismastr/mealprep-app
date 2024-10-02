package mealPrep

import "github.com/bismastr/mealprep-app/internal/db"

type MealPrepRepository interface {
	//Recipe
	CreateRecipe(recipe *Recipe) (*Recipe, error)
	GetRecipeByID(id int) (*Recipe, error)
	AddItemToRecipe(recipeID int, ingredient *Ingredient) (*Ingredient, error)
	GetRecipePaginated(page int, pageSize int) (*[]Recipe, error)
	//MealPrep
	CreateMealPrep(mealPrep *MealPrep) (*MealPrep, error)
	AddRecipeToMealPrep(mealPrepID int, recipeID int) error
}

type MealPrepRepositoryImpl struct {
	db *db.DB
}

// GetIngredientsForMealPrep implements MealPrepRepository.
func (m *MealPrepRepositoryImpl) GetIngredientsForMealPrep(mealPrepID int) (*[]Ingredient, error) {
	panic("unimplemented")
}

func NewMealPrepRepository(db *db.DB) *MealPrepRepositoryImpl {
	return &MealPrepRepositoryImpl{
		db: db,
	}
}

// Create a meal prep
func (m *MealPrepRepositoryImpl) CreateMealPrep(mealPrep *MealPrep) (*MealPrep, error) {
	var newMealPrep MealPrep
	err := m.db.DbClient.QueryRow("INSERT INTO meal_prep (user_id, name) VALUES ($1, $2) RETURNING id, user_id, name", mealPrep.UserID, mealPrep.Name).Scan(&newMealPrep.ID, &newMealPrep.UserID, &newMealPrep.Name)
	if err != nil {
		return nil, err
	}

	return &newMealPrep, nil
}

// Join recipe and mealprep
func (m *MealPrepRepositoryImpl) AddRecipeToMealPrep(mealPrepID int, recipeID int) error {
	_, err := m.db.DbClient.Exec("INSERT INTO meal_prep_recipe (meal_prep_id, recipe_id) VALUES ($1, $2)", mealPrepID, recipeID)
	if err != nil {
		return err
	}

	return nil
}

// CreateRecipe creates a new recipe and inserts it into-- the database.
func (m *MealPrepRepositoryImpl) CreateRecipe(recipe *Recipe) (*Recipe, error) {
	var newRecipe Recipe
	err := m.db.DbClient.QueryRow("INSERT INTO recipe (name) VALUES ($1) RETURNING id, name", recipe.Name).Scan(&newRecipe.ID, &newRecipe.Name)
	if err != nil {
		return nil, err
	}

	return &newRecipe, nil
}

// GetRecipeByID returns a recipe by the given ID or an error if the recipe does not exist.
func (m *MealPrepRepositoryImpl) GetRecipeByID(id int) (*Recipe, error) {
	var recipe Recipe

	rows, err := m.db.DbClient.Query("SELECT * FROM recipe WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&recipe.ID, &recipe.Name)
		if err != nil {
			return nil, err
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &recipe, nil
}

// AddItemToRecipe adds a new ingredient to a recipe and returns the newly created ingredient.
func (m *MealPrepRepositoryImpl) AddItemToRecipe(recipeID int, ingredient *Ingredient) (*Ingredient, error) {
	var newIngredient Ingredient
	row := m.db.DbClient.QueryRow("INSERT INTO ingredient (recipe_id, name, quantity, unit) VALUES ($1, $2, $3, $4) RETURNING id, recipe_id, name, quantity, unit", recipeID, ingredient.Name, ingredient.Quantity, ingredient.Unit)
	err := row.Scan(&newIngredient.ID, &newIngredient.RecipeID, &newIngredient.Name, &newIngredient.Quantity, &newIngredient.Unit)
	if err != nil {
		return nil, err
	}

	return &newIngredient, nil
}

// Get Paginated Recipe
func (m *MealPrepRepositoryImpl) GetRecipePaginated(page int, pageSize int) (*[]Recipe, error) {
	offset := (page - 1) * pageSize
	limit := pageSize

	rows, err := m.db.DbClient.Query("SELECT * FROM recipe LIMIT $1 OFFSET $2", limit, offset)
	if err != nil {
		return nil, err
	}

	var recipes []Recipe
	for rows.Next() {
		var recipe Recipe
		err := rows.Scan(&recipe.ID, &recipe.Name)
		if err != nil {
			return nil, err
		}

		recipes = append(recipes, recipe)
	}

	return &recipes, nil
}
