package mongodb

import (
	"gopkg.in/mgo.v2"
)

type DatabaseSession struct {
    *mgo.Session
    databaseName string
}

func NewSession(url string, name string) *DatabaseSession {

    session, err := mgo.Dial(url)
    if err != nil {
        panic(err)
    }

    return &DatabaseSession{session, name}

}

func (s *DatabaseSession) Database() *mgo.Database {

    return s.DB(s.databaseName)

}