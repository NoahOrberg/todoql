package main

import (
	log "log"
	http "net/http"
	os "os"

	handler "github.com/99designs/gqlgen/handler"
	todoql "github.com/NoahOrberg/todoql"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(todoql.NewExecutableSchema(todoql.Config{Resolvers: &todoql.Resolver{}})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}