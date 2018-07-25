package main

import (
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func main() {
	app := iris.New()
	/**
	* Logging
	 */
	app.Logger().SetLevel("debug")
	app.Use(recover.New())
	app.Use(logger.New())

	/**
	* Routes
	 */

	app.Get("/ping", func(ctx iris.Context) {
		ctx.Writef("Pong %d", int32(time.Now().Unix()))
	})

	//Start server
	app.Run(iris.Addr(":5555"), iris.WithoutServerError(iris.ErrServerClosed))
}
