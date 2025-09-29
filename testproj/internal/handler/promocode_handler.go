package handler

import (
	"encoding/json"
	"net/http"

	"testproj/internal/middleware"
	"testproj/internal/model"
	"testproj/internal/service"
)

type PromoCodeHandler struct {
	promoCodeService service.PromoCodeService
}

func NewPromoCodeHandler(s service.PromoCodeService) *PromoCodeHandler {
	return &PromoCodeHandler{promoCodeService: s}
}

func (h *PromoCodeHandler) GetPromoCodeByUserID(w http.ResponseWriter, r *http.Request) {
	UserID, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		http.Error(w, "Invalid user", http.StatusUnauthorized)
		return
	}

	promocodes, err := h.promoCodeService.GetPromoCodeByUserID(UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(promocodes)
}

func (h *PromoCodeHandler) GetPromoCodeByID(w http.ResponseWriter, r *http.Request) {
	_, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		http.Error(w, "Invalid user", http.StatusUnauthorized)
		return
	}

	var req struct {
		Discount_ID int `json:"discount_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "h.promocode.GetPromoCodeByID: invalid JSON body", http.StatusBadRequest)
		return
	}

	promocode, err := h.promoCodeService.GetPromoCodeByID(req.Discount_ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(promocode)
}

func (h *PromoCodeHandler) CreatePromoCode(w http.ResponseWriter, r *http.Request) {
	UserID, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		http.Error(w, "Invalid user", http.StatusUnauthorized)
		return
	}

	var promocode model.Discount_code
	if err := json.NewDecoder(r.Body).Decode(&promocode); err != nil {
		http.Error(w, "h.promocode.createPromoCode: invalid JSON", http.StatusBadRequest)
		return
	}

	promocode.Seller_ID = UserID

	if err := h.promoCodeService.CreatePromoCode(promocode); err != nil {
		http.Error(w, "h.promocode.createPromoCode: failed to create promocode", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(map[string]string{
		"message": "promocode created",
	})
}

func (h *PromoCodeHandler) UpdatePromoCode(w http.ResponseWriter, r *http.Request) {
	_, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		http.Error(w, "Invalid user", http.StatusUnauthorized)
		return
	}

	var promocode model.Discount_code
	if err := json.NewDecoder(r.Body).Decode(&promocode); err != nil {
		http.Error(w, "h.promocode.UpdatePromoCode: invalid JSON body", http.StatusBadRequest)
		return
	}

	if err := h.promoCodeService.UpdatePromoCode(promocode); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(map[string]string{
		"message": "promocode update",
	})
}

func (h *PromoCodeHandler) DeactivatePromoCode(w http.ResponseWriter, r *http.Request) {
	_, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		http.Error(w, "Invalid user", http.StatusUnauthorized)
		return
	}

	var req struct {
		Discount_ID int `json:"discount_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "h.promocode.DeactivatePromoCode: invalid JSON body", http.StatusBadRequest)
		return
	}

	if err := h.promoCodeService.DeactivatePromoCode(req.Discount_ID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(map[string]string{
		"message": "promocode update",
	})
}
