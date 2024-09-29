package mealPrep

import "github.com/bismastr/mealprep-app/internal/db"

type MealPrepRepository interface {
	CreateRecipe(recipe *Recipe) error
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
func (m *MealPrepRepositoryImpl) CreateRecipe(recipe *Recipe) error {
	_, err := m.db.DbClient.Exec("INSERT INTO recipe (name) VALUES ($1)", recipe.Name)
	if err != nil {
		return err
	}

	return nil
}

// GetRecipeByID returns a recipe by the given ID or an error if the recipe does not exist.
func (m *MealPrepRepositoryImpl) GetRecipeByID(id int64) (*Recipe, error) {
	var recipe Recipe

	rows, err := m.db.DbClient.Query("SELECT * FROM Recipe WHERE id = ?", id)
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
