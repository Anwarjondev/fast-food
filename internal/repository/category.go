package repository

import "github.com/Anwarjondev/fast-food/internal/db"

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func GetAllCategories() ([]Category, error) {
	var categories []Category
	err := db.DB.Select(&categories, "SELECT id, name FROM category")
	return categories, err
}

func GetCategoryById(id int) (Category, error) {
	var category Category
	err := db.DB.Get(&category, "SELECT id, name FROM category WHERE id = $1", id)
	return category, err
}
