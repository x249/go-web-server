package model

import (
	"github.com/globalsign/mgo/bson"
)

//Theater ...
type Theater struct {
	ID          bson.ObjectId    `bson:"id" json:"id"`
	Name        string           `bson:"name" json:"name"`
	Description string           `bson:"desc" json:"desc"`
	Photos      []string         `bson:"photos" json:"photos"`
	Location    *TheaterLocation `bson:"location" json:"location"`
}

//TheaterLocation ...
type TheaterLocation struct {
	Name      string  `bson:"name" json:"name"`
	Latitude  float32 `bson:"lat" json:"lat"`
	Longitude float32 `bson:"long" json:"long"`
}
