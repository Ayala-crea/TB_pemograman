package controller

import (
	"Ayala-Crea/TB-pemograman/model"
	repo "Ayala-Crea/TB-pemograman/repository"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"gorm.io/gorm"
)

func CreateCategory(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header is missing", http.StatusBadRequest)
			return
		}

		// Format header harus "Bearer token"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Authorization header format is incorrect", http.StatusBadRequest)
			return
		}

		tokenStr := parts[1]

		// Verifikasi token
		claims, err := repo.VerifyToken(tokenStr)
		if err != nil {
			http.Error(w, "Invalid token: "+err.Error(), http.StatusUnauthorized)
			return
		}

		// Debug print klaim token (sebaiknya dihapus atau diganti dengan logging dalam produksi)
		fmt.Printf("Token valid! IdUser: %d, IdRole: %d\n", claims.IdUser, claims.IdRole)

		var category model.Categories
		if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
			http.Error(w, "Error parsing request data", http.StatusBadRequest)
			return
		}

		// Tambahkan validasi untuk kategori jika diperlukan
		// Misalnya, memeriksa apakah semua field yang diperlukan telah diisi
		if category.Name == "" {
			http.Error(w, "Category name is required", http.StatusBadRequest)
			return
		}

		if err := repo.CreateCategory(db, &category); err != nil {
			http.Error(w, "Failed to create category: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Mengirim response sukses
		response := map[string]interface{}{
			"code":    http.StatusCreated,
			"success": true,
			"status":  "success",
			"message": "Category created successfully",
			"data":    category,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	}
}

func GetCategoryById(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id_category")
		if id == "" {
			http.Error(w, "Request must include id_category", http.StatusBadRequest)
			return
		}

		// Mengambil kategori dari repository berdasarkan id_category
		category, err := repo.GetCategoryById(db, id)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				http.Error(w, "Category not found", http.StatusNotFound)
			} else {
				http.Error(w, "Error retrieving category: "+err.Error(), http.StatusInternalServerError)
			}
			return
		}

		// Mengirim response sukses
		response := map[string]interface{}{
			"code":    http.StatusOK,
			"success": true,
			"status":  "success",
			"message": "Category retrieved successfully",
			"data":    category,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}
