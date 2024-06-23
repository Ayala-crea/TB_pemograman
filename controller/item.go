package controller

import (
	"Ayala-Crea/TB-pemograman/model"
	repo "Ayala-Crea/TB-pemograman/repository"
	"encoding/json"
	"fmt"

	"net/http"
	"strings"
	"time"

	"gorm.io/gorm"
)

func CreateItems(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header tidak ada", http.StatusBadRequest)
			return
		}

		// Format header harus "Bearer token"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Format Authorization header salah", http.StatusBadRequest)
			return
		}

		tokenStr := parts[1]

		// Verifikasi token
		claims, err := repo.VerifyToken(tokenStr)
		if err != nil {
			http.Error(w, "Token tidak valid: "+err.Error(), http.StatusUnauthorized)
			return
		}

		// Print klaim token untuk debug
		fmt.Fprintf(w, "Token valid!\n")
		fmt.Fprintf(w, "IdUser: %d, IdRole: %d\n", claims.IdUser, claims.IdRole)

		// Parsing body request ke struct Item
		var item model.Items
		if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
			http.Error(w, "Error parsing request data", http.StatusBadRequest)
			return
		}

		// Validasi data item (misalnya, pastikan Name dan Quantity tidak kosong)
		if item.Name == "" || item.Quantity <= 0 {
			http.Error(w, "Name atau Quantity tidak valid", http.StatusBadRequest)
			return
		}

		// Set CreatedAt ke waktu sekarang jika belum diatur
		if item.CreatedAt.IsZero() {
			item.CreatedAt = time.Now()
		}

		// Simpan item ke database
		if err := repo.CreateItems(db, &item); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Mengirim response sukses
		response := map[string]interface{}{
			"code":    http.StatusCreated,
			"success": true,
			"status":  "success",
			"message": "Item berhasil disimpan",
			"data":    item,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	}
}

func GetItem(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		items, err := repo.GetItem(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Mengirim response sukses
		response := map[string]interface{}{
			"code":    http.StatusOK,
			"success": true,
			"status":  "success",
			"message": "Berhasil mendapatkan data item",
			"data":    items,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

func GetItemById(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
			// Mendapatkan nilai id dari query parameter
			id := r.URL.Query().Get("id_items")
			if id == "" {
				http.Error(w, "Request must include id_item", http.StatusBadRequest)
				return
			}
	
			// Mengambil item dari repository menggunakan GORM
			item, err := repo.GetItemById(db, id)
			if err != nil {
				if err == gorm.ErrRecordNotFound {
					http.Error(w, "Item not found", http.StatusNotFound)
				} else {
					http.Error(w, "Error retrieving item: "+err.Error(), http.StatusInternalServerError)
				}
				return
			}

		// Mengirim response sukses
		response := map[string]interface{}{
			"code":    http.StatusOK,
			"success": true,
			"status":  "success",
			"message": "Item retrieved successfully",
			"data":    item,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

func GetItemByIdCategory(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("category_id")
		if id == "" {
			http.Error(w, "Request must include id_category", http.StatusBadRequest)
			return
		}

		// Mengambil item dari repository berdasarkan id_category
		item, err := repo.GetItemByIdCategory(db, id)
		if err != nil {
			// Mengirimkan response jika item tidak ditemukan
			http.Error(w, "Item not found", http.StatusNotFound)
			return
		}

		// Mengirim response sukses
		response := map[string]interface{}{
			"code":    http.StatusOK,
			"success": true,
			"status":  "success",
			"message": "Item retrieved successfully",
			"data":    item,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}
