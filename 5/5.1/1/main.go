package main

import (
	"fmt"
	"net/http"
)

func main() {
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello, World!")
	// })
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			//GETの処理
			fmt.Fprintf(w, "Hello, World!")
		default:
		}
	})
	http.ListenAndServe(":8080", nil)
}
