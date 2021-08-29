package main

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	"github.com/rs/cors"

	"github.com/n-kurasawa/blog-api/graph"
	"github.com/n-kurasawa/blog-api/graph/generated"
)

var adapter *httpadapter.HandlerAdapter

func init() {
	//db := initDB()
	//defer db.Close()
	// TODO
	var db *sql.DB

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
