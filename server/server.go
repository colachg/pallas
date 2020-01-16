package main

import (
	"github.com/colachg/pallas/graphql"
	"github.com/colachg/pallas/mysql"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"

	"github.com/99designs/gqlgen/handler"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	defaultPort = "8080"
	mysqlDSN    = "root:abc123_@tcp(127.0.0.1:3306)/dev?charset=utf8&parseTime=True&loc=Local"
)

func graphqlHandler(db *gorm.DB) echo.HandlerFunc {
	c := graphql.Config{Resolvers: &graphql.Resolver{
		ProjectRepo: mysql.ProjectRepo{DB: db},
	}}

	h := handler.GraphQL(graphql.NewExecutableSchema(c))
	return echo.WrapHandler(h)
}

func playgroundHandler() echo.HandlerFunc {
	h := handler.Playground("GraphQL playground", "/query")
	return echo.WrapHandler(h)
}

func main() {

	db, err := gorm.Open("mysql", mysqlDSN)
	defer db.Close()

	if err != nil {
		log.Panic(err)
	}
	//db.Debug().LogMode(true).AutoMigrate(&models.Project{})

	//start echo server
	e := echo.New()
	e.Use(middleware.Recover())
	e.POST("/query", graphqlHandler(db))
	e.GET("/", playgroundHandler())
	e.Logger.Fatal(e.Start(":" + defaultPort))
}
