package model

import "time"

type TagModel struct {
	ID             string
	Name           string
	Slug           string
	SeoTitle       string
	SeoDescription string
	CreatedAt      time.Time
}
