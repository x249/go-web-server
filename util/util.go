package util

import (
	"go-web-server/model"
	"github.com/kataras/iris"
	"regexp"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

const (
	nickNamePattern = "^[a-z0-9_.]{2,14}$"
	phonePattern    = "^[1,9]{1}[0,1,2,6,9]{1}[0-9]{7}$"
	passwordPattern = "^[a-zA-Z0-9@.!#$%&'*+=?^_`{|}~-]{6,}$"
	userTypePattern = "^[a-z ]{3,25}$"
)

//GeneratePasswordHash ...
func GeneratePasswordHash(password string) (string) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash)
}

//VerifyPassword ...
func  VerifyPassword(hash string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

//ValidateNickName ...
func ValidateNickName(nickName string) bool {
	matcher, err := regexp.Compile(nickNamePattern)
	if err != nil {
		panic(err)
	}
	return matcher.MatchString(nickName)
}

//ValidatePhone ...
func ValidatePhone(phone string) bool {
	matcher, err := regexp.Compile(phonePattern)
	if err != nil {
		panic(err)
	}
	return matcher.MatchString(phone)
}

//ValidatePassword ...
func ValidatePassword(passowrd string) bool {
	matcher, err := regexp.Compile(passwordPattern)
	if err != nil {
		panic(err)
	}
	return matcher.MatchString(passowrd)
}

//ValidateUserType ...
func ValidateUserType(userType string)bool{
	matcher, err := regexp.Compile(userTypePattern)
	if err != nil {panic(err)}
	return matcher.MatchString(userType)
}

//ValidateUserData ...
func ValidateUserData(user model.User)error{
	if !ValidateNickName(user.NickName){
		return errors.New("invalid nick name")
	}else if !ValidatePhone(user.Phone){
		return errors.New("invalid phone phone")
	}else if !ValidatePassword(user.PasswordHash){
		return errors.New("invalid password")
	}
	return nil
}

//JSON ...
func JSON(ctx iris.Context, data interface{}, code int, check bool){
	ctx.StatusCode(code)
	ctx.JSON(map[string]interface{}{"error": check, "data": data})
}
