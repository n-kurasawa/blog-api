package model

type Post struct {
	ID          string `json:"id"`
	Slug        string `json:"slug"`
	Title       string `json:"title"`
	Date        string `json:"date"`
	CoverImage  string `json:"coverImage"`
	ContentID   string `json:"content"`
	Description string `json:"description"`
}
