package routes

import "github.com/go-chi/chi"

func RegisterRoutes(mux *chi.Mux) {
	mux.Get("/client/viewProfile", clientProfile)
	mux.Get("/client/viewProfile/{clientId}", clientProfileByID)
}
