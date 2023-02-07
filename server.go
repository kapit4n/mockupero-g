package main

import (
	"log"
	"net/http"
	"os"
	"stats-mockupero/graph"
	"stats-mockupero/graph/common"
	"stats-mockupero/graph/resolvers"
	"stats-mockupero/services"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = defaultPort
	}

	db, err := common.InitDb()
	if err != nil {
		log.Fatal(err)
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &resolvers.Resolver{}}))

	customCtx := &common.CustomContext{
		Database: db,
	}

	var newsApi services.NewsApi

	newsApi = services.RapidApi{}
	newsApi.GetNews()

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", common.CreateContext(customCtx, srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
