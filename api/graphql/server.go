package graphql

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/gorilla/handlers"
)

const defaultPort = "8080"

func Server() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))

	http.Handle("/query", GetHandler())

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func GetHandler() http.Handler {
	return handlers.CORS(
		handlers.AllowedHeaders([]string{"content-type"}),
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowCredentials(),
	)(handlers.LoggingHandler(os.Stdout, handler.GraphQL(NewExecutableSchema(Config{Resolvers: &Resolver{}}))))
}
