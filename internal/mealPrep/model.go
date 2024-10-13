package mealPrep

type AppUser struct {
	ID    int    `db:"id"`
	Name  string `db:"name"`
	Email string `db:"email"`
}

type Recipe struct {
	ID             int    `db:"id" json:"id"`
	Name           string `db:"name" json:"name"`
	DifficultyTier int    `db:"difficulty_tier" json:"difficulty_tier"`
	Rating         int    `db:"rating" json:"rating"`
	CostTier       int    `db:"cost_tier" json:"cost_tier"`
}

type Ingredient struct {
	ID       int    `db:"id"`
	RecipeID int    `db:"recipe_id"`
	Name     string `db:"name"`
	Quantity string `db:"quantity"`
	Unit     string `db:"unit"`
}

type MealPrep struct {
	ID     int    `db:"id" json:"id"`
	UserID int    `db:"user_id" json:"user_id"`
	Name   string `db:"name" json:"name"`
}

type MealPrepRecipe struct {
	MealPrepID int `db:"meal_prep_id"`
	RecipeID   int `db:"recipe_id"`
}

type CreateMealPrepRequest struct {
	Name      string `json:"name"`
	UserID    int    `json:"user_id"`
	RecipeIds []int  `json:"recipe_ids"`
}
