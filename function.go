package german_grammar_server

import (
	"net/http"

	"github.com/tdidierjean/german_grammar/german_grammar_server/graphql"
)

// GraphQLServer Entry point for Google Cloud functions
// Deploy with: `gcloud functions deploy GraphQLServer --runtime go111 --trigger-http`
func GraphQLServer(w http.ResponseWriter, r *http.Request) {
	// graphql.Server()
	graphql.GetHandler().ServeHTTP(w, r)
}
