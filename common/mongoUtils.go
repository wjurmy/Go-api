package common

import (
	"gopkg.in/mgo.v2"
	"time"
	"log"
)

var session *mgo.Session

func GetSession() *mgo.Session  {
	if session == nil{
		var err error

		session, err = mgo.DialWithInfo(&mgo.DialInfo{
			Addrs: []string{AppConfig.SqlDatabaseHost},
			Username:	AppConfig.DatabaseUser,
			Password:AppConfig.DatabasePassword,
			Timeout:60*time.Second,


		})
		if err != nil {
			log.Fatalf("[createDbSession]: %s\n ", err)
		}
	}
	return session
}
func createDbSession(){
	var err error
	session, err = mgo.DialWithInfo(&mgo.DialInfo{
		Addrs: []string{AppConfig.SqlDatabaseHost},
		Username:AppConfig.DatabaseUser,
		Password:AppConfig.DatabasePassword,
		Timeout: 60*time.Second,
	})
	if err!=nil {
		log.Fatalf("[createDbSession]:  %s\n", err)
	}
}