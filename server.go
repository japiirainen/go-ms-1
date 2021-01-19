package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/japiirainen/go-ms-1/graphql/db"
	"github.com/japiirainen/go-ms-1/graphql/generated"
	"github.com/japiirainen/go-ms-1/graphql/resolvers"
	"github.com/joho/godotenv"
)

const defaultPort = "5000"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	err := godotenv.Load(".env")
	dbURL := os.Getenv("DATABASE_URL")
	if err != nil {
		log.Fatal("failed to lead env")
	}
	c := db.Connect(os.Getenv("DATABASE_URL"))
	db.MigrateUp(dbURL)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{
		Conn: c,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
