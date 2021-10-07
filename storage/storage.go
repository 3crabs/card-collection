package storage

import "card_collection/models"

//Storage хранилище данных
type Storage interface {
	AddCards(cards []models.Card) []models.Card
	GetAllCards() []models.Card
	GetCardById(id string) (models.Card, error)
}
