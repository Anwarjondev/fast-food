package repository

import "github.com/Anwarjondev/fast-food/internal/db"

type Food struct {
	ID         int     `json:"id" db:"id"`
	Name       string  `json:"name" db:"name"`
	Price      float64 `json:"price" db:"price"`
	CategoryID int     `json:"category_id" db:"category_id"`
	ImageURL   string  `json:"img_url" db:"img_url"`
	CountFood  int     `json:"count_food" db:"count_food"`
}

func GetFoodsByCategory(categoryID int) ([]Food, error) {
	var foods []Food
	err := db.DB.Select(&foods, `
		SELECT name, price, category_id, img_url, count_food
		FROM food 
		WHERE category_id = $1
	`, categoryID)
	return foods, err
}
