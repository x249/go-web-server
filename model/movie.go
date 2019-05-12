package model

import (
	"github.com/globalsign/mgo/bson"
)

//MovieCategory ...
type MovieCategory struct {
	ID   bson.ObjectId `bson:"id,omitempty" json:"id"`
	Name string        `bson:"name,omitempty" json:"name"`
}

//Movie ...
type Movie struct {
	ID          bson.ObjectId `bson:"id" json:"id"`
	Category    bson.ObjectId `bson:"cate" json:"cate"`
	Theater     bson.ObjectId `bson:"theater" json:"theater"`
	Name        string        `bson:"name" json:"name"`
	Poster      string        `bson:"poster" json:"poster"`
	Thmubnail   string        `bson:"thumb" json:"thumb"`
	Description string        `bson:"desc" json:"desc"`
}

//MovieRate ...
type MovieRate struct {
	Movie bson.ObjectId `bson:"movie" json:"movie"`
	User  bson.ObjectId `bson:"user" json:"user"`
	Rate  float32       `bson:"rate" json:"rate"`
}

//MovieComments ...
type MovieComments struct {
	Movie   bson.ObjectId `bson:"movie" json:"movie"`
	User    bson.ObjectId `bson:"user" json:"user"`
	Comment string        `bson:"comment" json:"comment"`
}
