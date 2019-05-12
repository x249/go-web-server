package db

import (
	"errors"

	"go-web-server/model"
	"go-web-server/util"

	"github.com/globalsign/mgo/bson"
)

//NewUser ...
func (d *DB) NewUser(user model.User) error {
	user.ID = bson.NewObjectId()
	user.PasswordHash = util.GeneratePasswordHash(user.PasswordHash)
	user.Phone = "+249" + user.Phone
	user.Active = false
	return d.EnsureUniqueIndex(UserCollection, []string{"nick_name", "phone"}, user)
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
func (d *DB) GetUserByID(ID string) (model.User, error) {
	var returnUser model.User
	err := d.C(UserCollection).Find(bson.M{"id": bson.ObjectIdHex(ID)}).One(&returnUser)
	return returnUser, err
}

//NewUserType ...
func (d *DB) NewUserType(userType model.UserType) error {
	userType.ID = bson.NewObjectId()
	return d.EnsureUniqueIndex(UserTypeCollection, []string{"type"}, userType)
}

//GetUserTypes ...
func (d *DB) GetUserTypes() ([]model.UserType, error) {
	var userTypes []model.UserType
	err := d.C(UserTypeCollection).Find(bson.M{}).All(&userTypes)
	return userTypes, err
}
