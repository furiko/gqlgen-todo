package main

import (
	"database/sql"
	"github.com/furiko/gqlgen-todo/graph/model"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/furiko/gqlgen-todo/graph"
	"github.com/furiko/gqlgen-todo/graph/generated"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/gorp.v1"

)

const defaultPort = "8080"
//var dbMap *gorp.DbMap

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	dbMap := initDB()
	defer dbMap.Db.Close()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: dbMap}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}


func initDB() *gorp.DbMap {
	db, err := sql.Open("mysql", "root:root@tcp(db:3306)/todo?parseTime=true")
	checkErr(err, "Connect DB failed")

	dbMap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"Innodb", "UTF8"}}
	dbMap.AddTableWithName(model.Todo{}, "todos").SetKeys(false, "ID")
	return dbMap
}

func checkErr(err error, message string) {
	if err != nil {
		log.Fatalln(message)
	}
}
