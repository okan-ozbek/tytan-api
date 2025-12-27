package router

import (
	"database/sql"
	"go/main/api/resource/health"
	"go/main/api/resource/user"

	"github.com/gorilla/mux"
)

func NewRouter(db *sql.DB) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/health", health.Check).Methods("GET")

	userHandler := user.NewUserHandler(db)
	r.HandleFunc("/users", userHandler.List).Methods("GET")
	r.HandleFunc("/users/{id}", userHandler.Read).Methods("GET")
	r.HandleFunc("/users", userHandler.Create).Methods("POST")
	r.HandleFunc("/users/{id}", userHandler.Update).Methods("PUT")
	r.HandleFunc("/users/{id}", userHandler.Delete).Methods("DELETE")

	return r
}
