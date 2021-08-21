package graph

import "github.com/n-kurasawa/blog-api/graph/model"

type Repository interface {
	GetContent(id int) *model.Content
	GetArticles() []model.Article
}
