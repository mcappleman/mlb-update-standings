package mongo

import (
	"gopkg.in/mgo.v2"
)

type DatabaseSession struct {
    *mgo.Session
    databaseName string
}

func NewSession(name string) *DatabaseSession {

    session, err := mgo.Dial("mongodb://localhost")
    if err != nil {
        panic(err)
    }

    return &DatabaseSession{session, name}

}

func (s *DatabaseSession) Database() *mgo.Database {

    return s.DB(s.databaseName)

}