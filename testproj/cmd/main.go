package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"testproj/redisclient"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/rs/cors"

	"testproj/internal/handler"
	"testproj/internal/middleware"
	"testproj/internal/repository"
	"testproj/internal/router"
	"testproj/internal/service"
)

func main() {
	_ = godotenv.Load()

	addr := os.Getenv("REDIS_ADDR")
	pass := os.Getenv("REDIS_PASSWORD")
	dbStr := os.Getenv("REDIS_DB")
	db, _ := strconv.Atoi(dbStr)

	fmt.Println("redis pass:", pass)

	rc, err := redisclient.New(addr, pass, db)
	if err != nil {
		log.Fatalf("connect redis failed: %v", err)
	}
	defer rc.Close()

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser, dbPass, dbHost, dbPort, dbName)

	dbpool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatal("DB connect failed:", err)
	}
	defer dbpool.Close()

	masterRepo := repository.NewMasterRepository(dbpool)
	masterService := service.NewMasterService(masterRepo, rc)
	masterHandler := handler.NewMasterHandler(masterService)

	productRepo := repository.NewProductRepository(dbpool)
	productService := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)

	usersRepo := repository.NewUsersRepository(dbpool)
	usersService := service.NewUsersService(usersRepo)
	usersHandler := handler.NewUsersHandler(usersService)

	purchaseRepo := repository.NewPurchaseRepository(dbpool)
	purchaseService := service.NewPurchaseSevice(purchaseRepo, productRepo, usersRepo)
	purchaseHandler := handler.NewPurchaseHandler(purchaseService)

	historyRepo := repository.NewHistoryRepository(dbpool)
	historyService := service.NewHistoryService(historyRepo)
	historyHandler := handler.NewHistoryHandler(historyService)

	myShopRepo := repository.NewMyShopRepository(dbpool)
	myShopService := service.NewMyShopService(myShopRepo)
	myShopHandler := handler.NewMyShopHandler(myShopService)

	promoCodeRepo := repository.NewPromoCodeRepository(dbpool)
	promoCodeService := service.NewPromoCodeService(promoCodeRepo)
	promoCodeHandler := handler.NewPromoCodeHandler(promoCodeService)

	r := mux.NewRouter()

	router.RegisterHomeRoutes(r, usersHandler)

	// 🔐 protected routes (ต้องมี JWT)
	protected := r.PathPrefix("/api").Subrouter()
	protected.Use(middleware.JWTAuthMiddleware)
	router.RegisterMasterRoutes(protected, masterHandler)
	router.RegisterUsersRoutes(protected, usersHandler)
	router.RegisterProductRoutes(protected, productHandler)
	router.TransactionRoutes(protected, purchaseHandler, productHandler)
	router.HistoryRoutes(protected, historyHandler)
	router.RegisterMyShopRoutes(protected, myShopHandler, promoCodeHandler)

	// ✅ ตั้งค่า CORS ที่นี่
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)

	fmt.Println("Listening on http://localhost:8180")
	http.ListenAndServe(":8180", handler)

	ctx := context.Background()

	if err := rc.Set(ctx, "hello", "world", time.Minute); err != nil {
		log.Fatalf("SET failed: %v", err)
	}
	val, err := rc.Get(ctx, "hello")
	if err != nil {
		log.Fatalf("GET failed: %v", err)
	}
	fmt.Println("hello =", val)
}
