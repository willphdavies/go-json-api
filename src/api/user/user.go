package user

import (
	"fmt"
	"net/http"
)

func Endpoint() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, "API / USER!!!")
	}
}
