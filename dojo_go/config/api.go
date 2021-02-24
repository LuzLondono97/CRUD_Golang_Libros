package config

import (
	"dojo_go/config/database"
	"dojo_go/handler"
	"github.com/go-chi/chi"
	"net/http"
)

// New returns the API V1 Handler with configuration.
func New(conn *database.Data) http.Handler {
	router := chi.NewRouter()
	ur := handler.NewUserHandler(conn)
	router.Mount("/users", RoutesUser(ur))
	return router
}
