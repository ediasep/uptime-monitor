package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"uptime-monitor/model"
	"uptime-monitor/service"
)

// UserHandler handles user-related requests
type UserHandler struct {
	service service.UserService
}

// NewUserHandler creates a new UserHandler
func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// GetUser godoc
// @Summary      Get a user
// @Description  get user by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   query      int  true  "User ID"
// @Success      200  {object}   model.User
// @Failure      400  {string}   string "Invalid id"
// @Failure      404  {string}   string "User not found"
// @Router       /users [get]
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	// Assuming URL: /users?id=1
	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	user, err := h.service.GetUser(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// CreateUser creates a new user
// @Summary      Create a user
// @Description  create user with given payload
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      model.User  true  "User payload"
// @Success      201  {string}   string "User created"
// @Failure      400  {string}   string "Invalid request body"
// @Failure      500  {string}   string "Failed to create user"
// @Router       /users [post]
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.service.CreateUser(&user); err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
