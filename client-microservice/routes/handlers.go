package routes

import (
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi"
)

func clientProfile(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("UserID")
	url := "http://localhost:8080/client-admin/viewProfile"
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	req.Header.Set("UserID", userId)

	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)

}

func clientProfileByID(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("UserID")
	clientId := chi.URLParam(r, "clientId")
	url := fmt.Sprintf("http://localhost:8080/client-admin/viewProfile/%s", clientId)
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	req.Header.Set("UserID", userId)

	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}
