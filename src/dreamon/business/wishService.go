package business

import (
	"dreamon/dbBase"
	"dreamon/models/mongo"
	"time"

	mgo "gopkg.in/mgo.v2"
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
	if bson.IsObjectIdHex(id) {
		col.FindId(bson.ObjectIdHex(id)).One(&result)
	}
	return
}

func (WishService) AddWish(title string, wish string) (result *mongo.Wish) {
	conn, col := initCollection()
	defer conn.Close()
	model := &mongo.Wish{
		Id:         bson.NewObjectId(),
		Title:      title,
		Wish:       wish,
		CreateTime: time.Now().UTC(),
		UpdateTime: time.Now().UTC(),
		IsDeleted:  false,
	}
	err := col.Insert(model)
	if err != nil {
		panic(err)
	}
	result = model
	//以下为用upsert
	//err := col.Insert(model)
	// info, err := col.Upsert(nil, model)
	// if err != nil {
	// 	panic(err)
	// } else {
	// 	model.Id = info.UpsertedId.(bson.ObjectId)
	// 	result = model
	// }
	return
}
