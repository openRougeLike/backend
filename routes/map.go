package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/openRougeLike/backend/database"
	"github.com/openRougeLike/backend/user"
)

func mapRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/", func(rw http.ResponseWriter, r *http.Request) {
		user := (r.Context().Value(USER)).(user.User)
		usrJson, _ := json.Marshal(user.Map)
		rw.Write(usrJson)
	})

	r.Post("/move/{dir:0|1|2|3}/{num:\\d+}", routeMove)
	r.Post("/move/{dir:0|1|2|3}", routeMove)

	r.Post("/action/{dir:0|1|2|3|4}", func(w http.ResponseWriter, r *http.Request) {
		usr := &database.GUser
		directionRaw := chi.URLParam(r, "dir")
		dirP, _ := strconv.ParseInt(directionRaw, 10, 64)
		dir := user.Direction(dirP)

		tile := usr.Map.Layout[usr.Map.User[0]+user.DirectionEnum[dir][0]][usr.Map.User[1]+user.DirectionEnum[dir][1]]

		switch tile {
		case user.MapArena:
			WriteErr("TODO:", 501, w)
			return
		case user.MapChest:
			WriteErr("TODO:", 501, w)
			return
		case user.MapExit:
			usr.Exit()
		default:
			WriteErr(user.ErrNoAction.Error(), 400, w)
			return
		}

		jsonUsr, _ := json.Marshal(usr)

		w.WriteHeader(200)
		w.Write(jsonUsr)
	})

	return r
}

func routeMove(rw http.ResponseWriter, r *http.Request) {
	directionRaw := chi.URLParam(r, "dir")
	dirP, _ := strconv.ParseInt(directionRaw, 10, 64)
	dir := user.Direction(dirP)
	amountRaw := chi.URLParam(r, "num")
	amount := 0
	if amountRaw != "" {
		amtP, _ := strconv.ParseInt(amountRaw, 10, 64)

		amount = int(amtP)
	}

	mErr := database.GUser.Move(dir, amount)

	if mErr != nil {
		WriteErr(mErr.Error(), 400, rw)
		return
	}

	js, _ := json.Marshal(database.GUser.Map)

	rw.WriteHeader(200)
	rw.Write(js)
}
