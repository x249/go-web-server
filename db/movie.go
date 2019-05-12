package db

import (
	"errors"
	"go-web-server/model"
	"go-web-server/util"

	"github.com/globalsign/mgo/bson"
)

//NewMovieCategory ...
func (d *DB) NewMovieCategory(category model.MovieCategory) error {
	category.ID = bson.NewObjectId()
	if !util.ValidateMovieCategory(category.Name){
		return errors.New("invalid movie category name")
	}
	return d.EnsureUniqueIndex(MovieCategoryCollection, []string{"name"}, category)
}

//GetCategories ...
func (d *DB) GetCategories()([]model.MovieCategory, error){
	var categories []model.MovieCategory
	err := d.C(MovieCategoryCollection).Find(bson.M{}).All(&categories)
	return categories, err
}

//RemoveMovieCategory ...
func (d *DB) RemoveMovieCategory (ID string)error{
	return d.C(MovieCategoryCollection).Remove(bson.M{"id": bson.ObjectIdHex(ID)})
}

//NewMovie ...
func (d *DB) NewMovie(movie model.Movie)error{
	movie.ID = bson.NewObjectId()
	return d.EnsureUniqueIndex(MoviesCollection, []string{}, movie)
}

//GetMovies ...
func (d *DB) GetMovies()([]model.Movie, error){
	 var movies []model.Movie
	 err := d.C(MoviesCollection).Find(bson.M{}).All(&movies)
	 return movies, err
}

//GetMoviesByCategory ...
func (d *DB) GetMoviesByCategory(category string)([]model.Movie, error){
	var movies []model.Movie
	err := d.C(MoviesCollection).Find(bson.M{"cate": bson.ObjectIdHex(category)}).All(&movies)
	return movies, err
}

//GetMoviesByTheater ...
func (d *DB) GetMoviesByTheater(theater string)([]model.Movie, error){
	var movies []model.Movie
	err := d.C(MoviesCollection).Find(bson.M{"theater": bson.ObjectIdHex(theater)}).All(&movies)
	return movies, err
}

//GetMoviesByID ...
func (d *DB) GetMoviesByID(ID string)(model.Movie, error){
	var movie model.Movie
	err := d.C(MoviesCollection).Find(bson.M{"id": bson.ObjectIdHex(ID)}).All(&movie)
	return movie, err
}

//RemoveMovie ...
func (d *DB) RemoveMovie(ID string)error{
	return d.C(MoviesCollection).Remove(bson.M{"id": bson.ObjectIdHex(ID)})
}

//RateMovie ...
func (d *DB) RateMovie(rate model.MovieRate)error{
	return d.C(MovieRateCollection).Insert(&rate)
}

//NewMovieComment ...
func (d *DB) NewMovieComment(comment model.MovieComments)error{
	return d.C(MovieCommentsCollection).Insert(&comment)
}