package middleware

import (
	"context"
	"net/http"

	"github.com/rajaabluu/ershop-api/internal/helper"
	"github.com/rajaabluu/ershop-api/internal/model"
)

func (middleware *Middleware) VerifyAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		SCHEMA := "Bearer "
		authHeader := r.Header.Get("Authorization")
		if len(authHeader) == 0 {
			helper.WriteJSONResponse(w, &model.ErrResponse{Message: "Invalid token, unauthorized user"}, http.StatusUnauthorized)
			return
		}
		tokenString := authHeader[len(SCHEMA):]
		auth, err := middleware.UserService.Verify(tokenString)
		if err != nil {
			helper.WriteJSONResponse(w, &model.ErrResponse{Message: "Unauthorized user", Errors: err.Error()}, http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), model.AuthContextKey, auth)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (middleware *Middleware) VerifyIsAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth, err := middleware.UserService.GetCurrentAuth(r.Context())
		if err != nil {
			helper.WriteJSONResponse(w, &model.ErrResponse{Message: err.Error()}, http.StatusInternalServerError)
			return
		}
		if auth.Role != 1 {
			helper.WriteJSONResponse(w, &model.ErrResponse{Message: "you're not permitted to do this operation", Errors: "unauthorized users"}, http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
