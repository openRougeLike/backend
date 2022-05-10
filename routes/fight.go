package routes

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/openRougeLike/backend/user"
)

func fightRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/", func(rw http.ResponseWriter, r *http.Request) {
		user := (r.Context().Value(USER)).(user.User)
		usrJson, err := json.Marshal(user)
		if err != nil {
			rw.Write([]byte(`{"err": "Unknown unmarshaling error! (Id: "}`))
			// TODO: Log errors or smt
			rw.WriteHeader(501)
			return
		}

		rw.Write(usrJson)
	})

	r.Get("/attack", func(rw http.ResponseWriter, r *http.Request) {
		user := (r.Context().Value(USER)).(user.User)
		stats := user.FetchStats()

		usrJson, err := json.Marshal(stats)

		if err != nil {
			rw.Write([]byte(`{"err": "Unknown unmarshaling error! (Id: "}`))
			// TODO: Log errors or smt
			rw.WriteHeader(501)
			return
		}

		rw.Write(usrJson)
	})

	return r
}
