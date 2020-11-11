package models

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var session, _ = mgo.Dial("127.0.0.1")

type ToDoItem struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	Date        time.Time
	Description string
	Done        bool
}
