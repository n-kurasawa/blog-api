package graph

import (
	"database/sql"
	"strconv"
	"strings"
	"time"

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
	row := r.db.QueryRow("select id, body from contents where id = ?", id)
	content := model.Content{}
	if err := row.Scan(&content.ID, &content.Body); err != nil {
		return nil, err
	}
	return &content, nil
}

func (r *SQLRepository) GetPosts() ([]*model.Post, error) {
	rows, err := r.db.Query("select id, slug, title, published_at, cover_image, description, content_id from posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	posts := make([]*model.Post, 0)
	for rows.Next() {
		post := model.Post{}
		var publishedAt time.Time
		if err := rows.Scan(&post.ID, &post.Slug, &post.Title, &publishedAt, &post.CoverImage, &post.Description, &post.ContentID); err != nil {
			return nil, err
		}
		post.PublishedAt = publishedAt.Format("2006-01-02")
		posts = append(posts, &post)
	}
	return posts, nil
}

func (r *SQLRepository) GetPost(slug string) (*model.Post, error) {
	row := r.db.QueryRow("select id, slug, title, published_at, cover_image, description, content_id from posts where slug = ?", slug)
	post := model.Post{}
	var publishedAt time.Time
	if err := row.Scan(&post.ID, &post.Slug, &post.Title, &publishedAt, &post.CoverImage, &post.Description, &post.ContentID); err != nil {
		return nil, err
	}
	post.PublishedAt = publishedAt.Format("2006-01-02")
	return &post, nil
}

func (r *SQLRepository) CreatePost(post model.NewPost) (*model.Post, error) {
	result, err := r.db.Exec("insert into contents (body) values (?)", post.Content)
	if err != nil {
		return nil, err
	}
	query := "insert into posts (slug, title, published_at, cover_image, description, content_id, created_at) values (?, ?, ?, ?, ?, ?, ?)"
	contentID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	result, err = r.db.Exec(query, post.Slug, post.Title, post.PublishedAt, post.CoverImage, post.Description, contentID, time.Now())
	if err != nil {
		return nil, err
	}
	postID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &model.Post{
		ID:          strconv.Itoa(int(postID)),
		Slug:        post.Slug,
		Title:       post.Title,
		PublishedAt: post.PublishedAt,
		CoverImage:  post.CoverImage,
		Description: post.Description,
		ContentID:   strconv.Itoa(int(contentID)),
	}, nil
}

type ContentSQLRepository struct {
	db *sql.DB
}

func NewContentSQLRepository(db *sql.DB) *ContentSQLRepository {
	return &ContentSQLRepository{db: db}
}

func (r *ContentSQLRepository) GetContents(ids []string) ([]*model.Content, []error) {
	placeholders := make([]string, len(ids))
	args := make([]interface{}, len(ids))
	for i := 0; i < len(ids); i++ {
		placeholders[i] = "?"
		args[i] = ids[i]
	}

	rows, err := r.db.Query("select id, body from contents where id in ("+strings.Join(placeholders, ",")+")", args...)
	if err != nil {
		return nil, []error{err}
	}
	defer rows.Close()

	// 引数のidsの順序のまま返す必要がるため一度mapにしてids順の配列に変換してる
	findByID := make(map[string]*model.Content)
	for rows.Next() {
		content := model.Content{}
		err := rows.Scan(&content.ID, &content.Body)
		if err != nil {
			return nil, []error{err}
		}
		findByID[content.ID] = &content
	}
	contents := make([]*model.Content, len(ids))
	for i, id := range ids {
		contents[i] = findByID[id]
	}
	return contents, nil
}
