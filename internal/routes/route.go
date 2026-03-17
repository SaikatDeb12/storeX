package routes

import (
	handler "github.com/SaikatDeb12/storeX/internal/handlers"
	"github.com/SaikatDeb12/storeX/internal/middleware"
	"github.com/go-chi/chi/v5"
)

func SetUpRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Route("/v1", func(v1 chi.Router) {
		v1.Get("/health", handler.CheckHealth)
		v1.Route("/auth", func(r chi.Router) {
			r.Post("/login", handler.Login)
			r.Post("/register", handler.Register)
			r.Group(func(r chi.Router) {
				r.Use(middleware.Authenticate)
				r.Post("/logout", handler.Logout)
			})
		})

		v1.Group(func(r chi.Router) {
			r.Use(middleware.Authenticate)
			r.Route("/users", func(r chi.Router) {
				r.Get("/", handler.GetAllUsers)
				r.Get("/{id}", handler.GetUserInfoByID)
				r.Group(func(r chi.Router) {
					r.Use(middleware.CheckUserRole("admin", "asset_manager"))
					r.Delete("/{id}", handler.DeleteUserByID)
				})
				r.Group(func(r chi.Router) {
					r.Use(middleware.CheckUserRole("admin"))
					r.Post("/assign-role/{id}", handler.AssignRole)
				})
			})
			r.Route("/assets", func(r chi.Router) {
				r.Group(func(r chi.Router) {
					r.Use(middleware.CheckUserRole("admin", "asset_manager"))
					r.Get("/", handler.FetchAssets)
					r.Post("/", handler.CreateAsset)
					r.Patch("/update/{id}", handler.UpdateAsset)
					r.Patch("/assign", handler.AssignAssets)
					r.Patch("/service/{id}", handler.SentToService)
				})
			})
		})
	})

	return router
}
