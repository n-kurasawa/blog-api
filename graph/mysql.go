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

func (r *SQLRepository) GetContent(id string) (*model.Content, error) {
	return nil, nil
}

func (r *SQLRepository) GetArticles() ([]*model.Article, error) {
	return nil, nil
}

func (r *SQLRepository) GetArticle(slug string) (*model.Article, error) {
	return nil, nil
}

func (r *SQLRepository) SaveArticle(article model.NewArticle) (*model.Article, error) {
	return nil, nil
}