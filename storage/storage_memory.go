package storage

import (
	"card_collection/models"
	"errors"
	"github.com/google/uuid"
)

//storageMemory хранилище данных в оперативной памяти
type storageMemory struct {
	cards map[string]models.Card
}

func newStorageMemory() *storageMemory {
	return &storageMemory{}
}

func (s *storageMemory) Init() {
	s.cards = make(map[string]models.Card)
}

func (s *storageMemory) AddCards(cards []models.Card) []models.Card {
	var tempCards []models.Card
	for _, c := range cards {
		id := uuid.New().String()
		c.Id = id
		s.cards[id] = c
		tempCards = append(tempCards, c)
	}
	return tempCards
}

func (s *storageMemory) GetAllCards() []models.Card {
	var tmp []models.Card
	for _, card := range s.cards {
		tmp = append(tmp, card)
	}
	return tmp
}

func (s *storageMemory) GetCardById(id string) (models.Card, error) {
	for _, card := range s.cards {
		if card.Id == id {
			return card, nil
		}
	}
	return models.Card{}, errors.New("card not found")
}
