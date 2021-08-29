package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/cors"

	"github.com/n-kurasawa/blog-api/graph"
	"github.com/n-kurasawa/blog-api/graph/generated"
)

var adapter *httpadapter.HandlerAdapter

func init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/blog?parseTime=true&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"))
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: graph.NewResolver(db)}))
	repo := graph.NewContentSQLRepository(db)
	http.Handle("/query", cors.Default().Handler(graph.Middleware(repo, srv)))
	http.Handle("/playground", playground.Handler("GraphQL playground", "/query"))

	adapter = httpadapter.New(http.DefaultServeMux)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return adapter.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
