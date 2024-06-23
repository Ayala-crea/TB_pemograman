package model

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type Users struct {
	IdUser    uint      `gorm:"primaryKey;column:id_user" json:"id_user"`
	IdRole    int       `gorm:"column:id_role" json:"id_role"`
	Nama      string    `gorm:"column:nama" json:"nama"`
	Username  string    `gorm:"column:username" json:"username"`
	Password  string    `gorm:"column:password" json:"password"`
	Email     string    `gorm:"column:email" json:"email"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"-"`
}

type Items struct {
	IdItem     uint      `gorm:"primaryKey;autoIncrement;column:id_items" json:"id_items"`       // ID unik barang
	Name       string    `gorm:"column:name" json:"name"`                            // Nama barang
	Quantity   int       `gorm:"column:quantity" json:"quantity"`                    // Jumlah barang dalam stok
	CategoryID uint      `gorm:"column:category_id" json:"category_id"`              // ID kategori barang
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"` // Waktu penciptaan barang
}

type Categories struct {
	ID   uint   `gorm:"primaryKey;autoIncrement;column:id_category" json:"id_category"` // ID unik kategori
	Name string `gorm:"column:name" json:"name"`                      // Nama kategori
}

type ItemCategory struct {
	ItemID     uint `gorm:"primaryKey;column:item_id" json:"item_id"`         // ID barang
	CategoryID uint `gorm:"primaryKey;column:category_id" json:"category_id"` // ID kategori
}

type JWTClaims struct {
	IdUser uint `json:"id_user"`
	IdRole int  `json:"id_role"`
	jwt.StandardClaims
}
