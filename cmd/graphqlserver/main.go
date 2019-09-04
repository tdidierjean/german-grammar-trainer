package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/tdidierjean/german_grammar/german_grammar_server/api/graphql"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(os.ExpandEnv("$GOPATH/src/github.com/tdidierjean/german_grammar/german_grammar_server/.env")); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	graphql.Server()
}
