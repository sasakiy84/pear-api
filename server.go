package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Article struct {
	Title    string   `json:"title"`
	Content  string   `json:"content"`
	Url      string   `json:"url"`
	Keywords []string `json:"keywords"`
}

type ApiResponse struct {
	Ok   bool        `json:"ok"`
	Data interface{} `json:"data"`
}

func main() {

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/article", func(c echo.Context) error {
		a := new(Article)
		if err := c.Bind(a); err != nil {
			return c.JSON(http.StatusBadRequest, ApiResponse{false, err})
		}
		return c.JSON(http.StatusOK, ApiResponse{true, a})
	})
	e.Logger.Fatal(e.Start(":1323"))
}
