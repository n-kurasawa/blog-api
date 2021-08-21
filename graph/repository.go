package graph

import "github.com/n-kurasawa/blog-api/graph/model"

type Repository interface {
	GetContent(id string) (*model.Content, error)
	GetArticles() ([]*model.Article, error)
	GetArticle(slug string) (*model.Article, error)
	SaveArticle(article model.NewArticle) (*model.Article, error)
}
