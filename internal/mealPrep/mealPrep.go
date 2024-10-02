package mealPrep

type MealPrepService struct {
	Repository MealPrepRepository
}

func NewMealPrepService(repository MealPrepRepository) *MealPrepService {
	return &MealPrepService{
		Repository: repository,
	}
}

func (m *MealPrepService) CreateRecipe(recipe *Recipe) (*Recipe, error) {
	result, err := m.Repository.CreateRecipe(recipe)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (m *MealPrepService) GetRecipeByID(id int) (*Recipe, error) {
	recipe, err := m.Repository.GetRecipeByID(id)
	if err != nil {
		return nil, err
	}

	return recipe, nil
}

func (m *MealPrepService) GetRecipePaginated(page int, pageSize int) (*[]Recipe, error) {
	recipe, err := m.Repository.GetRecipePaginated(page, pageSize)
	if err != nil {
		return nil, err
	}

	return recipe, nil
}

func (m *MealPrepService) CreateMealPrep(mealPrep *MealPrep) (*MealPrep, error) {
	result, err := m.Repository.CreateMealPrep(mealPrep)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (m *MealPrepService) AddRecipeToMealprep(mealPrepId int, recipeId int) error {
	err := m.Repository.AddRecipeToMealPrep(mealPrepId, recipeId)
	if err != nil {
		return err
	}

	return nil
}
