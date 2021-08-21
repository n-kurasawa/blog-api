package graph

import (
	"database/sql"
	"github.com/n-kurasawa/blog-api/graph/model"
)

type SQLRepository struct {
	db *sql.DB
}

func NewSQLRepository(db *sql.DB) *SQLRepository {
	return &SQLRepository{
		db: db,
	}
}

func (r *SQLRepository) GetContent(id int) *model.Content {
	return nil
}

func (r *SQLRepository) GetArticles() []model.Article {
	return nil
}