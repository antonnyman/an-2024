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

	app.GET("", routes.IndexView)
	app.GET("/embroidery", routes.EmbroideryView)
	app.GET("/blog/my-coffee-recipes", routes.CoffeeRecipesView)
	app.GET("/blog/i-captured-60-dv-tapes-in-2020", routes.DVTapesView)

	app.HTTPErrorHandler = routes.HTTPErrorHandler

	port := os.Getenv("PORT")

	if len(port) == 0 {
		port = ":3456"
	}

	if err := app.Start(port); !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}
