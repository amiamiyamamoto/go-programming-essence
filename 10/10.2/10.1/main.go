package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type Todo struct {
	ID        int64
	Context   string
	Done      bool
	Until     time.Time
	CreatedAt time.Time
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "")
	})
	e.Logger.Fatal(e.Start(":8989"))
}
