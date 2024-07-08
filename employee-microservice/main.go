package main

import (
	"employee-microservice/routes"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	mux := chi.NewRouter()
	routes.RegisterRoutes(mux)
	http.ListenAndServe(":8082", mux)
}
