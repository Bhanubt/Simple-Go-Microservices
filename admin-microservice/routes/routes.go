package routes

import "github.com/go-chi/chi"

func RegisterRoutes(mux *chi.Mux) {
	mux.Post("/admin/addUser", addUser)
	mux.Delete("/admin/removeUser/{userId}", removeUser)
	mux.Get("/admin/viewClientProfile/{clientId}", fetchClientProfile)
	mux.Get("/admin/viewEmployeeProfile/{employeeId}", fetchEmployeeProfile)

	mux.Get("/client-admin/viewProfile", viewClientProfile)
	mux.Get("/client-admin/viewProfile/{clientId}", viewClientProfileByID)
	mux.Get("/employee-admin/viewProfile", viewEmployeeProfile)
	mux.Get("/employee-admin/viewProfile/{employeeId}", viewEmployeeProfileByID)
}
