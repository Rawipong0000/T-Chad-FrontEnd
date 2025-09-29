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

type ProductHandler struct {
	productService service.ProductService
}

func NewProductHandler(s service.ProductService) *ProductHandler {
	return &ProductHandler{productService: s}
}

// //////////////////////////////Product///////////////////////////////////
func (h *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.productService.GetAllProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func (h *ProductHandler) GetMultiProductsForCart(w http.ResponseWriter, r *http.Request) {
	var productIDs []int
	if err := json.NewDecoder(r.Body).Decode(&productIDs); err != nil {
		http.Error(w, "h.product: invalid JSON product array", http.StatusBadRequest)
		return
	}
	products, err := h.productService.GetMultiProductsForCart(productIDs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		http.Error(w, "h.product: missing product ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "h.product: invalid product ID", http.StatusBadRequest)
		return
	}

	products, err := h.productService.GetProduct(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		http.Error(w, "h.product: missing product ID", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "h.product: invalid product ID", http.StatusBadRequest)
		return
	}

	var product model.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "h.product: invalid JSON body", http.StatusBadRequest)
		return
	}
	product.Product_ID = id

	if err := h.productService.UpdateProduct(product); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(map[string]string{
		"message": "product update",
	})
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product model.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "h.product: invalid JSON", http.StatusBadRequest)
		return
	}

	if err := h.productService.CreateProduct(product); err != nil {
		http.Error(w, "h.product: failed to create product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("product created"))
}

func (h *ProductHandler) CreatePageProduct(w http.ResponseWriter, r *http.Request) {
	UserID, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		http.Error(w, "Invalid user", http.StatusUnauthorized)
		return
	}

	var product model.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "h.product: invalid JSON", http.StatusBadRequest)
		return
	}

	if err := h.productService.CreatePageProduct(product, UserID); err != nil {
		http.Error(w, "h.product: failed to create product", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(map[string]string{
		"message": "pooduct created",
	})
}

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		http.Error(w, "h.product: missing product ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "h.product: invalid product ID", http.StatusBadRequest)
		return
	}

	if err := h.productService.DeleteProduct(id); err != nil {
		http.Error(w, "h.product: failed to delete product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("product deleted"))
}

// //////////////////////////////////////////////////////////////////////
