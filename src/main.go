package main

import (
	"fmt"
	"go-json-api/src/api"
	"log"
	"net/http"
)

func welcome() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Holy McShizzle!")
	}
}

func main() {
	http.HandleFunc("/", welcome())
	http.HandleFunc("/api", api.Endpoint())

	log.Fatal(http.ListenAndServe(":8080", nil))
}
