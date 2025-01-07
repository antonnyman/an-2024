package routes

import (
	"an-2024/views/blog"

	"github.com/labstack/echo/v4"
)

func AppStorageObservable(ctx echo.Context) error {
	return html(ctx, blog.AppStorageInObservable())
}
