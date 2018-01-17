package dbBase

import (
	"time"

	"gopkg.in/mgo.v2"
)

var (
	GlobalSession *mgo.Session
)

func init() {
	loadGlobalSession()
}

func loadGlobalSession() error {

	mongoDBInfo := &mgo.DialInfo{
		Addrs:     []string{C.MongoDB.Addr},
		Timeout:   C.MongoDB.Timeout * time.Second,
		PoolLimit: C.MongoDB.PoolLimit,
	}
	session, err := mgo.DialWithInfo(mongoDBInfo)
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	GlobalSession = session
	return nil
}
