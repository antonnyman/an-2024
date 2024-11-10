package routes

import (
	"an-2024/views/dv_tapes"

	"github.com/labstack/echo/v4"
)

func DVTapesView(ctx echo.Context) error {
	return html(ctx, dv_tapes.DVTapesView())
}
