package graph

import (
	"context"
	"net/http"
	"time"

	"github.com/n-kurasawa/blog-api/graph/model"
)

const loadersKey = "dataloaders"

type Loaders struct {
	ContentByID ContentLoader
}

func Middleware(repo ContentRepository, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), loadersKey, &Loaders{
			ContentByID: ContentLoader{
				maxBatch: 100,
				wait:     1 * time.Millisecond,
				fetch: func(ids []string) ([]*model.Content, []error) {
					return repo.GetContents(ids)
				},
			},
		})
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}
