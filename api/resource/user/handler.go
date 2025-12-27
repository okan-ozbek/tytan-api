package user

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type API struct {
	repository *UserRepository
	db         *sql.DB
}

func NewUserHandler(db *sql.DB) *API {
	return &API{
		repository: NewUserRepository(db),
		db:         db,
	}
}

func (a *API) List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users, err := a.repository.List()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Failed to retrieve users")
		return
	}

	if len(users) == 0 {
		json.NewEncoder(w).Encode([]*UserDTO{})
		return
	}

	dtos := make([]*UserDTO, len(users))
	for i, v := range users {
		dtos[i] = ToUserDTO(v)
	}

	if err := json.NewEncoder(w).Encode(dtos); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Failed to encode users")
		return
	}
}

func (a *API) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userForm := &UserForm{}
	_ = json.NewDecoder(r.Body).Decode(&userForm)

	if err := a.repository.Create(ToUserModel(userForm)); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Failed to create user")
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (a *API) Read(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid user ID")
		return
	}

	user, err := a.repository.Read(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("User not found")
		return
	}

	dto := ToUserDTO(user)
	if err := json.NewEncoder(w).Encode(dto); err != nil {
		http.Error(w, "Failed to encode user", http.StatusInternalServerError)
		return
	}
}

func (a *API) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid user ID")
		return
	}

	var userForm *UserForm
	_ = json.NewDecoder(r.Body).Decode(&userForm)

	if err := a.repository.Update(id, ToUserModel(userForm)); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Failed to update user")
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (a *API) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid user ID")
		return
	}

	_, err = a.db.Exec("DELETE FROM user WHERE id = ?", id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Failed to delete user")
		return
	}

	w.WriteHeader(http.StatusOK)
}
