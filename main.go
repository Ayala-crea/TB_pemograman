package main

import (
	"Ayala-Crea/TB-pemograman/config"
	"Ayala-Crea/TB-pemograman/routes"
	"log"
	"net/http"
)

func main() {
	// Koneksi ke database
	db := config.CreateDBConnection()

	// Membuat router
	router := routes.NewRouter(db)

	// Menjalankan server
	log.Println("Server berjalan pada port 3000")
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		log.Fatalf("Server gagal berjalan: %s", err)
	}
}
