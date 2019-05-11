package main

import (
	"github.com/kataras/iris"
	"go-web-server/db"
	"go-web-server/controller"
	"github.com/BurntSushi/toml"
	"fmt"
)

//Config ...
type Config struct {
	Host string
	DB string
	Port int
}

func main() {
	var conf Config
	if _, err := toml.DecodeFile("config.toml", &conf); err != nil {
		fmt.Printf("%v", err.Error())
	}
	db := db.NewDB(conf.Host, conf.DB)
	cntr := controller.NewController(db)

	app := iris.Default()

	app.PartyFunc("/type", func(t iris.Party){
		t.Get("/", cntr.GetAllUserTypes)
		t.Post("/", cntr.AddUserType)
	})

	app.PartyFunc("/user", func(u iris.Party){
		u.Get("/", cntr.AllUsers)
		u.Get("/{id:string}", cntr.GetByID)
		u.Post("/login", cntr.Login)
		u.Post("/signup", cntr.Signup)
	})
	
	app.Run(iris.Addr(fmt.Sprintf(":%d", conf.Port)))
}