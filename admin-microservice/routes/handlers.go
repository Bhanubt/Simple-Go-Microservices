package routes

import (
	"admin-microservice/models"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

func addUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}
	userId := r.Header.Get("UserID")

	err := models.CheckAdmin(userId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	var user models.User

	err = json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = user.AddUser()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Succesfully Added User"))
}

func removeUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalied Request", http.StatusBadRequest)
		return
	}

	userId := r.Header.Get("UserID")

	err := models.CheckAdmin(userId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	removeUserId := chi.URLParam(r, "userId")

	err = models.RemoveUser(removeUserId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Succesfully Removed User"))

}

func fetchClientProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalied Request", http.StatusBadRequest)
		return
	}
	userId := r.Header.Get("UserID")

	err := models.CheckAdmin(userId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	clientId := chi.URLParam(r, "clientId")

	err = models.CheckClient(clientId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	client, err := models.ViewClientProfile(clientId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	clientDetails, err := json.Marshal(client)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(clientDetails)

}

func fetchEmployeeProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalied Request", http.StatusBadRequest)
		return
	}
	userId := r.Header.Get("UserID")

	err := models.CheckAdmin(userId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	employeeId := chi.URLParam(r, "employeeId")

	err = models.CheckEmployee(employeeId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	employee, err := models.ViewClientProfile(employeeId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	employeeDetails, err := json.Marshal(employee)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(employeeDetails)

}

func viewClientProfile(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("UserID")
	err := models.CheckClient(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	client, err := models.ViewClientProfile(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	clientDetails, err := json.Marshal(client)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(clientDetails)
}

func viewClientProfileByID(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "clientId")
	adminId := r.Header.Get("UserID")
	err := models.CheckAdmin(adminId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = models.CheckClient(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	client, err := models.ViewClientProfile(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	clientDetails, err := json.Marshal(client)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(clientDetails)
}

func viewEmployeeProfile(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("UserID")
	err := models.CheckEmployee(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	employee, err := models.ViewEmployeeProfile(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	employeeDetails, err := json.Marshal(employee)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(employeeDetails)
}

func viewEmployeeProfileByID(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "employeeId")
	adminId := r.Header.Get("UserID")
	err := models.CheckAdmin(adminId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = models.CheckEmployee(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	employee, err := models.ViewClientProfile(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	employeeDetails, err := json.Marshal(employee)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(employeeDetails)
}
