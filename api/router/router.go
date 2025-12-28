package router

import (
	"database/sql"
	"tytan-api/api/resource/food"
	"tytan-api/api/resource/health"
	"tytan-api/api/resource/user"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

func NewRouter(validator *validator.Validate, database *sql.DB) *chi.Mux {
	router := chi.NewRouter()

	router.Route("/api", func(router chi.Router) {
		router.Get("/health", health.Check)

		router.Route("/v1", func(router chi.Router) {
			router.Route("/users", func(router chi.Router) {
				userHandler := user.NewUserHandler(validator, database)

				router.Get("/", userHandler.List)
				router.Post("/", userHandler.Create)

				router.Route("/{id}", func(r chi.Router) {
					router.Get("/", userHandler.Read)
					router.Put("/", userHandler.Update)
					router.Delete("/", userHandler.Delete)
				})
			})

			router.Route("/foods", func(router chi.Router) {
				foodHandler := food.NewFoodHandler(validator, database)

				router.Get("/", foodHandler.List)
				router.Post("/", foodHandler.Create)

				router.Route("/{id}", func(router chi.Router) {
					router.Get("/", foodHandler.Read)
					router.Put("/", foodHandler.Update)
					router.Delete("/", foodHandler.Delete)
				})
			})
		})
	})

	return router
}
