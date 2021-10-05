package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

//AppInfo информация о пиложении
type AppInfo struct {
	AppName    string `json:"app_name"`
	AppVersion string `json:"app_version"`
}

//Card карта - единица коллекционирования
type Card struct {
	Name string `json:"name"`
}

func main() {
	e := echo.New()

	// обязательный роут с информацией о приложении
	e.GET("/info", func(c echo.Context) error {
		return c.JSON(http.StatusOK, AppInfo{
			AppName:    "card-collection",
			AppVersion: "0.1.0",
		})
	})

	// создание карт
	e.POST("/cards", func(c echo.Context) error {
		var cards []Card
		if err := c.Bind(&cards); err != nil {
			return err
		}
		return c.JSON(http.StatusOK, cards)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
