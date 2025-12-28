package router

import (
	"database/sql"
	"tytan-api/api/resource/food"
	"tytan-api/api/resource/health"
	"tytan-api/api/resource/user"
	"tytan-api/api/router/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

func NewRouter(validator *validator.Validate, database *sql.DB) *chi.Mux {
	router := chi.NewRouter()
	router.Route("/api", func(router chi.Router) {
		router.Use(middleware.ContentTypeJSON)

		router.Get("/health", health.Check)

		router.Route("/v1", func(router chi.Router) {

			// Protected routes
			router.Group(func(router chi.Router) {
				userHandler := user.NewUserHandler(validator, database)

				router.Get("/users", userHandler.List)
				router.Post("/users", userHandler.Create)
				router.Get("/users/{id:[0-9]+}", userHandler.Read)
				router.Put("/users/{id:[0-9]+}", userHandler.Update)
				router.Delete("/users/{id:[0-9]+}", userHandler.Delete)

				foodHandler := food.NewFoodHandler(validator, database)

				router.Get("/foods", foodHandler.List)
				router.Post("/foods", foodHandler.Create)
				router.Get("/foods/{id:[0-9]+}", foodHandler.Read)
				router.Put("/foods/{id:[0-9]+}", foodHandler.Update)
				router.Delete("/foods/{id:[0-9]+}", foodHandler.Delete)
			})
		})
	})

	return router
}
