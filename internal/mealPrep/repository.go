package mealPrep

import "github.com/bismastr/mealprep-app/internal/db"

type MealPrepRepository interface {
	//Recipe
	CreateRecipe(recipe *Recipe) (*Recipe, error)
	GetRecipeByID(id int) (*Recipe, error)
	AddItemToRecipe(recipeID int, ingredient *Ingredient) (*Ingredient, error)
	GetRecipePaginated(page int, pageSize int) (*[]Recipe, error)
	//MealPrep
	CreateMealPrep(name string, userId int) (*MealPrep, error)
	AddRecipeToMealPrep(mealPrepID int, recipeID int) error
	GetIngredientsForMealPrep(mealPrepID int) (*[]Ingredient, error)
}

type MealPrepRepositoryImpl struct {
	db *db.DB
}

func NewMealPrepRepository(db *db.DB) *MealPrepRepositoryImpl {
	return &MealPrepRepositoryImpl{
		db: db,
	}
}

func (m *MealPrepRepositoryImpl) GetIngredientsForMealPrep(mealPrepID int) (*[]Ingredient, error) {
	rows, err := m.db.DbClient.Query("SELECT i.name, i.quantity, i.unit FROM ingredient i JOIN recipe r ON i.recipe_id = r.id JOIN meal_prep_recipe mpr ON mpr.recipe_id = r.id JOIN meal_prep mp ON mpr.meal_prep_id = mp.id WHERE mp.id = $1", mealPrepID)
	if err != nil {
		return nil, err
	}

	var ingredients []Ingredient
	for rows.Next() {
		var ingredient Ingredient
		err := rows.Scan(&ingredient.Name, &ingredient.Quantity, &ingredient.Unit)
		if err != nil {
			return nil, err
		}
		ingredients = append(ingredients, ingredient)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &ingredients, nil
}

// Create a meal prep
func (m *MealPrepRepositoryImpl) CreateMealPrep(name string, userId int) (*MealPrep, error) {
	var newMealPrep MealPrep
	err := m.db.DbClient.QueryRow("INSERT INTO meal_prep (user_id, name) VALUES ($1, $2) RETURNING id, user_id, name", userId, name).Scan(&newMealPrep.ID, &newMealPrep.UserID, &newMealPrep.Name)
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
		err := rows.Scan(&recipe.ID, &recipe.Name, &recipe.DifficultyTier, &recipe.CostTier, &recipe.Rating)
		if err != nil {
			return nil, err
		}

		recipes = append(recipes, recipe)
	}

	return &recipes, nil
}
