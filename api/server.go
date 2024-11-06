package api

import (
	"github.com/go-chi/chi/v5"
	v1 "instashop/api/v1"
	"net/http"
)

func NewServer(repo v1.Repository) http.Handler {
	mux := chi.NewRouter()
	v1.AddRoutes(mux, repo)
	return mux
}
