package main

import (
	"fmt"

	"github.com/kataras/iris"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

const (
	host = "127.0.0.1"
	db   = "test"

	user    = "users"
	movie   = "movies"
	comment = "comments"
)

//User ...
type User struct {
	ID   bson.ObjectId `bson:"id" json:"id"`
	Name string        `bson:"name" json:"name"`
}

//Movie ..
type Movie struct {
	ID   bson.ObjectId `bson:"id" json:"id"`
	Name string        `bson:"name" json:"name"`
}

//Comment ...
type Comment struct {
	ID      bson.ObjectId `bson:"id" json:"id"`
	User    bson.ObjectId `bson:"user" json:"user"`
	Movie   bson.ObjectId `bson:"movie" json:"movie"`
	Comment string        `bson:"comment" json:"comment"`
}

//DB ...
type DB struct {
	*mgo.Database
}

//NewDB ...
func NewDB(host, db string) *DB {
	session, err := mgo.Dial(host)
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	defer session.Close()
	sess := session.Clone()
	return &DB{sess.DB(db)}
}

func main() {
	db := NewDB(host, db)
	collection := db.C(comment)
	pipeline := []bson.M{
		{"$lookup": bson.M{
			"from":         user,
			"localField":   "user",
			"foreignField": "id",
			"as":           "user_comment",
		}},
		{"$unwind": "$user_comment"},
		{"$lookup": bson.M{
			"from":         movie,
			"localField":   "movie",
			"foreignField": "id",
			"as":           "movie_comment",
		}},
		{"$unwind": "$movie_comment"},
		{"$project": bson.M{
			"_id":        1,
			"id":         1,
			"comment":    1,
			"user_id":    "$user_comment.id",
			"user_name":  "$user_comment.name",
			"movie_id":   "$movie_comment.id",
			"movie_name": "$movie_comment.name",
		}},
	}

	pipe := collection.Pipe(pipeline)
	rest := []bson.M{}
	err := pipe.All(&rest)
	if err != nil {
		fmt.Printf("%v", err.Error())
	}
	app := iris.Default()
	app.Get("/", func(ctx iris.Context) {
		ctx.JSON(rest)
	})
	app.Run(iris.Addr(":8080"))
}

func (d *DB) addNewUsers() {
	user1 := User{bson.NewObjectId(), "majid"}
	user2 := User{bson.NewObjectId(), "sayed"}
	user3 := User{bson.NewObjectId(), "ahmed"}
	user4 := User{bson.NewObjectId(), "yassin"}

	d.C(user).Insert(user1, user2, user3, user4)

	movie1 := Movie{bson.NewObjectId(), "the avenger"}
	movie2 := Movie{bson.NewObjectId(), "the equalizer"}
	movie3 := Movie{bson.NewObjectId(), "the rite"}
	movie4 := Movie{bson.NewObjectId(), "the end"}

	d.C(movie).Insert(movie1, movie2, movie3, movie4)
}

func (d *DB) addNewComments() {
	comment1 := Comment{bson.NewObjectId(), bson.ObjectIdHex("5cd87e2f8bdc5b1740820e6f"), bson.ObjectIdHex("5cd87e2f8bdc5b1740820e73"), "nice one"}
	comment2 := Comment{bson.NewObjectId(), bson.ObjectIdHex("5cd87e2f8bdc5b1740820e70"), bson.ObjectIdHex("5cd87e2f8bdc5b1740820e74"), "nice two"}
	comment3 := Comment{bson.NewObjectId(), bson.ObjectIdHex("5cd87e2f8bdc5b1740820e71"), bson.ObjectIdHex("5cd87e2f8bdc5b1740820e74"), "nice three"}
	comment4 := Comment{bson.NewObjectId(), bson.ObjectIdHex("5cd87e2f8bdc5b1740820e72"), bson.ObjectIdHex("5cd87e2f8bdc5b1740820e76"), "nice four"}
	comment5 := Comment{bson.NewObjectId(), bson.ObjectIdHex("5cd87e2f8bdc5b1740820e6f"), bson.ObjectIdHex("5cd87e2f8bdc5b1740820e76"), "nice five"}
	comment6 := Comment{bson.NewObjectId(), bson.ObjectIdHex("5cd87e2f8bdc5b1740820e71"), bson.ObjectIdHex("5cd87e2f8bdc5b1740820e75"), "nice six"}
	comment7 := Comment{bson.NewObjectId(), bson.ObjectIdHex("5cd87e2f8bdc5b1740820e70"), bson.ObjectIdHex("5cd87e2f8bdc5b1740820e75"), "nice seven"}
	comment8 := Comment{bson.NewObjectId(), bson.ObjectIdHex("5cd87e2f8bdc5b1740820e72"), bson.ObjectIdHex("5cd87e2f8bdc5b1740820e73"), "nice eight"}
	d.C(comment).Insert(comment1, comment2, comment3, comment4, comment5, comment6, comment7, comment8)
}
