# german-grammar-trainer
Generate random German grammar exercises, comes with a GraphQL server and a CLI interface.

See [german-grammar-webapp](https://github.com/tdidierjean/german-grammar-webapp) for a web client to the GraphQL API ([see it live here](https://german-grammar.netlify.com/)).

## Running locally

```
# run the command line app
go run cmd/cli/main.go

# run the GraphQL server
go run cmd/graphqlserver/main.go

# run tests
go test ./...
```

## Deploying to Google cloud functions
```
gcloud functions deploy GraphQLServer --runtime go111 --trigger-http
```
