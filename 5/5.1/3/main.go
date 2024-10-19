package main

import (
	"database/sql"
	"net/http"
)

type MyContext struct {
	db *sql.DB
}

func NewMyContext() *MyContext {
	return &MyContext{}
}

func (m *MyContext) handle(w http.ResponseWriter, r *http.Request) {
	// m.dbを使った処理
}

func main() {
	myctx := NewMyContext()
	http.HandleFunc("/", myctx.handle)

	http.ListenAndServe(":8080", nil)

}
