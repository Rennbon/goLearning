package mongo

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Wish struct {
	Id    bson.ObjectId `bson:"_id,omitempty"`
	Title string        `bson:"title"`
	Wish  string        `bson:"wish"`
	//Sort  int           `bson:"sort"`

	CreateTime time.Time `bson:"ct"`
	UpdateTime time.Time `bson:"ut"`
	IsDeleted  bool      `bson:"isd"`
}
