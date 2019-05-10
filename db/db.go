package db

import (
	"errors"
	"go-web-server/model"
	"go-web-server/util"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

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

const (
	//UserCollection ...
	UserCollection = "users"
)

//NewUser ...
func (d *DB) NewUser(user model.User) error {
	return d.C(UserCollection).Insert(&user)
}

//CheckUser ...
func (d *DB) CheckUser(user model.User) (model.User, error) {
	var returnedUser model.User
	err := d.C(UserCollection).Find(bson.M{"phone": user.Phone}).One(&returnedUser)
	if err != nil {
		return returnedUser, err
	}
	if util.VerifyPassword(returnedUser.PasswordHash, user.PasswordHash) != nil {
		return returnedUser, errors.New("invalid user password")
	}
	return returnedUser, nil
}

//GetUsers ...
func (d *DB) GetUsers() ([]model.User, error) {
	var returnedUsers []model.User
	err := d.C(UserCollection).Find(bson.M{"active": false}).All(&returnedUsers)
	return returnedUsers, err
}

//GetUserByID ...
func (d *DB) GetUserByID(ID bson.ObjectId) (model.User, error) {
	var returnUser model.User
	err := d.C(UserCollection).Find(bson.M{"id": ID}).One(&returnUser)
	return returnUser, err
}