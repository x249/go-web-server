package model

import (
	"github.com/globalsign/mgo/bson"
)

//MovieCategory ...
type MovieCategory struct {
	ID   bson.ObjectId `bson:"id,omitempty" json:"id"`
	Name string        `bson:"name,omitempty" json:"name"`
}
