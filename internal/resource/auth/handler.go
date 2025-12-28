package auth

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"tytan-api/internal/resource/user"

	"github.com/go-playground/validator/v10"

	validatorUtil "tytan-api/util/validator"
)

type API struct {
	validator  *validator.Validate
	repository *user.UserRepository
	database   *sql.DB
}

func NewAuthHandler(validator *validator.Validate, database *sql.DB) *API {
	return &API{
		validator:  validator,
		repository: user.NewUserRepository(database),
		database:   database,
	}
}

func (a *API) Login(w http.ResponseWriter, r *http.Request) {
	// Todo add actual login functionality
	w.Header().Set("Content-Type", "application/json")

	authForm := &AuthForm{}
	if err := json.NewDecoder(r.Body).Decode(&authForm); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(validatorUtil.ToErrResponse(err))
		return
	}

	if err := a.validator.Struct(authForm); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(validatorUtil.ToErrResponse(err))
		return
	}

	user, err := a.repository.FindByCredentials(authForm.Username, authForm.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Failed to fetch user")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (a *API) Logout(w http.ResponseWriter, r *http.Request) {
	// Todo add actual logout functionality
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Successfully logged out")
}
