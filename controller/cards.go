package controller

import (
	"card_collection/models"
	"card_collection/storage"
	"net/http"

	"github.com/labstack/echo/v4"
)

// HandleInfo обязательный роут с информацией о приложении
func HandleInfo(c echo.Context) error {
	return c.JSON(http.StatusOK, models.AppInfo{
		AppName:    "card-collection",
		AppVersion: "0.2.0",
	})
}

// AddCards создание карт
func AddCards(c echo.Context) error {
	var cards []models.Card
	if err := c.Bind(&cards); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	tempCards := storage.AddCards(cards)
	return c.JSON(http.StatusOK, tempCards)
}

// GetAllCards получаем все карты
func GetAllCards(c echo.Context) error {
	return c.JSON(http.StatusOK, storage.GetAllCards())
}

// GetCardById получаем карту по uuid
func GetCardById(c echo.Context) error {
	id := c.Param("id")
	card, err := storage.GetCardById(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusOK, card)
}
