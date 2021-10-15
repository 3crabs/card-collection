package storage

import (
	"card_collection/config"
	"card_collection/models"
	"errors"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

//storageDb хранилище данных в postgres
type storageDb struct {
	Db *gorm.DB
}

func newStorageDb() *storageDb {
	return &storageDb{}
}

func (s *storageDb) Init() {
	var err error
	conString := config.GetPostgresConnectionString()

	s.Db, err = gorm.Open("postgres", conString)

	if err != nil {
		log.Panic(err)
	}
	s.Db.AutoMigrate(&models.Card{})
}

func (s storageDb) AddCards(cards []models.Card) []models.Card {
	var tempCards []models.Card
	for _, c := range cards {
		id := uuid.New().String()
		c.Id = id
		s.Db.Create(&c)
		tempCards = append(tempCards, c)
	}
	return tempCards
}

func (s storageDb) GetAllCards() []models.Card {
	var cards []models.Card
	_ = s.Db.Find(&cards)
	return cards
}

func (s storageDb) GetCardById(id string) (models.Card, error) {
	card := models.Card{}
	_ = s.Db.Find(&card, "id = ?", id)
	if result := s.Db.First("id = ?", id).First(&card); result.Error != nil {
		return models.Card{}, errors.New("card not found")
	}
	return card, nil
}
