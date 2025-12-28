package food

import "time"

type FoodDTO struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

type FoodForm struct {
	Name        string `json:"name" validate:"required,alpha_space"`
	Description string `json:"description" validate:"max=500"`
}

type Food struct {
	ID          int
	Name        string
	Description string
	CreatedAt   time.Time
}

func ToFoodDTO(food *Food) *FoodDTO {
	return &FoodDTO{
		Name:        food.Name,
		Description: food.Description,
	}
}

func ToFoodModel(form *FoodForm) *Food {
	return &Food{
		Name:        form.Name,
		Description: form.Description,
		CreatedAt:   time.Now(),
	}
}
