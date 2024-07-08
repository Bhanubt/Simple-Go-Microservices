package main

import (
	"admin-microservice/db"
	"admin-microservice/routes"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	db.Connect()
	mux := chi.NewRouter()
	routes.RegisterRoutes(mux)
	http.ListenAndServe(":8080", mux)
}
