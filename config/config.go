package config

import (
	"encoding/json"
	"os"
	"log"
)

type Config struct {
	LOG_FILE string
	DATABASE_URL string
	DATABASE_NAME string
}

func DecodeConfig() Config{

	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	configuration := Config{}
	err := decoder.Decode(&configuration)
	if err != nil {
		log.Fatalln(err)
	}

	return configuration

}