package http

import (
	"encoding/json"
	"net/http"

	"github.com/AntonyIS/todo-app-one/app"
	"github.com/gorilla/mux"
)

type UserHandler interface {
	CreateUser(http.ResponseWriter, *http.Request)
	ReadUser(http.ResponseWriter, *http.Request)
	ReadAllUsers(http.ResponseWriter, *http.Request)
	UpdateUser(http.ResponseWriter, *http.Request)
	DeleteUser(http.ResponseWriter, *http.Request)
}

type userhandler struct {
	userService app.UserService
}

func NewUserHandler(userService app.UserService) UserHandler {
	return &userhandler{
		userService,
	}
}

func (h *userhandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user app.User
	json.NewDecoder(r.Body).Decode(&user)
	h.userService.Create(&user)
	json.NewEncoder(w).Encode(user)

}

func (h *userhandler) ReadUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	user, err := h.userService.Read(id)
	if err != nil {
		json.NewEncoder(w).Encode("User Not Found!")
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
func (h *userhandler) ReadAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.userService.ReadAll()
	if err != nil {
		json.NewEncoder(w).Encode("User Not Found!")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)

}

func (h *userhandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	user := app.User{}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)

}

func (h *userhandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	h.userService.Delete(id)
	json.NewEncoder(w).Encode("User Deleted Successfully!")
}
