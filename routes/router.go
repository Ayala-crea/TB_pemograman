package routes

import (
	"Ayala-Crea/TB-pemograman/controller"
	"Ayala-Crea/TB-pemograman/middleware"


	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) *mux.Router {
	router := mux.NewRouter()

	// Middleware CORS
	router.Use(middleware.CORSMiddleware)

	// Handler untuk rute register
	router.HandleFunc("/register", controller.RegisterUser(db)).Methods("POST")
	router.HandleFunc("/login", controller.LoginUser(db)).Methods("POST")

	router.HandleFunc("/item", controller.CreateItems(db)).Methods("POST")
	router.HandleFunc("/item", controller.GetItem(db)).Methods("GET")
	router.HandleFunc("/item/id", controller.GetItemById(db)).Methods("GET")
	router.HandleFunc("/item/categoryId", controller.GetItemByIdCategory(db)).Methods("GET")

	router.HandleFunc("/item/category", controller.CreateCategory(db)).Methods("POST")
	router.HandleFunc("/item/category", controller.GetCategoryById(db)).Methods("GET")

	return router
}
