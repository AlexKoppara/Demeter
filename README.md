# Demeter
Restaurant info GraphQL API

## Set Up
install dependencies with `go get ./...`

edit schema in `./graph/schema.graphqls`

regenerate files with `go run github.com/99designs/gqlgen`

update `entity.resolvers.go` and `schema.resolvers.go`

to run: `go run ./server.go`

visit `http://localhost:8080/` for GraphQL playground
