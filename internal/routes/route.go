package routes

import (
	"net/http"

	handler "github.com/SaikatDeb12/storeX/internal/handlers"
	"github.com/SaikatDeb12/storeX/internal/middleware"
	"github.com/SaikatDeb12/storeX/internal/utils"
	"github.com/go-chi/chi/v5"
)

func SetUpRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Route("/v1", func(v1 chi.Router) {
		v1.Get("/health", func(w http.ResponseWriter, r *http.Request) {
			utils.RespondJSON(w, http.StatusOK, map[string]string{
				"status": "server is running",
			})
		})
		v1.Post("/auth/login", handler.Login)
		v1.Post("/auth/register", handler.Register)
		v1.Group(func(r chi.Router) {
			r.Use(middleware.Authenticate)
			r.Route("/users", func(r chi.Router) {
				r.Get("/", handler.GetAllUsers)
			})
			r.Route("/asset", func(r chi.Router) {
				r.Get("/", handler.ShowAssets)
				r.Group(func(r chi.Router) {
					r.Use(middleware.CheckUserRole)
					r.Post("/", handler.CreateAsset)
					r.Patch("/assign", handler.AssignedAssets)
				})
			})
		})
	})

	return router
}

// TODO : when the jwt is expiring add archived_at timestamp on users
