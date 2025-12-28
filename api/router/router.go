package router

import (
	"database/sql"
	"tytan-api/api/resource/food"
	"tytan-api/api/resource/health"
	"tytan-api/api/resource/user"

	"github.com/go-playground/validator/v10"

	"github.com/gorilla/mux"
)

func NewRouter(validator *validator.Validate, db *sql.DB) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/health", health.Check).Methods("GET")

	userHandler := user.NewUserHandler(validator, db)
	r.HandleFunc("/users", userHandler.List).Methods("GET")
	r.HandleFunc("/users/{id}", userHandler.Read).Methods("GET")
	r.HandleFunc("/users", userHandler.Create).Methods("POST")
	r.HandleFunc("/users/{id}", userHandler.Update).Methods("PUT")
	r.HandleFunc("/users/{id}", userHandler.Delete).Methods("DELETE")

	foodHandler := food.NewFoodHandler(validator, db)
	r.HandleFunc("/foods", foodHandler.List).Methods("GET")
	r.HandleFunc("/foods/{id}", foodHandler.Read).Methods("GET")
	r.HandleFunc("/foods", foodHandler.Create).Methods("POST")
	r.HandleFunc("/foods/{id}", foodHandler.Update).Methods("PUT")
	r.HandleFunc("/foods/{id}", foodHandler.Delete).Methods("DELETE")

	return r
}
