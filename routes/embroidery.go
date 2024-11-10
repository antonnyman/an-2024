package routes

import (
	"an-2024/views/embroidery"

	"github.com/labstack/echo/v4"
)

func EmbroideryView(ctx echo.Context) error {
	return html(ctx, embroidery.EmbroideryView())
}
