package main

import (
	"log"
	"os"

	"github.com/mcappleman/mlb-update-standings/config"
	"github.com/mcappleman/mlb-update-standings/mongodb"
	"github.com/mcappleman/mlb-update-standings/teams"
)

func main() {

	conf := config.DecodeConfig()
	
	file, err := os.OpenFile(conf.LOG_FILE, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Unable to open log file.")
	}

	defer file.Close()

	log.SetOutput(file)
	log.Println("Logging started")

	session := mongodb.NewSession(conf.DATABASE_URL, conf.DATABASE_NAME)

	teamList := teams.GetTeams(session.Database())

	log.Println(teamList)

}