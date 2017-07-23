package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	LOG_FILE      string
	DATABASE_URL  string
	DATABASE_NAME string
}

func DecodeConfig() Config {

	file, _ := os.Open("/home/matt_cappleman/projects/go/src/github.com/mcappleman/mlb-update-standings/config.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Config{}
	err := decoder.Decode(&configuration)
	if err != nil {
		log.Println("ERROR")
		log.Fatalln(err)
	}

	return configuration

}
