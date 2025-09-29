package router

import (
	"net/http"
	"testproj/internal/handler"

	"github.com/gorilla/mux"
)

func RegisterMasterRoutes(router *mux.Router, masterHandler *handler.MasterHandler) {
	router.HandleFunc("/users/profile/province", masterHandler.GetProvince).Methods(http.MethodGet)
	router.HandleFunc("/users/profile/district", masterHandler.GetDistrict).Methods(http.MethodPut)
	router.HandleFunc("/users/profile/subdistrict", masterHandler.GetSubDistrict).Methods(http.MethodPut)
}

func RegisterProductRoutes(router *mux.Router, productHandler *handler.ProductHandler) {
	router.HandleFunc("/products", productHandler.GetAllProducts).Methods(http.MethodGet)
	router.HandleFunc("/cart/products", productHandler.GetMultiProductsForCart).Methods(http.MethodPost)
	router.HandleFunc("/products/{id}", productHandler.GetProduct).Methods(http.MethodGet)
	router.HandleFunc("/products/{id}", productHandler.UpdateProduct).Methods(http.MethodPut)
	router.HandleFunc("/products", productHandler.CreateProduct).Methods(http.MethodPost)
	router.HandleFunc("/add-product", productHandler.CreatePageProduct).Methods(http.MethodPost)
	router.HandleFunc("/products/{id}", productHandler.DeleteProduct).Methods(http.MethodDelete)
}

func RegisterUsersRoutes(router *mux.Router, userHandler *handler.UserHandler) {
	router.HandleFunc("/api/users", userHandler.GetAllUsers).Methods(http.MethodGet)
	router.HandleFunc("/api/users/{id}", userHandler.GetUser).Methods(http.MethodGet)
	router.HandleFunc("/users/id", userHandler.GetUserForPage).Methods(http.MethodGet)
	router.HandleFunc("/users/profile/update", userHandler.UpdateUser).Methods(http.MethodPut)
	router.HandleFunc("/api/users/{id}", userHandler.DeleteUser).Methods(http.MethodDelete)
}

func TransactionRoutes(router *mux.Router, purchaseHandler *handler.PurchaseHandler, productHandler *handler.ProductHandler) {
	router.HandleFunc("/transaction", purchaseHandler.CreateTransaction).Methods(http.MethodPost)
	router.HandleFunc("/redeem-code", purchaseHandler.RedeemCode).Methods(http.MethodPost)
}

func HistoryRoutes(router *mux.Router, historyHandler *handler.HistoryHandler) {
	router.HandleFunc("/users/History", historyHandler.GetHistoryTransaction).Methods(http.MethodGet)
	router.HandleFunc("/users/History/update-complete", historyHandler.CompleteTransaction).Methods(http.MethodPut)
	router.HandleFunc("/users/History/update-refund", historyHandler.RefundTransaction).Methods(http.MethodPut)
	router.HandleFunc("/users/History/update-refund-approve", historyHandler.RefundApprove).Methods(http.MethodPut)
	router.HandleFunc("/users/History/update-refund-reject", historyHandler.RefundReject).Methods(http.MethodPut)
	router.HandleFunc("/users/History/update-cancel", historyHandler.CancelTransaction).Methods(http.MethodPut)
}

func RegisterHomeRoutes(router *mux.Router, userHandler *handler.UserHandler) {
	router.HandleFunc("/api/users/email", userHandler.GetUserEmail).Methods(http.MethodPost)
	router.HandleFunc("/api/users/email/{email}", userHandler.GetUserEmail).Methods(http.MethodGet)
	router.HandleFunc("/api/users/check-email", userHandler.CheckUserEmail).Methods(http.MethodPost)
	router.HandleFunc("/api/users", userHandler.CreateUser).Methods(http.MethodPost)
}

func RegisterMyShopRoutes(router *mux.Router, myShopHandler *handler.MyShopHandler, promoCodeHandler *handler.PromoCodeHandler) {
	router.HandleFunc("/users/MyShop", myShopHandler.GetShopNameByID).Methods(http.MethodGet)
	router.HandleFunc("/users/MyShop/Edit/Shopname", myShopHandler.EditShopName).Methods(http.MethodPut)
	router.HandleFunc("/users/MyShop/products", myShopHandler.GetMyShopAllProducts).Methods(http.MethodGet)
	router.HandleFunc("/users/MyShop/order-manage", myShopHandler.GetMyShopTransaction).Methods(http.MethodGet)
	router.HandleFunc("/users/MyShop/order-manage/edit-tracking", myShopHandler.EditTracking).Methods(http.MethodPut)
	router.HandleFunc("/users/MyShop/promocodes", promoCodeHandler.GetPromoCodeByUserID).Methods(http.MethodGet)
	router.HandleFunc("/users/MyShop/promocodes/get-by-id", promoCodeHandler.GetPromoCodeByID).Methods(http.MethodPut)
	router.HandleFunc("/users/MyShop/promocodes/create", promoCodeHandler.CreatePromoCode).Methods(http.MethodPost)
	router.HandleFunc("/users/MyShop/promocodes/update", promoCodeHandler.UpdatePromoCode).Methods(http.MethodPut)
	router.HandleFunc("/users/MyShop/promocodes/deactivate", promoCodeHandler.DeactivatePromoCode).Methods(http.MethodPut)
}
