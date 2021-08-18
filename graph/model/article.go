package model

import "time"

type Article struct {
	ID          string    `json:"id"`
	Slug        string    `json:"slug"`
	Title       string    `json:"title"`
	Date        time.Time `json:"date"`
	CoverImage  string    `json:"coverImage"`
	ContentID   string    `json:"content"`
	Description string    `json:"description"`
}