package routes

import (
	"an-2024/views/coffee_recipes"

	"github.com/labstack/echo/v4"
)

func CoffeeRecipesView(ctx echo.Context) error {
	return html(ctx, coffee_recipes.CoffeeRecipesView())
}
