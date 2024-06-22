package repository

import (
	"Ayala-Crea/TB-pemograman/model"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, user *model.Users) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	// Jika id_role tidak diisi, atur nilainya ke 2 (atau nilai default yang diinginkan)
	if user.IdRole == 0 {
		user.IdRole = 2 // atau nilai default yang diinginkan
	}

	// Simpan user ke database
	result := db.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetUserByUsername(db *gorm.DB, username string) (*model.Users, error) {
	var user model.Users
	result := db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// GetUserByEmail mengambil user berdasarkan email
func GetUserByEmail(db *gorm.DB, email string) (*model.Users, error) {
	var user model.Users
	result := db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// GetUserById mengambil user berdasarkan ID
func GetUserById(db *gorm.DB, userID uint) (*model.Users, error) {
	var user model.Users
	result := db.First(&user, userID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &user, nil
}

// GenerateToken adalah placeholder untuk fungsi yang menghasilkan token JWT
func GenerateToken(user *model.Users) (string, error) {
	// Buat klaim token dengan detail user dan klaim standar
	claims := &model.JWTClaims{
		IdUser: user.IdUser,
		IdRole: user.IdRole,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(), // Token berlaku selama 24 jam
			Issuer:    "myapp", // Pemberi token
		},
	}

	// Buat token baru dengan klaim yang sudah ditentukan
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Tandatangani token dengan kunci rahasia
	secretKey := []byte("your_secret_key") // Gantilah dengan kunci rahasia yang aman
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	// Logging untuk debug
	fmt.Println("Generated token:", tokenString)
	fmt.Printf("Claims in token: %+v\n", claims)

	return tokenString, nil
}