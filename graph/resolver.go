//go:generate go run github.com/99designs/gqlgen
package graph

import (
	"database/sql"
	"github.com/n-kurasawa/blog-api/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	articles []*model.Article
	contents []*model.Content
	repository Repository
}

func NewResolver(db *sql.DB) *Resolver {
	return &Resolver{
		repository: NewSQLRepository(db),
	}
}