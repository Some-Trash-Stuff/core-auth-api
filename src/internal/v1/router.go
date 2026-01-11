package v1

import (
	"net/http"

	"core-auth-api/internal/v1/controller"

	"github.com/go-chi/chi/v5"
)

var (
	authController = controller.AuthController{}
)

func NewRouter() http.Handler {
	r := chi.NewRouter()

	r.Route("/v1", func(r chi.Router) {
		r.Get("/ping", authController.Ping)
	})

	return r
}
