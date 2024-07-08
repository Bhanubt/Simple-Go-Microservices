package routes

import "github.com/go-chi/chi"

func RegisterRoutes(mux *chi.Mux) {
	mux.Get("/employee/viewProfile", employeeProfile)
	mux.Get("/employee/viewProfile/{employeeId}", employeeProfileByID)
}
