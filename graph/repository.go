package graph

import "github.com/n-kurasawa/blog-api/graph/model"

type Repository interface {
	GetContent(id string) (*model.Content, error)
	GetPosts() ([]*model.Post, error)
	GetPost(slug string) (*model.Post, error)
	CreatePost(post model.NewPost) (*model.Post, error)
}

type ContentRepository interface {
	GetContents(ids []string) ([]*model.Content, []error)
}
