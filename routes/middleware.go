package routes

import (
	"context"
	"net/http"

	"github.com/openRougeLike/backend/database"
	"github.com/openRougeLike/backend/user"
)

type CtxKeys int8

const (
	USER CtxKeys = iota
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("authorization")
		if len(auth) == 0 {
			w.WriteHeader(401)
			w.Write([]byte("No authorization header provided"))
		} else {
			userFetched, ok := database.FetchUser(auth)
			if !ok {
				w.WriteHeader(401)
				w.Write([]byte("Not Authorized"))
			} else {
				ctx := context.WithValue(r.Context(), USER, userFetched)
				next.ServeHTTP(w, r.WithContext(ctx))
			}
		}
	})
}

func SuccessMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		next.ServeHTTP(w, r)
	})
}

func FightOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		usr := (r.Context().Value(USER)).(user.User)
		if usr.State != user.STATE_FIGHT {
			// log out or smt idk
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
