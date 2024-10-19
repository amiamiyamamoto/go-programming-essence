package main

import (
	"fmt"
	"net/http"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, my handler!")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello,World!")
	})

	http.Handle("/ami", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Ami!")
	}))

	http.HandleFunc("/myhandler", myHandler)
	http.ListenAndServe(":8080", nil)
}
