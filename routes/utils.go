package routes

import (
	"an-2024/views/not_found"
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func html(ctx echo.Context, component templ.Component) error {
	return component.Render(ctx.Request().Context(), ctx.Response())
}

func HTTPErrorHandler(err error, ctx echo.Context) {
	if he, ok := err.(*echo.HTTPError); ok {
		switch he.Code {
		case http.StatusNotFound:
			// Custom 404 page
			if ctx.Path() == "/api/*" {
				ctx.JSON(http.StatusNotFound, map[string]string{
					"error":   "404",
					"message": "Resource not found",
				})
			} else {
				html(ctx, not_found.FourOhFourView())
			}
			return
		case http.StatusInternalServerError:
			html(ctx, not_found.FiveHundredView())
			return
		}
	}
	log.Printf("Unexpected error: %v", err)
	ctx.HTML(http.StatusInternalServerError, "<h1>500 Internal Server Error</h1><p>Unexpected error occurred.</p>")
}
