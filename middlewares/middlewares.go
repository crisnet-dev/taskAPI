package middlewares

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/crisnet-dev/utils"
)

func SetCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func VerifyToken(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		authorization = strings.TrimSpace(authorization)

		if authorization == "" {
			utils.HttpError(w, "Missing token", http.StatusBadRequest)
			return
		}

		// Bearer ufurf
		if !strings.Contains(authorization, "Bearer ") {
			utils.HttpError(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		token := strings.Split(authorization, " ")[1]

		claims, err := utils.ValidateToken(token)
		if err != nil {
			log.Println(err)
			utils.HttpError(w, err.Error(), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), "email", claims["email"])
		ctx = context.WithValue(ctx, "name", claims["name"])
		ctx = context.WithValue(ctx, "user_id", claims["subject"])

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
