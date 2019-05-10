package controller

import (
	"github.com/globalsign/mgo/bson"
	"go-web-server/util"
	"go-web-server/model"
	"go-web-server/db"

	"github.com/kataras/iris"
)

//Controller ...
type Controller struct {
	*db.DB
}

//NewController ...
func NewController(db *db.DB) *Controller {
	return &Controller{db}
}

//Signup ...
func (c *Controller)Signup(ctx iris.Context){
	var user model.User
	err := ctx.ReadJSON(&user)
	if err != nil{
		panic(err)
	}
	err = c.NewUser(user)
	if err != nil{
		util.JSON(ctx, err, 400, true)
	}
	util.JSON(ctx, err, 200, false)
}

//Login ...
func (c *Controller) Login(ctx iris.Context) {
	var user model.User
	err := ctx.ReadJSON(&user)
	if err != nil{
		panic(err)
	}
	returnedUser, err := c.CheckUser(user)
	if err != nil{
		util.JSON(ctx, err, 400, true)
	}
	util.JSON(ctx, returnedUser, 200, false)
}


//AllUsers ...
func (c *Controller)AllUsers(ctx iris.Context){
	users, err := c.GetUsers()
	if err != nil{
		util.JSON(ctx, err, 400, true)
	}
	util.JSON(ctx, users, 200, false)
}

//GetByID ...
func (c *Controller)GetByID(ctx iris.Context){
	ID := ctx.Params().Get("id")
	user, err := c.GetUserByID(bson.ObjectIdHex(ID))
	if err != nil{
		util.JSON(ctx, err, 400, true)
	}
	util.JSON(ctx, user, 200, false)
}