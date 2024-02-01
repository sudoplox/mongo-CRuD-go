package models

import "gopkg.in/mgo.v2/bson"

type User struct {
	//  json:"id"	-> get in request
	//  bson:"_id"	-> would go in db
	Id     bson.ObjectId `json:"id" bson:"_id"`
	Name   string        `json:"name" bson:"name"`
	Gender string        `json:"gender" bson:"gender"`
	Age    int           `json:"age" bson:"age"`
}
