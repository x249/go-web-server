package model

import (
	"github.com/globalsign/mgo/bson"
)

//MovieTicket ...
type MovieTicket struct {
	ID         bson.ObjectId `bson:"id" json:"id"`
	Movie      bson.ObjectId `bson:"movie" json:"movie"`
	ShowTime   string        `bson:"show_time" json:"show_time"`
	Date       string        `bson:"date" json:"date"`
	ExpireAt   string        `bson:"expire_at" json:"expire_at"`
	Price      string        `bson:"price" json:"price"`
	SeatNumber int           `bson:"seat" json:"seat"`
}

//UserTicker ...
type UserTicker struct {
	MovieTicket bson.ObjectId `bson:"movie_ticker" json:"movie_ticket"`
	User        bson.ObjectId `bson:"user" json:"user"`
	BuyDate     string        `bson:"buy_date" json:"buy_date"`
	Scanned     bool          `bson:"scanned" json:"scanned"`
}
