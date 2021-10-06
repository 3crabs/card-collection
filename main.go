package main

import (
	"github.com/google/uuid"
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
	Id   string `json:"id"`
	Name string `json:"name"`
}

func main() {
	e := echo.New()

	m := make(map[string]Card)

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
		tempCards := make(map[string]Card)

		for _, c := range cards {
			id := uuid.New().String()
			c.Id = id
			m[id] = c
			tempCards[id] = c
		}
		return c.JSON(http.StatusOK, tempCards)
	})

	// получаем все карты
	e.GET("/cards", func(c echo.Context) error {
		return c.JSON(http.StatusOK, m)
	})

	// получаем карту по uuid
	e.GET("/cards/:id", func(c echo.Context) error {
		id := c.Param("id")
		card, ok := m[id]
		if !ok {
			return c.JSON(http.StatusNotFound, "card not found")
		}
		return c.JSON(http.StatusOK, card)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
