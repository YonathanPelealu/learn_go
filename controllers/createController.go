package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"gopkg.in/mgo.v2/bson"
)

var c = session.DB("shopifyClone").C("Shopify")

func AddToDo(w http.ResponseWriter, r *http.Request) {
	_ = c.Insert(models.ToDoItem{
		bson.NewObjectId(),
		time.Now(),
		r.FormValue("description"),
		false,
	})

	result := models.ToDoItem{}
	_ = c.Find(bson.M{"description": r.FormValue("description")}).One(&result)
	json.NewEncoder(w).Encode(result)
}
