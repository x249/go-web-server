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
	//UserTypeCollection ...
	UserTypeCollection = "user_types"
	//UserCollection ...
	UserCollection = "users"
)

//NewUser ...
func (d *DB) NewUser(user model.User) error {
	user.ID = bson.NewObjectId()
	user.PasswordHash = util.GeneratePasswordHash(user.PasswordHash)
	user.Phone = "+249" + user.Phone
	user.Active = false
	for _, key := range []string{"nick_name", "phone"} {
		index := mgo.Index{
			Key:    []string{key},
			Unique: true,
		}
		if err := d.C(UserCollection).EnsureIndex(index); err != nil {
			return err
		}
	}
	return d.C(UserCollection).Insert(&user)
}

//CheckUser ...
func (d *DB) CheckUser(user model.User) (model.User, error) {
	var returnedUser model.User
	user.Phone = "+249" + user.Phone
	err := d.C(UserCollection).Find(bson.M{"phone": user.Phone, "active": false}).One(&returnedUser)
	if err != nil {
		return returnedUser, errors.New("invalid user phone")
	} else if util.VerifyPassword(returnedUser.PasswordHash, user.PasswordHash) != nil {
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

//NewUserType ...
func (d *DB) NewUserType(userType model.UserType) error {
	userType.ID = bson.NewObjectId()
	for _, key := range []string{"type"} {
		index := mgo.Index{
			Key:    []string{key},
			Unique: true,
		}
		if err := d.C(UserTypeCollection).EnsureIndex(index); err != nil {
			return err
		}
	}
	return d.C(UserTypeCollection).Insert(&userType)
}

//GetUserTypes ...
func (d *DB) GetUserTypes() ([]model.UserType, error) {
	var userTypes []model.UserType
	err := d.C(UserTypeCollection).Find(bson.M{}).All(&userTypes)
	return userTypes, err
}
