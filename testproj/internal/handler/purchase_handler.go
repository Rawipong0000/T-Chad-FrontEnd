package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"testproj/internal/middleware"
	"testproj/internal/model"
	"testproj/internal/service"
)

type PurchaseHandler struct {
	purchaseService service.PurchaseService
}

func NewPurchaseHandler(s service.PurchaseService) *PurchaseHandler {
	return &PurchaseHandler{purchaseService: s}
}

type CreateTransactionRequest struct {
	Transaction     model.Transaction       `json:"transaction"`
	Sub_Transaction []model.Sub_Transaction `json:"sub_transaction"`
	Purchasing      []model.Purchasing      `json:"purchasing"`
	CartIds         []int                   `json:"cartIds"`
	DiscountCode    []string                `json:"discountcode"`
}

func (h *PurchaseHandler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var req CreateTransactionRequest

	UserID, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		http.Error(w, "Invalid user", http.StatusUnauthorized)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON format", http.StatusBadRequest)
		return
	}

	err := h.purchaseService.CreateTransaction(req.Purchasing, req.Sub_Transaction, req.Transaction, req.CartIds, req.DiscountCode, UserID)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, service.ErrAddress) ||
			errors.Is(err, service.ErrSold) ||
			errors.Is(err, service.ErrDelete) {
			status = http.StatusBadRequest
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json") // ✅ แจ้ง Content-Type
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(map[string]string{
		"message": "transaction created",
	})
}

func (h *PurchaseHandler) RedeemCode(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Discount_code string `json:"discount_code"`
		Seller_ID     int    `json:"seller_id"`
		Total         float32
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "h.RedeemCode: invalid input", http.StatusBadRequest)
		return
	}

	row, err := h.purchaseService.RedeemCode(req.Discount_code, req.Seller_ID, req.Total)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, service.ErrInvalidCode) ||
			errors.Is(err, service.ErrInvalidSeller) ||
			errors.Is(err, service.ErrBelowMinimum) ||
			errors.Is(err, service.ErrExceedLimit) {
			status = http.StatusBadRequest
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(row)
}
