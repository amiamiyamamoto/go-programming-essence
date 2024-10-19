package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
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
	http.HandleFunc("/ami", func(w http.ResponseWriter, r *http.Request) {
		f, err := os.Open("content.txt")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()
		io.Copy(w, f)
	})
	http.ListenAndServe(":8080", nil)
}
