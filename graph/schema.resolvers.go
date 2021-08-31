package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/n-kurasawa/blog-api/graph/generated"
	"github.com/n-kurasawa/blog-api/graph/model"
)

func (r *mutationResolver) CreatePost(ctx context.Context, input model.NewPost) (*model.Post, error) {
	post, err := r.repository.CreatePost(input)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (r *postResolver) Content(ctx context.Context, obj *model.Post) (*model.Content, error) {
	return For(ctx).ContentByID.Load(obj.ContentID)
}

func (r *queryResolver) Posts(ctx context.Context) ([]*model.Post, error) {
	posts, err := r.repository.GetPosts()
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *queryResolver) Post(ctx context.Context, slug string) (*model.Post, error) {
	post, err := r.repository.GetPost(slug)
	if err != nil {
		return nil, err
	}
	return post, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Post returns generated.PostResolver implementation.
func (r *Resolver) Post() generated.PostResolver { return &postResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type postResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
