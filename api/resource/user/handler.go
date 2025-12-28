package user

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"

	validatorUtil "tytan-api/util/validator"
)

type API struct {
	validator  *validator.Validate
	repository *UserRepository
	db         *sql.DB
}

func NewUserHandler(validator *validator.Validate, db *sql.DB) *API {
	return &API{
		validator:  validator,
		repository: NewUserRepository(db),
		db:         db,
	}
}

// List godoc
// @Summary      List Users
// @Description  Get a list of all users
// @Tags         users
// @Produce      json
// @Success      200  {array}   UserDTO
// @Failure      500  {string}  string "Failed to retrieve users"
// @Router       /users [get]
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

// Create godoc
// @Summary      Create User
// @Description  Create a new user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      UserForm  true  "User Form"
// @Success      201   {string}  string    "Created"
// @Failure      400   {object}  validatorUtil.ErrResponse
// @Failure      400   {string}  string    "Failed to create user"
// @Router       /users [post]
func (a *API) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userForm := &UserForm{}
	_ = json.NewDecoder(r.Body).Decode(&userForm)

	if err := a.validator.Struct(userForm); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(validatorUtil.ToErrResponse(err))
		return
	}

	if err := a.repository.Create(ToUserModel(userForm)); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Failed to create user")
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// Read godoc
// @Summary      Read User
// @Description  Get a user by ID
// @Tags         users
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  UserDTO
// @Failure      400  {string}  string  "Invalid user ID"
// @Failure      404  {string}  string  "User not found"
// @Router       /users/{id} [get]
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

	w.WriteHeader(http.StatusOK)
}

// Update godoc
// @Summary      Update User
// @Description  Update an existing user by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id    path      int       true  "User ID"
// @Param        user  body      UserForm  true  "User Form"
// @Success      200   {string}  string    "OK"
// @Failure      400   {object}  validatorUtil.ErrResponse
// @Failure      400   {string}  string    "Invalid user ID"
// @Failure      500   {string}  string    "Failed to update user"
// @Router       /users/{id} [put]
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

	if err := a.validator.Struct(userForm); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(validatorUtil.ToErrResponse(err))
		return
	}

	if err := a.repository.Update(id, ToUserModel(userForm)); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Failed to update user")
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Delete godoc
// @Summary      Delete User
// @Description  Delete a user by ID
// @Tags         users
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {string}  string  "OK"
// @Failure      400  {string}  string  "Invalid user ID"
// @Failure      500  {string}  string  "Failed to delete user"
// @Router       /users/{id} [delete]
func (a *API) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid user ID")
		return
	}

	if err := a.repository.Delete(id); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Failed to delete user")
		return
	}

	w.WriteHeader(http.StatusOK)
}
