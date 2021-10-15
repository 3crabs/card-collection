package config

import (
	"fmt"
	"github.com/tkanos/gonfig"
	"log"
)

type Configuration struct {
	DbUsername string
	DbPassword string
	DbPort     string
	DbHost     string
	DbName     string
}

func GetConfig(params ...string) Configuration {
	configuration := Configuration{}
	env := "dev"
	if len(params) > 0 {
		env = params[0]
	}
	fileName := fmt.Sprintf("./%s_config.json", env)
	err := gonfig.GetConf(fileName, &configuration)
	if err != nil {
		log.Panicln(err)
	}
	return configuration
}

func GetPostgresConnectionString() string {
	config := GetConfig("dev")
	dataBase := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		config.DbHost,
		config.DbPort,
		config.DbUsername,
		config.DbName,
		config.DbPassword)
	return dataBase
}
