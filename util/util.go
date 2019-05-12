package util

import (
	"errors"
	"go-web-server/model"
	"regexp"

	"github.com/kataras/iris"

	"golang.org/x/crypto/bcrypt"
)

const (
	nickNamePattern      = "^[a-z0-9_.]{2,14}$"
	phonePattern         = "^[1,9]{1}[0,1,2,6,9]{1}[0-9]{7}$"
	passwordPattern      = "^[a-zA-Z0-9@.!#$%&'*+=?^_`{|}~-]{6,}$"
	userTypePattern      = "^[a-z ]{3,25}$"
	movieCategoryPattern = "^[a-z -]{3,25}$"
	movieNamePattern     = "^[a-zA-Z- 0-9]{2,25}$"
	theaterNamePattern   = "^[a-zA-Z 0-9]{1,25}$"
	descriptionPattern   = "^[a-zA-Z 0-9]{25,250}$"
	movieCommentPattern  = "^[a-zA-Z ]{1,}&"
)

//GeneratePasswordHash ...
func GeneratePasswordHash(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash)
}

//VerifyPassword ...
func VerifyPassword(hash string, password string) error {
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
func ValidateUserType(userType string) bool {
	matcher, err := regexp.Compile(userTypePattern)
	if err != nil {
		panic(err)
	}
	return matcher.MatchString(userType)
}

//ValidateMovieCategory ...
func ValidateMovieCategory(category string) bool {
	matcher, err := regexp.Compile(movieCategoryPattern)
	if err != nil {
		panic(err)
	}
	return matcher.MatchString(category)
}

//ValidateMovieName ...
func ValidateMovieName(name string) bool {
	matcher, err := regexp.Compile(movieNamePattern)
	if err != nil {
		panic(err)
	}
	return matcher.MatchString(name)
}

//ValidateTheaterName ...
func ValidateTheaterName(name string) bool {
	matcher, err := regexp.Compile(theaterNamePattern)
	if err != nil {
		panic(err)
	}
	return matcher.MatchString(name)
}

//ValidateDescription ...
func ValidateDescription(desc string) bool {
	matcher, err := regexp.Compile(descriptionPattern)
	if err != nil {
		panic(err)
	}
	return matcher.MatchString(desc)
}

//ValidateMovieComment ...
func ValidateMovieComment(comment string) bool {
	matcher, err := regexp.Compile(movieCommentPattern)
	if err != nil {
		panic(err)
	}
	return matcher.MatchString(comment)
}

//ValidateUserData ...
func ValidateUserData(user model.User) error {
	if !ValidateNickName(user.NickName) {
		return errors.New("invalid nick name")
	} else if !ValidatePhone(user.Phone) {
		return errors.New("invalid phone phone")
	} else if !ValidatePassword(user.PasswordHash) {
		return errors.New("invalid password")
	}
	return nil
}

//ValidateMovieData ...
func ValidateMovieData(movie model.Movie) error {
	if !ValidateMovieName(movie.Name) {
		return errors.New("invalid movie name")
	} else if !ValidateDescription(movie.Description) {
		return errors.New("invalid description characters")
	}
	return nil
}

//ValidateTheaterData ...
func ValidateTheaterData(theater model.Theater) error {
	if !ValidateTheaterName(theater.Name) {
		return errors.New("invalid theater name")
	} else if !ValidateDescription(theater.Description) {
		return errors.New("invalid description characters")
	}
	return nil
}

//JSON ...
func JSON(ctx iris.Context, data interface{}, code int, check bool) {
	ctx.StatusCode(code)
	ctx.JSON(map[string]interface{}{"error": check, "data": data})
}
