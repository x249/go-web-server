package model

import (
	"github.com/globalsign/mgo/bson"
)

//User ...
type User struct {
	ID           bson.ObjectId `bson:"id,omitempty" json:"id"`
	Type         bson.ObjectId `bson:"type,omitempty" json:"type"`
	NickName     string        `bson:"nick_name,omitempty" json:"nick_name"`
	Phone        string        `bson:"phone,omitempty" json:"phone"`
	PasswordHash string        `bson:"hash,omitempty" json:"hash"`
	Active       bool          `bson:"active" json:"active"`
}
