package business

import (
	"dreamon/dbBase"

	"dreamon/models/mongo"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type WishService struct{}

func initCollection() (*mgo.Session, *mgo.Collection) {
	conn := dbBase.GlobalSession.Clone()
	col := conn.DB("dreamon").C("wish")
	return conn, col
}

func (WishService) GetWish(id string) (result *mongo.Wish) {
	conn, col := initCollection()
	defer conn.Close()
	col.FindId(bson.ObjectIdHex(id)).One(&result)
	return
}

func (WishService) AddWish(title string, wish string) (result *mongo.Wish) {
	conn, col := initCollection()
	defer conn.Close()
	model := &mongo.Wish{
		Title:      title,
		Wish:       wish,
		CreateTime: time.Now().UTC(),
		UpdateTime: time.Now().UTC(),
		IsDeleted:  false,
	}
	err := col.Insert(model)
	if err != nil {
		panic(err)
	} else {
		result = model
	}
	return
}
