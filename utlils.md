
# Modify models
- update graph/schema.graphqls

go get -d github.com/99designs/gqlgen
go run github.com/99designs/gqlgen generate

# Run it
- go run main.go
