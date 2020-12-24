package main

import (
	"log"

	"github.com/dev-sota/echo-gorm-graphql-example/datastore"
	"github.com/dev-sota/echo-gorm-graphql-example/graphql"
	"github.com/dev-sota/echo-gorm-graphql-example/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db := datastore.NewDB()
	defer db.Close()
	handler.SetDB(db)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/users", handler.GetUsers)

	// endpoint using graphql
	h, _ := graphql.NewHandler(db)
	e.POST("/graphql", echo.WrapHandler(h))

	if err := e.Start("localhost:1323"); err != nil {
		log.Fatalln(err)
	}

}
