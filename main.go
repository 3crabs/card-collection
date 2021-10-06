package main

import (
	"errors"
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

//Storage хранилище данных
type Storage interface {
	AddCards(cards []Card) []Card
	GetAllCards() []Card
	GetCardById(id string) (Card, error)
}

//StorageMemory хранилище данных в оперативной памяти
type StorageMemory struct {
	cards map[string]Card
}

func NewStorageMemory() *StorageMemory {
	return &StorageMemory{cards: make(map[string]Card)}
}

func (s StorageMemory) AddCards(cards []Card) []Card {
	var tempCards []Card
	for _, c := range cards {
		id := uuid.New().String()
		c.Id = id
		s.cards[id] = c
		tempCards = append(tempCards, c)
	}
	return tempCards
}

func (s StorageMemory) GetAllCards() []Card {
	var tmp []Card
	for _, card := range s.cards {
		tmp = append(tmp, card)
	}
	return tmp
}

func (s StorageMemory) GetCardById(id string) (Card, error) {
	for _, card := range s.cards {
		if card.Id == id {
			return card, nil
		}
	}
	return Card{}, errors.New("card not found")
}

func main() {
	e := echo.New()

	// потом поменять на интерфейс и выбирать в зависимости от среды реализации
	storage := NewStorageMemory()

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
			return c.JSON(http.StatusInternalServerError, err)
		}
		tempCards := storage.AddCards(cards)
		return c.JSON(http.StatusOK, tempCards)
	})

	// получаем все карты
	e.GET("/cards", func(c echo.Context) error {
		return c.JSON(http.StatusOK, storage.GetAllCards())
	})

	// получаем карту по uuid
	e.GET("/cards/:id", func(c echo.Context) error {
		id := c.Param("id")
		card, err := storage.GetCardById(id)
		if err != nil {
			return c.JSON(http.StatusNotFound, err)
		}
		return c.JSON(http.StatusOK, card)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
