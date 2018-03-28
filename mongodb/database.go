package mongodb

import (
	"log"
	"gopkg.in/mgo.v2"
)

type DatabaseSession struct {
    *mgo.Session
    databaseName string
}

func NewSession(url string, name string) *DatabaseSession {

	log.Println(url)
	log.Println(name)

	url = url + "/" + name
    session, err := mgo.Dial(url)
    if err != nil {
    	log.Println(err)
        panic(err)
    }

    return &DatabaseSession{session, name}

}

func (s *DatabaseSession) Database() *mgo.Database {

    return s.DB(s.databaseName)

}
