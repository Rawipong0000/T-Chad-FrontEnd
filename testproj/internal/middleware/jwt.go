package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("FHJKDFjksdczvfvbd45")

// ใช้ context key แบบ type-safe
type contextKey string

const UserIDKey contextKey = "user_id"

func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// ดึง Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Missing or invalid Authorization header", http.StatusUnauthorized)
			return
		}

		// ตัด prefix "Bearer "
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// แกะและตรวจสอบ token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// ตรวจว่าลายเซ็นใช้ HMAC
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return jwtSecret, nil
		})

		if err != nil {
			fmt.Println("JWT parse error:", err)
			http.Error(w, "Token parse error", http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			fmt.Println("JWT not valid")
			http.Error(w, "Invalid token signature or expired", http.StatusUnauthorized)
			return
		}

		// ดึง user_id จาก claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || claims["user_id"] == nil {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}

		userID := int(claims["user_id"].(float64))

		// แนบ user_id เข้า context
		ctx := context.WithValue(r.Context(), UserIDKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
