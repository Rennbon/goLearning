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
		Addrs:     []string{"0.0.0.0:32768"},
		Timeout:   60 * time.Second,
		PoolLimit: 200,
	}
	session, err := mgo.DialWithInfo(mongoDBInfo)
	if err != nil {
		panic(err)
	}
	GlobalSession = session
	return nil
}
