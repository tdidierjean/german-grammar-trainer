package graphql

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/gorilla/handlers"
	grammarexercise "github.com/tdidierjean/german_grammar/german_grammar_server/app"
)

const defaultPort = "8080"

// Server start a GraphQL server
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

// GetHandler configure the GraphQL handler
func GetHandler() http.Handler {
	connection := grammarexercise.NewDatabaseConnection()

	resolver := Resolver{
		connection: connection,
	}

	return handlers.CORS(
		handlers.AllowedHeaders([]string{"content-type"}),
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowCredentials(),
	)(handlers.LoggingHandler(os.Stdout, handler.GraphQL(NewExecutableSchema(Config{Resolvers: &resolver}))))
}
