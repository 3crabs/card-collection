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

var DB *gorm.DB

func Init() {
	var err error
	conString := config.GetPostgresConnectionString()

	log.Print(conString)

	DB, err = gorm.Open("postgres", conString)

	if err != nil {
		log.Panic(err)
		panic("DB Connection Error")
	}
	DB.AutoMigrate(&models.Card{})
}

func GetDBInstance() *gorm.DB {
	return DB
}

func AddCards(cards []models.Card) []models.Card {
	var tempCards []models.Card
	for _, c := range cards {
		id := uuid.New().String()
		c.Id = id
		DB.Create(&c)
		tempCards = append(tempCards, c)
	}
	return tempCards
}

func GetAllCards() []models.Card {
	var cards []models.Card
	_ = DB.Find(&cards)
	return cards
}

func GetCardById(id string) (models.Card, error) {
	card := models.Card{}
	_ = DB.Find(&card, "id = ?", id)
	if result := DB.First("id = ?", id).First(&card); result.Error != nil {
		return models.Card{}, errors.New("card not found")
	}
	return card, nil
}
