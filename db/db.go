package db

import "github.com/globalsign/mgo"

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
	//MovieCategoryCollection ...
	MovieCategoryCollection= "movie_category"
	//MoviesCollection ...
	MoviesCollection = "movies"
	//MovieRateCollection ...
	MovieRateCollection = "movie_rate"
	//MovieCommentsCollection ...
	MovieCommentsCollection = "movie_comments"
)

//EnsureUniqueIndex ...
func (d *DB) EnsureUniqueIndex(collection string, indexes []string, data interface{}) error {
	for _, key := range indexes {
		index := mgo.Index{
			Key:    []string{key},
			Unique: true,
		}
		if err := d.C(collection).EnsureIndex(index); err != nil {
			return err
		}
	}
	return d.C(collection).Insert(&data)
}