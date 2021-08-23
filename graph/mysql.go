package graph

import (
	"database/sql"
	"strconv"
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

func (r *SQLRepository) GetArticles() ([]*model.Article, error) {
	rows, err := r.db.Query("select id, slug, title, date, cover_image, description, content_id from articles")
	if err != nil {
		return nil, err
	}
	articles := make([]*model.Article, 0)
	for rows.Next() {
		article := model.Article{}
		if err := rows.Scan(&article.ID, &article.Slug, &article.Title, &article.Date, &article.CoverImage, &article.Description, &article.ContentID); err != nil {
			return nil, err
		}
		articles = append(articles, &article)
	}
	return articles, nil
}

func (r *SQLRepository) GetArticle(slug string) (*model.Article, error) {
	row := r.db.QueryRow("select id, slug, title, date, cover_image, description, content_id from articles where slug = ?", slug)
	article := model.Article{}
	if err := row.Scan(&article.ID, &article.Slug, &article.Title, &article.Date, &article.CoverImage, &article.Description, &article.ContentID); err != nil {
		return nil, err
	}
	return &article, nil
}

func (r *SQLRepository) CreateArticle(article model.NewArticle) (*model.Article, error) {
	result, err := r.db.Exec("insert into contents (body) values (?)", article.Content)
	if err != nil {
		return nil, err
	}
	query := "insert into articles (slug, title, date, cover_image, description, content_id) values (?, ?, ?, ?, ?, ?)"
	contentID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	date := time.Now()
	result, err = r.db.Exec(query, article.Slug, article.Title, date, article.CoverImage, article.Description, contentID)
	if err != nil {
		return nil, err
	}
	articleID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &model.Article{
		ID:          strconv.Itoa(int(articleID)),
		Slug:        article.Slug,
		Title:       article.Title,
		Date:        date.Format(time.RFC3339),
		CoverImage:  article.CoverImage,
		Description: article.Description,
		ContentID:   strconv.Itoa(int(contentID)),
	}, nil
}
