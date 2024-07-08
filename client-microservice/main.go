package main

import (
	"client-microservice/routes"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	mux := chi.NewRouter()
	routes.RegisterRoutes(mux)
	http.ListenAndServe(":8081", mux)
}
