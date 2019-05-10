package model

import (
	"github.com/globalsign/mgo/bson"
)

//User ...
type User struct {
	ID           bson.ObjectId `bson:"id" json:"id"`
	NickName     string        `bson:"nick_name" json:"nick_name"`
	Phone        string        `bson:"phone" json:"phone"`
	PasswordHash []byte        `bson:"hash" json:"hash"`
	Active       bool          `bson:"active" json:"active"`
}
