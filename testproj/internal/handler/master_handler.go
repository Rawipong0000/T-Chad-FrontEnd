package handler

import (
	"encoding/json"
	"net/http"

	"testproj/internal/middleware"
	"testproj/internal/service"
)

type MasterHandler struct {
	masterService service.MasterService
}

func NewMasterHandler(s service.MasterService) *MasterHandler {
	return &MasterHandler{masterService: s}
}

func (h *MasterHandler) GetProvince(w http.ResponseWriter, r *http.Request) {
	_, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		http.Error(w, "Invalid user", http.StatusUnauthorized)
		return
	}

	ctx := r.Context()

	province, err := h.masterService.GetProvince(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(province)
}

func (h *MasterHandler) GetDistrict(w http.ResponseWriter, r *http.Request) {
	_, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		http.Error(w, "Invalid user", http.StatusUnauthorized)
		return
	}

	type Province_ID struct {
		Province_ID int `json:"province_id"`
	}

	var payload Province_ID
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	district, err := h.masterService.GetDistrict(payload.Province_ID, ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(district)
}

func (h *MasterHandler) GetSubDistrict(w http.ResponseWriter, r *http.Request) {
	_, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		http.Error(w, "Invalid user", http.StatusUnauthorized)
		return
	}

	type District_ID struct {
		District_ID int `json:"district_id"`
	}

	var payload District_ID
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	subdistrict, err := h.masterService.GetSubDistrict(payload.District_ID, ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(subdistrict)
}
