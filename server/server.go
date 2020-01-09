package main

import (
	"github.com/colachg/pallas/graphql"
	"github.com/colachg/pallas/models"
	"github.com/colachg/pallas/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	defaultPort = "8080"
	mysqlDSN    = "root:abc123_@tcp(127.0.0.1:3306)/dev?charset=utf8&parseTime=True&loc=Local"
)

func main() {

	db, err := gorm.Open("mysql", mysqlDSN)
	defer db.Close()

	if err != nil {
		log.Panic(err)
	}
	db.Debug().AutoMigrate(&models.Project{})

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	c := graphql.Config{Resolvers: &graphql.Resolver{
		ProjectRepo: mysql.ProjectRepo{DB: db},
	}}

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(graphql.NewExecutableSchema(c)))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
