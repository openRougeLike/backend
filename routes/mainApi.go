package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// /game routes
func GameRouter() http.Handler {
	r := chi.NewRouter()

	// Middleware that can cause errors & such
	r.Use(AuthMiddleware)

	// r.Use(middleware.Compress(5))
	r.Use(middleware.CleanPath)
	r.Use(middleware.GetHead)
	r.Use(middleware.AllowContentType("application/json"))
	r.Use(middleware.NoCache)

	// Last middleware!
	r.Use(SuccessMiddleware)

	r.Mount("/self", userRouter())
	r.Mount("/fight", fightRouter())
	r.Mount("/map", mapRouter())

	return r
}
