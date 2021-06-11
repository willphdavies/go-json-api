package api

import (
	"fmt"
	"go-json-api/src/api/user"
	"net/http"
)

func Endpoint() http.HandlerFunc {
	http.HandleFunc("/api/user", user.Endpoint())
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "API!!!")
	}
}
