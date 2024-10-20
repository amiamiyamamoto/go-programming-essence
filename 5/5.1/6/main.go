package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gocraft/web"
)

type AppContext struct {
	HelloCount int
}

func (c *AppContext) SetHelloCount(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	if c.HelloCount == 0 {
		c.HelloCount = 3
	}
	next(rw, req)
}

func (c *AppContext) SayHello(rw web.ResponseWriter, req *web.Request) {
	// "Hello "をカウント分だけ出力
	fmt.Fprint(rw, strings.Repeat("Hello ", c.HelloCount), "World!")
}

func main() {
	// 構造体の値型を渡す
	router := web.New(AppContext{}).
		Middleware(web.LoggerMiddleware).
		Middleware(web.ShowErrorsMiddleware).
		Middleware((*AppContext).SetHelloCount).
		Get("/", (*AppContext).SayHello)
	http.ListenAndServe("localhost:3000", router)
}
