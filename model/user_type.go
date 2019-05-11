package model

import (
	"github.com/globalsign/mgo/bson"
)

//UserType ...
type UserType struct {
	ID   bson.ObjectId `bson:"id,omitempty" json:"id"`
	Type string        `bson:"type,omitempty" json:"type"`
}
