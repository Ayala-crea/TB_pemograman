package repository

import (
	"Ayala-Crea/TB-pemograman/model"
	"fmt"

	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

var secretKey = []byte("your_secret_key")

func CreateItems(db *gorm.DB, item *model.Items) error {
	result := db.Create(item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func VerifyToken(tokenStr string) (*model.JWTClaims, error) {
	// Parse token dengan klaim JWT
	token, err := jwt.ParseWithClaims(tokenStr, &model.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	// Pastikan token valid dan klaim bertipe JWTClaims
	if claims, ok := token.Claims.(*model.JWTClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid token")
	}
}

func GetItem(db *gorm.DB) (*model.Items, error) {
	var item model.Items
	result := db.Find(&item)
	if result.Error != nil {
		return nil, result.Error
	}
	return &item, nil
}

func GetItemById(db *gorm.DB, id string) (model.Items, error) {
	var item model.Items
	if err := db.First(&item, "id_items = ?", id).Error; err != nil {
		return item, err
	}
	return item, nil
}


func GetItemByIdCategory(db *gorm.DB, id string) (model.Items, error) {
	var item model.Items
	// Mencari item berdasarkan id_category
	if err := db.Where("category_id = ?", id).First(&item).Error; err != nil {
		return item, err
	}
	return item, nil
}