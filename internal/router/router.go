package router

import (
	"database/sql"
	"tytan-api/internal/resource/health"
	"tytan-api/internal/resource/user"
	"tytan-api/internal/router/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

func NewRouter(validator *validator.Validate, database *sql.DB) *chi.Mux {
	router := chi.NewRouter()
	router.Route("/api", func(router chi.Router) {
		router.Use(middleware.ContentTypeJSON)

		router.Get("/health", health.Check)

		router.Route("/v1", func(router chi.Router) {
			// Guest routes
			router.Group(func(router chi.Router) {
				// Middlewares
				//

				// router.Post("/auth/login", handler)
			})

			// Authenticated routes
			router.Group(func(router chi.Router) {
				// Middlewares
				//

				// router.Post("/auth/logout", handler)
				// router.Post("/auth/refresh", handler)

				userHandler := user.NewUserHandler(validator, database)

				router.Get("/users", userHandler.List)
				router.Post("/users", userHandler.Create)
				router.Get("/users/{id:[0-9]+}", userHandler.Read)
				router.Put("/users/{id:[0-9]+}", userHandler.Update)
				router.Delete("/users/{id:[0-9]+}", userHandler.Delete)
			})
		})
	})

	return router
}
