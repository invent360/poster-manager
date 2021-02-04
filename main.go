package main

import (
	"github.com/labstack/echo/v4"
	"github.com/post-manager/controllers"
)

func main()  {
	e := echo.New()

	e.GET("/", controllers.GetMessage)
	e.GET("/posts", controllers.FetchPosts)
	e.GET("/posts/:id", controllers.FetchPost)
	e.Logger.Fatal(e.Start(":8080"))
}
