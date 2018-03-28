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

	dir, err := os.Getwd()
	if err != nil {
		log.Println("Error getting the current directory in the DecodeConfig function")
		log.Fatalln(err)
	}
	file, _ := os.Open(dir + "/config.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Config{}
	err = decoder.Decode(&configuration)
	if err != nil {
		log.Println("ERROR")
		log.Fatalln(err)
	}

	return configuration

}
