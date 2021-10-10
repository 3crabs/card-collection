package route

import (
	"card_collection/controller"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	//e.GET("/", controller.Home)

	//e.GET("/users", controller.GetUsers)

	e.GET("/info", controller.HandleInfo)
	e.POST("/cards", controller.AddCards)
	e.GET("/cards", controller.GetAllCards)
	e.GET("/cards/:id", controller.GetCardById)
	return e
}
