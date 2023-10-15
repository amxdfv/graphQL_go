package main

import (
	"go-graphql-api/graph"
	"go-graphql-api/logs_logic"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	var Os_log logs_logic.Simple_log
	Os_log.Message = "Запускаем сервер, работяги"

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	logs_logic.Write_usual_log(Os_log)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
