package handler

import (
	"encoding/json"
	"net/http"
	"testproj/internal/middleware"

	"testproj/internal/service"
)

type HistoryHandler struct {
	historyService service.HistoryService
}

func NewHistoryHandler(s service.HistoryService) *HistoryHandler {
	return &HistoryHandler{historyService: s}
}

func (h *HistoryHandler) GetHistoryTransaction(w http.ResponseWriter, r *http.Request) {
	UserID, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		http.Error(w, "Invalid user", http.StatusUnauthorized)
		return
	}

	transactions, err := h.historyService.GetHistoryTransaction(UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transactions)
}

func (h *HistoryHandler) CompleteTransaction(w http.ResponseWriter, r *http.Request) {
	_, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		http.Error(w, "Invalid user", http.StatusUnauthorized)
		return
	}

	var req struct {
		SubTranID int `json:"sub_tran_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "h.CompleteTransaction: invalid input", http.StatusBadRequest)
		return
	}

	err := h.historyService.CompleteTransaction(req.SubTranID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "edit status successful"})
}

func (h *HistoryHandler) RefundTransaction(w http.ResponseWriter, r *http.Request) {
	_, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		http.Error(w, "Invalid user", http.StatusUnauthorized)
		return
	}

	var req struct {
		SubTranID int `json:"sub_tran_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "h.RefundTransaction: invalid input", http.StatusBadRequest)
		return
	}

	err := h.historyService.RefundTransaction(req.SubTranID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "edit status successful"})
}

func (h *HistoryHandler) RefundApprove(w http.ResponseWriter, r *http.Request) {
	_, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		http.Error(w, "Invalid user", http.StatusUnauthorized)
		return
	}

	var req struct {
		SubTranID int `json:"sub_tran_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "h.RefundApprove: invalid input", http.StatusBadRequest)
		return
	}

	err := h.historyService.RefundApprove(req.SubTranID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "edit status successful"})
}

func (h *HistoryHandler) RefundReject(w http.ResponseWriter, r *http.Request) {
	_, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		http.Error(w, "Invalid user", http.StatusUnauthorized)
		return
	}

	var req struct {
		SubTranID int `json:"sub_tran_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "h.RefundReject: invalid input", http.StatusBadRequest)
		return
	}

	err := h.historyService.RefundReject(req.SubTranID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "edit status successful"})
}

func (h *HistoryHandler) CancelTransaction(w http.ResponseWriter, r *http.Request) {
	_, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		http.Error(w, "Invalid user", http.StatusUnauthorized)
		return
	}

	var req struct {
		SubTranID int `json:"sub_tran_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "h.CancelTransaction: invalid input", http.StatusBadRequest)
		return
	}

	err := h.historyService.CancelTransaction(req.SubTranID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "edit status successful"})
}
