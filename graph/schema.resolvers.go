package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/n-kurasawa/blog-api/graph/generated"
	"github.com/n-kurasawa/blog-api/graph/model"
)

func (r *articleResolver) Content(ctx context.Context, obj *model.Article) (*model.Content, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateArticle(ctx context.Context, input model.NewArticle) (*model.Article, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Articles(ctx context.Context) ([]*model.Article, error) {
	panic(fmt.Errorf("not implemented"))
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
