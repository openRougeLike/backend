package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/openRougeLike/backend/routes"
)

func main() {

	// uMap := user.NewMap(0)

	// f, err := os.Create("map.json")

	// PanicIfErr(err)

	// b, err := json.Marshal(uMap)
	// PanicIfErr(err)

	// f.Write(b)
	// f.Close()

	r := chi.NewRouter()
	r.Mount("/game", routes.GameRouter())

	err := http.ListenAndServe(":3000", r)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Woooohooo!!! Started the server!!!")
	}
}
