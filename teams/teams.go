package teams

import (
	"log"
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
	Id      	bson.ObjectId    	`json:"id" bson:"_id"`
	Team    	bson.ObjectId    	`json:"team" bson:"team"`
	Wins		int              	`json:"wins" bson:"wins"`
	Losses		int              	`json:"losses" bson:"losses"`
	Year		int              	`json:"year" bson:"year"`
	WinPercent	float64			`json:"win_percent" bson:"win_percent"`
	EloRating	float64			`json:"elo_rating" bson:"elo_rating"`
}

const collection = "teams"
const games = "games"
const records = "records"

func GetTeams(db *mgo.Database) []Team {

	teamList := []Team{}
	err := db.C(collection).Find(nil).All(&teamList)
	if err != nil {
		panic(err);
	}

	return teamList

}

func (t *Team) GetAndUpdateRecord(db *mgo.Database) {

	teamSchedule := []Game{}
	currentYear := time.Now().Year()
	fromDate := time.Date(currentYear, time.January, 1, 0, 0, 0, 0, time.UTC)
	query := bson.M{ "status": "Final", "date": bson.M{ "$gte": fromDate }, "$or": []bson.M{ bson.M{"home_team": t.Id}, bson.M{"away_team": t.Id} } }
	err := db.C(games).Find( query ).All(&teamSchedule)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	var wins int = 0
	var losses int = 0

	for _, game := range teamSchedule {

		if t.Id == game.HomeTeam {

			if game.HomeRuns > game.AwayRuns {
				wins += 1
			} else {
				losses += 1
			}

		} else {

			if game.AwayRuns > game.HomeRuns {
				wins += 1
			} else {
				losses += 1
			}

		}

	}

	winPercent := float64(wins)/(float64(wins) + float64(losses))
	var record Record
	err = db.C(records).Find(bson.M{ "team": t.Id, "year": currentYear }).One(&record)
	if err != nil {
		record = Record{bson.NewObjectId(), t.Id, wins, losses, currentYear, winPercent, 0}
		err := db.C(records).Insert(record)
		if err != nil {
			log.Println("Error inserting Team Record", err)
			panic(err)
		}
	}

	record.Wins = wins
	record.Losses = losses
	record.WinPercent = winPercent
	err = db.C(records).Update(bson.M{ "_id": record.Id }, record)
	if err != nil {
		log.Println("Error updating Team Record", record.Team, err)
		panic(err)
	}

}
