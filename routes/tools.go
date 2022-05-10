package routes

import (
	"fmt"
	"net/http"
)

func Log(str string) {
	fmt.Println(str)
}

func LogIfErr(err error) {
	if err != nil {
		Log(err.Error())
	}
}

func WriteErr(err string, status int, rw http.ResponseWriter) {
	rw.WriteHeader(status)
	rw.Write([]byte(fmt.Sprintf(`{"err":"%v"}`, err)))
}

const (
	ERR_JSON = "Unknown marshal error"
)