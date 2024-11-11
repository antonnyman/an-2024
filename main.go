package main

import (
	"an-2024/lib"
	"an-2024/routes"
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()
	// app.Use(middleware.Gzip())
	lib.CombineCSS("./assets/stylesheets/partials", "./assets/stylesheets/styles.css")

	app.Static("/assets", "assets")

	app.GET("/healthcheck", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	app.GET("", routes.IndexView)
	app.GET("/embroidery", routes.EmbroideryView)
	app.GET("/blog/my-coffee-recipes", routes.CoffeeRecipesView)
	app.GET("/blog/i-captured-60-dv-tapes-in-2020", routes.DVTapesView)

	app.HTTPErrorHandler = routes.HTTPErrorHandler

	env := os.Getenv("APP_ENV")
	port := ":3456"

	if env == "production" {
		port = ":5000"
	}

	if err := app.Start(port); !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}
