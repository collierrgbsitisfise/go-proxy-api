package main

import (
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	mgo "gopkg.in/mgo.v2"
)

//Proxy model
type Proxy struct {
	Time    string `json:"time"`
	IP      bool   `json:"ip"`
	Port    int    `json:"port"`
	Country string `json:"country"`
}

func main() {
	app := iris.New()
	/**
	* Logging
	 */
	app.Logger().SetLevel("debug")
	app.Use(recover.New())
	app.Use(logger.New())

	/**
	* Mongo
	 */
	const (
		Database   = "easy-links-db"
		Collection = "proxies"
	)

	session, err := mgo.Dial("mongodb://admin:vadim1@ds247330.mlab.com:47330/easy-links-db")

	// If there is an error connecting to Mongo - panic
	if err != nil {
		panic(err)
	}

	defer session.Close()

	db := session.DB(Database)
	collection := db.C(Collection)
	/**
	* Routes
	 */

	app.Get("/ping", func(ctx iris.Context) {
		ctx.Writef("Pong %d", int32(time.Now().Unix()))
	})

	app.Get("/proxies", func(ctx iris.Context) {
		var proxies []Proxy
		err := collection.Find(nil).All(&proxies)

		if err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.JSON(iris.Map{"message": "Can not fetch all proxies", "error": err})
			return
		}

		ctx.JSON(proxies)
	})
	//Start server
	app.Run(iris.Addr(":5555"), iris.WithoutServerError(iris.ErrServerClosed))
}
