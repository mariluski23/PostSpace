package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, world!")
	})
	http.ListenAndServe(":8080", nil)
}