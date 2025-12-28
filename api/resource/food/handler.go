package food

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	validatorUtil "tytan-api/util/validator"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type API struct {
	validator  *validator.Validate
	repository *FoodRepository
	database   *sql.DB
}

func NewFoodHandler(validator *validator.Validate, database *sql.DB) *API {
	return &API{
		validator:  validator,
		repository: NewFoodRepository(database),
		database:   database,
	}
}

// List godoc
// @Summary      List Foods
// @Description  Get a list of all foods
// @Tags         foods
// @Produce      json
// @Success      200  {array}   FoodDTO
// @Failure      500  {string}  string "Failed to retrieve foods"
// @Router       /foods [get]
func (a *API) List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users, err := a.repository.List()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Failed to retrieve foods")
		return
	}

	if len(users) == 0 {
		json.NewEncoder(w).Encode([]*FoodDTO{})
		return
	}

	dtos := make([]*FoodDTO, len(users))
	for i, v := range users {
		dtos[i] = ToFoodDTO(v)
	}

	if err := json.NewEncoder(w).Encode(dtos); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Failed to encode foods")
		return
	}
}

// Create godoc
// @Summary      Create Food
// @Description  Create a new food item
// @Tags         foods
// @Accept       json
// @Produce      json
// @Param        food  body      FoodDTO  true  "Food to create"
// @Success      201   {string}  string   "Food created successfully"
// @Failure      400   {object}  validatorUtil.ErrResponse  "Validation errors"
// @Failure      500   {string}  string   "Failed to create food"
// @Router       /foods [post]
func (a *API) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	foodForm := &FoodForm{}
	_ = json.NewDecoder(r.Body).Decode(&foodForm)

	if err := a.validator.Struct(foodForm); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(validatorUtil.ToErrResponse(err))
		return
	}

	if err := a.repository.Create(ToFoodModel(foodForm)); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Failed to create food")
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// Read godoc
// @Summary      Read Food
// @Description  Get a food item by ID
// @Tags         foods
// @Produce      json
// @Param        id   path      int  true  "Food ID"
// @Success      200  {object}  FoodDTO
// @Failure      400  {string}  string  "Invalid food ID"
// @Failure      404  {string}  string  "Food not found"
// @Failure      500  {string}  string  "Failed to encode food"
// @Router       /foods/{id} [get]
func (a *API) Read(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid food ID")
		return
	}

	food, err := a.repository.Read(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Food not found")
		return
	}

	dto := ToFoodDTO(food)
	if err := json.NewEncoder(w).Encode(dto); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Failed to encode food")
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Update godoc
// @Summary      Update Food
// @Description  Update an existing food item by ID
// @Tags         foods
// @Accept       json
// @Produce      json
// @Param        id    path      int      true  "Food ID"
// @Param        food  body      FoodDTO  true  "Food to update"
// @Success      200   {string}  string   "Food updated successfully"
// @Failure      400   {object}  validatorUtil.ErrResponse  "Validation errors"
// @Failure      400   {string}  string   "Invalid food ID"
// @Failure      500   {string}  string   "Failed to update food"
// @Router       /foods/{id} [put]
func (a *API) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid food ID")
		return
	}

	var foodForm *FoodForm
	_ = json.NewDecoder(r.Body).Decode(&foodForm)

	if err := a.validator.Struct(foodForm); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(validatorUtil.ToErrResponse(err))
		return
	}

	if err := a.repository.Update(id, ToFoodModel(foodForm)); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Failed to update food")
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Delete godoc
// @Summary      Delete Food
// @Description  Delete a food item by ID
// @Tags         foods
// @Produce      json
// @Param        id   path      int  true  "Food ID"
// @Success      200  {string}  string  "Food deleted successfully"
// @Failure      400  {string}  string  "Invalid food ID"
// @Failure      500  {string}  string  "Failed to delete food"
// @Router       /foods/{id} [delete]
func (a *API) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid food ID")
		return
	}

	if err := a.repository.Delete(id); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Failed to delete food")
		return
	}

	w.WriteHeader(http.StatusOK)
}
