// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Content struct {
	ID   string `json:"id"`
	Body string `json:"body"`
}

type EditPost struct {
	ID          string `json:"id"`
	Slug        string `json:"slug"`
	Title       string `json:"title"`
	CoverImage  string `json:"coverImage"`
	Content     string `json:"content"`
	Description string `json:"description"`
	PublishedAt string `json:"publishedAt"`
}

type NewPost struct {
	Slug        string `json:"slug"`
	Title       string `json:"title"`
	CoverImage  string `json:"coverImage"`
	Content     string `json:"content"`
	Description string `json:"description"`
	PublishedAt string `json:"publishedAt"`
}
