package controller

import (
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
	user, err := c.GetUserByID(ctx.Params().Get("id"))
	if err != nil{
		util.JSON(ctx, err.Error(), 404, true)
	}else{
		util.JSON(ctx, user, 200, false)
	}
}

//AddMovieCategory ...
func (c *Controller) AddMovieCategory(ctx iris.Context){
	var category model.MovieCategory
	ctx.ReadJSON(&category)
	if 	!util.ValidateMovieCategory(category.Name){
		util.JSON(ctx, "invalid movie category name", 400, false)
	}else if err := c.NewMovieCategory(category); err != nil{
		util.JSON(ctx, "this category is already exists", 200, true)
	}else{
		util.JSON(ctx, "new movie category is added", 200, false)
	}
}

//GetMovieCategories ...
func (c *Controller)GetMovieCategories(ctx iris.Context){
	categories, err := c.GetCategories()
	if err != nil{
		util.JSON(ctx, err.Error(), 400, true)
	}else if categories != nil{
		util.JSON(ctx, categories, 200, false)
	}else {
		util.JSON(ctx, "no movie categories", 200, true)
	}
}

//DeleteMovieCategory ...
func (c *Controller)DeleteMovieCategory(ctx iris.Context){
	if err := c.RemoveMovieCategory(ctx.Params().Get("id")) ; err != nil{
		util.JSON(ctx, err.Error(), 200, true)
	}else{
		util.JSON(ctx, "movie category deleted", 200, false)
	}
}

//AddMovie ...
func (c *Controller) AddMovie(ctx iris.Context){
	var movie model.Movie
	ctx.ReadJSON(&movie)
	if err := util.ValidateMovieData(movie); err != nil{
		util.JSON(ctx, err.Error(), 400, true)
	}else if err := c.NewMovie(movie); err != nil{
		util.JSON(ctx, "error adding new movie", 400, true)
	}else {
		util.JSON(ctx, "new movie is added !", 200, false)		
	}
}

//AllMovies ...
func (c *Controller)AllMovies(ctx iris.Context){
	movies, err := c.GetMovies()
	if err != nil{
		util.JSON(ctx, err.Error(), 400, true)
	}else if movies != nil{
		util.JSON(ctx, movies, 200, false)
	}else {
		util.JSON(ctx, "no movies", 200, true)
	}
}

//MovieByCategory ...
func (c *Controller)MovieByCategory(ctx iris.Context){
	movies, err := c.GetMoviesByCategory(ctx.Params().Get("category"))
	if err != nil {
		util.JSON(ctx, err.Error(), 400, true)
	}else if movies != nil{
		util.JSON(ctx, movies, 200, false)
	}else{
		util.JSON(ctx, "no movies for this category", 200, true)
	}
}

//TheaterMovies ...
func (c *Controller)TheaterMovies(ctx iris.Context){
	movies, err := c.GetMoviesByTheater(ctx.Params().Get("theater"))
	if err != nil {
		util.JSON(ctx, err.Error(), 400, true)
	}else if movies != nil{
		util.JSON(ctx, movies, 200, false)
	}else{
		util.JSON(ctx, "no movies for this theater", 200, true)
	}
}

//MovieByID ...
func (c *Controller)MovieByID(ctx iris.Context){
	movie, err := c.GetMoviesByTheater(ctx.Params().Get("id"))
	if err != nil {
		util.JSON(ctx, err.Error(), 404, true)
	}else if movie != nil{
		util.JSON(ctx, movie, 200, false)
	}
}

//DeleteMovie ...
func (c *Controller)DeleteMovie(ctx iris.Context){
	if err := c.RemoveMovie(ctx.Params().Get("id")); err != nil{
		util.JSON(ctx, err.Error(), 400, true)
	}else{
		util.JSON(ctx, "movie deleted", 200, false)
	}
}

//MovieRate ...
func (c *Controller)MovieRate(ctx iris.Context){
	var rate model.MovieRate
	ctx.ReadJSON(&rate)
	if err := c.RateMovie(rate); err != nil{
		util.JSON(ctx, err.Error(), 200, true)
	}else{
		util.JSON(ctx, "movie rated", 200, true)
	}
}

//CommentMovie ...
func (c *Controller)CommentMovie(ctx iris.Context){
	var comment model.MovieComments
	ctx.ReadJSON(&comment)
	if err := c.NewMovieComment(comment); err != nil{
		util.JSON(ctx, err.Error(), 400, true)
	}else{
		util.JSON(ctx, "new comment", 200, false)
	}
}