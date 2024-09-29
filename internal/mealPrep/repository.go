package mealPrep

import "github.com/bismastr/mealprep-app/internal/db"

type MealPrepRepository interface {
	CreateRecipe(recipe *Recipe) (*Recipe, error)
	GetRecipeByID(id int) (*Recipe, error)
}

type MealPrepRepositoryImpl struct {
	db *db.DB
}

func NewMealPrepRepository(db *db.DB) *MealPrepRepositoryImpl {
	return &MealPrepRepositoryImpl{
		db: db,
	}
}

// CreateRecipe creates a new recipe and inserts it into the database.
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
