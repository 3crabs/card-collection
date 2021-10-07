package main

import (
	"card_collection/models"
	"card_collection/storage"
	"github.com/labstack/echo/v4"
	"net/http"
)

//AppInfo информация о пиложении
type AppInfo struct {
	AppName    string `json:"app_name"`
	AppVersion string `json:"app_version"`
}

func main() {
	e := echo.New()

	// потом поменять на интерфейс и выбирать в зависимости от среды реализации
	store := storage.NewStorageMemory()

	// обязательный роут с информацией о приложении
	e.GET("/info", func(c echo.Context) error {
		return c.JSON(http.StatusOK, AppInfo{
			AppName:    "card-collection",
			AppVersion: "0.1.0",
		})
	})

	// создание карт
	e.POST("/cards", func(c echo.Context) error {
		var cards []models.Card
		if err := c.Bind(&cards); err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		tempCards := store.AddCards(cards)
		return c.JSON(http.StatusOK, tempCards)
	})

	// получаем все карты
	e.GET("/cards", func(c echo.Context) error {
		return c.JSON(http.StatusOK, store.GetAllCards())
	})

	// получаем карту по uuid
	e.GET("/cards/:id", func(c echo.Context) error {
		id := c.Param("id")
		card, err := store.GetCardById(id)
		if err != nil {
			return c.JSON(http.StatusNotFound, err)
		}
		return c.JSON(http.StatusOK, card)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
