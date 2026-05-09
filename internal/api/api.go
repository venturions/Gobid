package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/venturions/gobid/internal/services"
)

type Api struct {
	Router      *chi.Mux
	UserService services.UserService
}

func (api *Api) handleCreateUser(w http.ResponseWriter, r *http.Request) {}
