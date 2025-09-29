package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"testproj/internal/middleware"
	"testproj/internal/model"
	"testproj/internal/service"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	usersService service.UsersService
}

func NewUsersHandler(s service.UsersService) *UserHandler {
	return &UserHandler{usersService: s}
}

// ////////////////////////////// User///////////////////////////////////
func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	user, err := h.usersService.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		http.Error(w, "h.user: missing User ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "h.user: invalid User ID", http.StatusBadRequest)
		return
	}

	products, err := h.usersService.GetUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func (h *UserHandler) GetUserForPage(w http.ResponseWriter, r *http.Request) {
	UserID, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		http.Error(w, "Invalid user", http.StatusUnauthorized)
		return
	}

	products, err := h.usersService.GetUserForPage(UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func (h *UserHandler) GetUserEmail(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "h.user: invalid input", http.StatusBadRequest)
		return
	}

	token, err := h.usersService.GetUserEmail(req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if token == "" {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(token)
}

func (h *UserHandler) CheckUserEmail(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email string `json:"email"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "h.user: invalid input", http.StatusBadRequest)
		return
	}

	user, err := h.usersService.CheckUserEmail(req.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if user {
		http.Error(w, "h.user: This email has already been used", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "email is available"})
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	UserID, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		http.Error(w, "Invalid user", http.StatusUnauthorized)
		return
	}

	var user model.Users
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "h.user: invalid JSON body", http.StatusBadRequest)
		return
	}

	user.User_ID = UserID

	if err := h.usersService.UpdateUser(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "user update",
	})
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.Users
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "h.user: invalid JSON", http.StatusBadRequest)
		return
	}

	if err := h.usersService.CreateUser(user); err != nil {
		http.Error(w, "h.user: failed to create User", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "user created",
	})
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		http.Error(w, "h.user: missing User ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "h.user: invalid User ID", http.StatusBadRequest)
		return
	}

	if err := h.usersService.DeleteUser(id); err != nil {
		http.Error(w, "h.user: failed to delete User", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User deleted"))
}
