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

func (m *MealPrepService) CreateMealPrep(mealPrep *CreateMealPrepRequest) (*MealPrep, error) {
	mp, err := m.Repository.CreateMealPrep(mealPrep.Name, mealPrep.UserID)
	if err != nil {
		return nil, err
	}

	for _, r := range mealPrep.RecipeIds {
		err = m.Repository.AddRecipeToMealPrep(mp.ID, r)
		if err != nil {
			return nil, err
		}
	}

	return mp, nil
}

func (m *MealPrepService) AddRecipeToMealprep(mealPrepId int, recipeId int) error {
	err := m.Repository.AddRecipeToMealPrep(mealPrepId, recipeId)
	if err != nil {
		return err
	}

	return nil
}

func (m *MealPrepService) GetIngredientsForMealPrep(mealPrepId int) (*Recipe, error) {
	recipe, err := m.Repository.GetIngredientsForMealPrep(mealPrepId)
	if err != nil {
		return nil, err
	}

	return recipe, nil
}
