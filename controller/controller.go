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

//AddUserType ...
func (c *Controller)AddUserType(ctx iris.Context){
	var userType model.UserType
	ctx.ReadJSON(&userType)
	if !util.ValidateUserType(userType.Type){
		util.JSON(ctx, "invalid user type name", 400, true)
	}else if err := c.NewUserType(userType); err != nil{
		util.JSON(ctx, "this is user type is already exits", 200, true)
	}else {
		util.JSON(ctx, "new user type is added", 200, false)
	}
}

//GetAllUserTypes ...
func (c *Controller)GetAllUserTypes(ctx iris.Context){
	userTypes, err := c.GetUserTypes()
	if err != nil{
		util.JSON(ctx, err, 400, true)
	}else if userTypes != nil{
		util.JSON(ctx, userTypes, 200, false)
	}else {
		util.JSON(ctx, "no user types", 200, true)
	}
}

//Signup ...
func (c *Controller)Signup(ctx iris.Context){
	var user model.User
	ctx.ReadJSON(&user)
	if err := util.ValidateUserData(user); err != nil{
		util.JSON(ctx, err.Error(), 400, true)
	}else if err := c.NewUser(user); err != nil{
		util.JSON(ctx, "this is user is already exist", 200, true)
	}else {
		util.JSON(ctx, "account created !", 200, false)		
	}
}

//Login ...
func (c *Controller) Login(ctx iris.Context) {
	var user model.User
	ctx.ReadJSON(&user)
	
	returnedUser, err := c.CheckUser(user)
	if err != nil{
		util.JSON(ctx, err.Error(), 400, true)
	}else {
		util.JSON(ctx, returnedUser, 200, false)	
	}
}


//AllUsers ...
func (c *Controller)AllUsers(ctx iris.Context){
	users, err := c.GetUsers()
	if err != nil{
		util.JSON(ctx, err.Error(), 400, true)
	}

	if users != nil{
		util.JSON(ctx, users, 200, false)
	}else {
		util.JSON(ctx, "no users", 200, true)	
	}
}

//GetByID ...
func (c *Controller)GetByID(ctx iris.Context){
	ID := ctx.Params().Get("id")
	user, err := c.GetUserByID(bson.ObjectIdHex(ID))
	if err != nil{
		util.JSON(ctx, err.Error(), 400, true)
	}else{
		util.JSON(ctx, user, 200, false)
	}
}