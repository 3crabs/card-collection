package route

import (
	"card_collection/controller"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/info", controller.HandleInfo)
	e.POST("/cards", controller.AddCards)
	e.GET("/cards", controller.GetAllCards)
	e.GET("/cards/:id", controller.GetCardById)

	return e
}
