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

type server struct {
	store  storage.Storage
	router *echo.Echo
}

func newServer() *server {
	s := &server{
		store:  storage.NewStorageMemory(),
		router: echo.New(),
	}
	s.routes()
	return s
}

func (s *server) routes() {
	s.router.GET("/info", s.handleInfo())
	s.router.POST("/cards", s.handleAddCards())
	s.router.GET("/cards", s.handleGetAllCards())
	s.router.GET("/cards/:id", s.handleGetCardById())
}

// обязательный роут с информацией о приложении
func (s *server) handleInfo() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, AppInfo{
			AppName:    "card-collection",
			AppVersion: "0.2.0",
		})
	}
}

// создание карт
func (s *server) handleAddCards() echo.HandlerFunc {
	return func(c echo.Context) error {
		var cards []models.Card
		if err := c.Bind(&cards); err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		tempCards := s.store.AddCards(cards)
		return c.JSON(http.StatusOK, tempCards)
	}
}

// получаем все карты
func (s *server) handleGetAllCards() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, s.store.GetAllCards())
	}
}

// получаем карту по uuid
func (s *server) handleGetCardById() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		card, err := s.store.GetCardById(id)
		if err != nil {
			return c.JSON(http.StatusNotFound, err)
		}
		return c.JSON(http.StatusOK, card)
	}
}

func main() {
	s := newServer()
	s.router.Logger.Fatal(s.router.Start(":8080"))
}
