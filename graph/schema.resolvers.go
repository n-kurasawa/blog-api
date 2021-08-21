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
	var content *model.Content
	for _, v := range r.contents {
		if obj.ContentID == v.ID {
			return v, nil
		}
	}
	return content, nil
}

func (r *mutationResolver) CreateArticle(ctx context.Context, input model.NewArticle) (*model.Article, error) {
	contentID := fmt.Sprintf("content: %d", len(r.contents)+1)
	content := &model.Content{
		ID:   contentID,
		Body: input.Content,
	}
	article := &model.Article{
		ID:          fmt.Sprintf("article: %d", len(r.articles)+1),
		Slug:        input.Slug,
		Title:       input.Title,
		Date:        input.Date,
		CoverImage:  input.CoverImage,
		ContentID:   contentID,
		Description: input.Description,
	}
	r.contents = append(r.contents, content)
	r.articles = append(r.articles, article)
	return article, nil
}

func (r *queryResolver) Articles(ctx context.Context) ([]*model.Article, error) {
	return r.articles, nil
}

func (r *queryResolver) Article(ctx context.Context, slug string) (*model.Article, error) {
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
