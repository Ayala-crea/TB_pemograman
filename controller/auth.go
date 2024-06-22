package controller

import (
	"Ayala-Crea/TB-pemograman/model"
	repo "Ayala-Crea/TB-pemograman/repository"
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// RegisterUser adalah handler untuk rute register
func RegisterUser(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user model.Users

		// Parsing body request ke struct User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Request Body Invalid", http.StatusBadRequest)
			return
		}

		// Menyimpan user ke database menggunakan repository
		if err := repo.CreateUser(db, &user); err != nil {
			http.Error(w, "Gagal Register", http.StatusInternalServerError)
			return
		}

		// Mengatur response status dan message
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"message": "Berhasil Register!"})
	}
}

func LoginUser(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user model.Users

		// Parsing request body ke struct User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Error parsing request data", http.StatusBadRequest)
			return
		}

		// Validasi email atau username harus ada
		if user.Email == "" && user.Username == "" {
			http.Error(w, "Request must include email or username", http.StatusBadRequest)
			return
		}

		// Mendapatkan data user dari database
		var userData *model.Users
		var err error

		if user.Email != "" {
			userData, err = repo.GetUserByEmail(db, user.Email)
		} else if user.Username != "" {
			userData, err = repo.GetUserByUsername(db, user.Username)
		}

		if err != nil || userData == nil {
			http.Error(w, "Username or password incorrect", http.StatusUnauthorized)
			return
		}

		// Verifikasi password
		if err := bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(user.Password)); err != nil {
			http.Error(w, "Username or password incorrect", http.StatusUnauthorized)
			return
		}

		// Generate JWT token
		token, err := repo.GenerateToken(userData)
		if err != nil {
			http.Error(w, "Failed to generate token", http.StatusInternalServerError)
			return
		}

		// Mendapatkan data user detail berdasarkan ID
		detailedUserData, err := repo.GetUserById(db, userData.IdUser)
		if err != nil {
			http.Error(w, "Failed to retrieve user data", http.StatusInternalServerError)
			return
		}

		// Membuat response sukses dengan token dan data user
		response := map[string]interface{}{
			"token": token,
			"user": map[string]interface{}{
				"id_user":  detailedUserData.IdUser,
				"id_role":  detailedUserData.IdRole,
				"nama":     detailedUserData.Nama,
				"username": detailedUserData.Username,
				"email":    detailedUserData.Email,
				// Tambahkan detail user lain di sini jika perlu
			},
		}

		// Mengirim response dalam format JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}