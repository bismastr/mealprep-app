package mealPrep

// AppUser represents a user in the application
type AppUser struct {
	ID    int    `db:"id"`    // Primary key for the AppUser table
	Name  string `db:"name"`  // Name of the user
	Email string `db:"email"` // Email of the user
}

// Recipe represents a recipe created by a user
type Recipe struct {
	ID   int    `db:"id" json:"id"`     // Primary key for the Recipe table
	Name string `db:"name" json:"name"` // Name of the recipe
}

// Ingredient represents an ingredient associated with a recipe
type Ingredient struct {
	ID       int    `db:"id"`        // Primary key for the Ingredient table
	RecipeID int    `db:"recipe_id"` // Foreign key referencing the Recipe table
	Name     string `db:"name"`      // Name of the ingredient
	Quantity string `db:"quantity"`  // e.g., "1/4", "1 1/2", "40"
	Unit     string `db:"unit"`      // Unit of the quantity, e.g., "teaspoon", "cup", "g"
}

// MealPrep represents a meal preparation plan created by a user
type MealPrep struct {
	ID     int    `db:"id" json:"id"`           // Primary key for the MealPrep table
	UserID int    `db:"user_id" json:"user_id"` // Foreign key referencing the AppUser table
	Name   string `db:"name" json:"name"`       // Name of the meal prep plan
}

// MealPrepRecipe is a join table linking meal preps and recipes
type MealPrepRecipe struct {
	MealPrepID int `db:"meal_prep_id"` // Foreign key referencing the MealPrep table
	RecipeID   int `db:"recipe_id"`    // Foreign key referencing the Recipe table
}

type GetRecipeResponse struct {
	Data     []Recipe `json:"data"`
	Message  string   `json:"message"`
	Page     int      `json:"page"`
	PageSize int      `json:"pageSize"`
}

type GetRecipeRequest struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}
