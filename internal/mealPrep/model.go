package mealPrep

type AppUser struct {
	ID    int    `db:"id"`
	Name  string `db:"name" json:"name,omitempty"`
	Email string `db:"email" json:"email,omitempty"`
}

type Recipe struct {
	ID             int          `db:"id" json:"id,omitempty"`
	Name           string       `db:"name" json:"name,omitempty"`
	DifficultyTier int          `db:"difficulty_tier" json:"difficulty_tier,omitempty"`
	Rating         int          `db:"rating" json:"rating,omitempty"`
	CostTier       int          `db:"cost_tier" json:"cost_tier,omitempty"`
	Ingredients    []Ingredient `json:"ingredients,omitempty"`
	Description    string       `db:"description" json:"description,omitempty"`
}

type Ingredient struct {
	ID       int    `db:"id" json:"id,omitempty"`
	RecipeID int    `db:"recipe_id" json:"recipe_id,omitempty"`
	Name     string `db:"name" json:"name,omitempty"`
	Quantity string `db:"quantity" json:"quantity,omitempty"`
	Unit     string `db:"unit" json:"unit,omitempty"`
}

type MealPrep struct {
	ID     int    `db:"id" json:"id,omitempty"`
	UserID int    `db:"user_id" json:"user_id,omitempty"`
	Name   string `db:"name" json:"name,omitempty"`
}

type MealPrepRecipe struct {
	MealPrepID int `db:"meal_prep_id" json:"meal_prep_id,omitempty"`
	RecipeID   int `db:"recipe_id" json:"recipe_id,omitempty"`
}

type CreateMealPrepRequest struct {
	Name      string `json:"name,omitempty"`
	UserID    int    `json:"user_id,omitempty"`
	RecipeIds []int  `json:"recipe_ids,omitempty"`
}
