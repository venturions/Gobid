package api

import (
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
	"github.com/venturions/gobid/internal/services"
)

type Api struct {
	Router         *chi.Mux
	UserService    services.UserService
	ProductService services.ProductService
	Sessions       *scs.SessionManager
	WsUpgrader     websocket.Upgrader
}

func (api *Api) handleCreateUser(w http.ResponseWriter, r *http.Request) {}
