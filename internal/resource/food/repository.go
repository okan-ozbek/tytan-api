package food

import (
	"database/sql"
)

type FoodRepository struct {
	DB *sql.DB
}

func NewFoodRepository(db *sql.DB) *FoodRepository {
	return &FoodRepository{DB: db}
}

func (r *FoodRepository) List() ([]*Food, error) {
	rows, err := r.DB.Query("SELECT id, name, description, created_at FROM foods")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var foods []*Food
	for rows.Next() {
		food := &Food{}
		if err := rows.Scan(&food.ID, &food.Name, &food.Description, &food.CreatedAt); err != nil {
			return nil, err
		}
		foods = append(foods, food)
	}
	return foods, nil
}

func (r *FoodRepository) Create(food *Food) error {
	_, err := r.DB.Exec(
		"INSERT INTO foods (name, description, created_at) VALUES (?, ?, ?)",
		food.Name,
		food.Description,
		food.CreatedAt,
	)

	return err
}

func (r *FoodRepository) Read(id int) (*Food, error) {
	row := r.DB.QueryRow("SELECT id, name, description FROM foods WHERE id = ?", id)
	food := &Food{}
	err := row.Scan(&food.ID, &food.Name, &food.Description)
	if err != nil {
		return nil, err
	}

	return food, nil
}

func (r *FoodRepository) Update(id int, food *Food) error {
	_, err := r.DB.Exec(
		"UPDATE foods SET name = ?, description = ?, created_at = ? WHERE id = ?",
		food.Name,
		food.Description,
		food.CreatedAt,
		id,
	)
	return err
}

func (r *FoodRepository) Delete(id int) error {
	_, err := r.DB.Exec("DELETE FROM foods WHERE id = ?", id)
	return err
}
