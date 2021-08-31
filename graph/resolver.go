//go:generate go install github.com/99designs/gqlgen@latest
//go:generate gqlgen
package graph

import (
	"database/sql"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	repository Repository
}

func NewResolver(db *sql.DB) *Resolver {
	return &Resolver{
		repository: NewSQLRepository(db),
	}
}
