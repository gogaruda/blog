package model

import "time"

type CategoryModel struct {
	ID             string
	Name           string
	Slug           string
	Description    string
	ParentID       *string
	SeoTitle       string
	SeoDescription string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
