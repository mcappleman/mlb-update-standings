package main

import (
	"log"
	"os"

	"github.com/mcappleman/mlb-update-standings/mongo"

	"gopkg.in/mgo.v2"
)

func main() {
	
	file, err := os.OpenFile("/home/ubuntu/log/mlb-update-standings.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Unable to open log file.")
	}

	defer file.Close()

	log.SetOutput(file)
	log.Println("Logging started")

	session := mongo.NewSession("mlb-feed")

}