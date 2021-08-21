package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/n-kurasawa/blog-api/graph/generated"
	"github.com/n-kurasawa/blog-api/graph/model"
)

func (r *articleResolver) Content(ctx context.Context, obj *model.Article) (*model.Content, error) {
	content, err := r.repository.GetContent(obj.ContentID)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func (r *mutationResolver) CreateArticle(ctx context.Context, input model.NewArticle) (*model.Article, error) {
	article, err := r.repository.SaveArticle(input)
	if err != nil {
		return nil, err
	}
	return article, nil
}

func (r *queryResolver) Articles(ctx context.Context) ([]*model.Article, error) {
	articles, err := r.repository.GetArticles()
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (r *queryResolver) Article(ctx context.Context, slug string) (*model.Article, error) {
	article, err := r.repository.GetArticle(slug)
	if err != nil {
		return nil, err
	}
	return article, nil
}

// Article returns generated.ArticleResolver implementation.
func (r *Resolver) Article() generated.ArticleResolver { return &articleResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type articleResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
