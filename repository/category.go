package repository

import (
	"Ayala-Crea/TB-pemograman/model"

	"gorm.io/gorm"
)

func CreateCategory(db *gorm.DB, category *model.Categories)  error {
	result := db.Create(category)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetCategoryById(db *gorm.DB, id string) (model.Categories, error) {
	var category model.Categories
	if err := db.First(&category, "id_category = ?", id).Error; err != nil {
		return category, err
	}
	return category, nil
}
