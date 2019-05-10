package util

import (
	"github.com/kataras/iris"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

const (
	nickNamePattern = "^[a-z0-9_.]{2,14}$"
	phonePattern    = "^[0]{1}[1,9]{1}[0,1,2,6,9]{1}[0-9]{7}$"
	passwordPattern = "^[a-zA-Z0-9@.!#$%&'*+=?^_`{|}~-]{6,}$"
)

//GeneratePasswordHash ...
func GeneratePasswordHash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

//VerifyPassword ...
func VerifyPassword(hash []byte, password []byte) error {
	return bcrypt.CompareHashAndPassword(hash, password)
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


//JSON ...
func JSON(ctx iris.Context, data interface{}, code int, check bool){
	ctx.StatusCode(code)
	ctx.JSON(map[string]interface{}{"error": check, "data": data})
}
