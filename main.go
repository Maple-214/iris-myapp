package main

import (
	"iris-myapp/routers"
	"iris-myapp/utils"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	app.Use(utils.Cors)

	booksAPI := app.Party("/v3/api/homePageContent")
	{
		booksAPI.Use(iris.Compression)
		booksAPI.Get("/", routers.Api_homePageContent)
	}
	app.Listen(":1111")
}
