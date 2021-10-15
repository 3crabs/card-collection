package config

import (
	"fmt"
	"github.com/tkanos/gonfig"
	"log"
)

type config struct {
	Username string
	Password string
	Port     string
	Host     string
	Name     string
}

func getConfig(params ...string) config {
	cfg := config{}
	env := "dev"
	if len(params) > 0 {
		env = params[0]
	}
	fileName := fmt.Sprintf("./%s_config.json", env)
	err := gonfig.GetConf(fileName, &cfg)
	if err != nil {
		log.Panicln(err)
	}
	return cfg
}

func GetPostgresConnectionString() string {
	config := getConfig("dev")
	dataBase := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		config.Host,
		config.Port,
		config.Username,
		config.Name,
		config.Password)
	return dataBase
}
