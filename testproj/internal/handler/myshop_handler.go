package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testproj/internal/middleware"

	"testproj/internal/service"
)

type MyShopHandler struct {
	myShopService service.MyShopService
}

func NewMyShopHandler(s service.MyShopService) *MyShopHandler {
	return &MyShopHandler{myShopService: s}
}

func (h *MyShopHandler) GetShopNameByID(w http.ResponseWriter, r *http.Request) {
	UserID, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		http.Error(w, "Invalid user", http.StatusUnauthorized)
		return
	}

	MyShop, err := h.myShopService.GetShopNameByID(UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(MyShop)
}

func (h *MyShopHandler) EditShopName(w http.ResponseWriter, r *http.Request) {
	UserID, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		http.Error(w, "Invalid user", http.StatusUnauthorized)
		return
	}

	type ShopnamePayload struct {
		Shopname string `json:"shopname"`
	}

	var payload ShopnamePayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	if payload.Shopname == "" {
		http.Error(w, "Shopname is required", http.StatusBadRequest)
		return
	}

	if err := h.myShopService.EditShopName(payload.Shopname, UserID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Shopname updated"})
}

func (h *MyShopHandler) GetMyShopAllProducts(w http.ResponseWriter, r *http.Request) {
	UserID, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		http.Error(w, "Invalid user", http.StatusUnauthorized)
		return
	}

	products, err := h.myShopService.GetMyShopAllProducts(UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func (h *MyShopHandler) GetMyShopTransaction(w http.ResponseWriter, r *http.Request) {
	UserID, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		http.Error(w, "Invalid user", http.StatusUnauthorized)
		return
	}

	transactions, err := h.myShopService.GetMyShopTransaction(UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transactions)
}

func (h *MyShopHandler) EditTracking(w http.ResponseWriter, r *http.Request) {
	_, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		http.Error(w, "Invalid user", http.StatusUnauthorized)
		return
	}

	var req struct {
		Tran_id  int    `json:"tran_id"`
		Tracking string `json:"tracking"`
	}

	fmt.Println("req :", req)

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "h.EditTracking: invalid input", http.StatusBadRequest)
		return
	}

	err := h.myShopService.EditTracking(req.Tran_id, req.Tracking)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "edit tracking successful"})
}
