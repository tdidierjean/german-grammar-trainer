package german_grammar_server

import (
	"net/http"

	"github.com/tdidierjean/german_grammar/german_grammar_server/api/graphql"
)

// GraphQLServer Entry point for Google Cloud functions
// Deploy with: `gcloud functions deploy GraphQLServer --runtime go111 --trigger-http`
// Note that ENV variables need to be set up in the cloud function configuration
func GraphQLServer(w http.ResponseWriter, r *http.Request) {
	graphql.GetHandler().ServeHTTP(w, r)
}
