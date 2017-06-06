package teams

import (
	"time"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Team struct {
	Id		 bson.ObjectId  `json:"id" bson:"_id"`
	Name     string 		`json:"name" bson:"name"`
	City     string 		`json:"city" bson:"city"`
	Abbrev   string         `json:"abbrev" bson:"abbrev"`
}

type Game struct {
	Id		     bson.ObjectId      `json:"id" bson:"_id"`
	Date         time.Time          `json:"date" bson:"date"`
	HomeTeam     bson.ObjectId 		`json:"home_team" bson:"home_team"`
	AwayTeam     bson.ObjectId 		`json:"away_team" bson:"away_team"`
	HomeRuns     int                `json:"home_runs" bson:"home_runs"`
	AwayRuns     int                `json:"away_runs" bson:"away_runs"`
	Status       string             `json:"status" bson:"status"`
}

type Record struct {
	Id      bson.ObjectId    `json:"id" bson:"_id"`
	Team    bson.ObjectId    `json:"team" bson:"team"`
	Wins	int              `json:"wins" bson:"wins"`
	Losses	int              `json:"losses" bson:"losses"`
	Year	int              `json:"year" bson"year"`
}

const collection = "teams"

func GetTeams(db *mgo.Database) []Team {

	teamList := []Team{}
	err := db.C(collection).Find(nil).All(&teamList)
	if err != nil {
		panic(err);
	}

	return teamList

}