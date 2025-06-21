package repository

import (
	"database/sql"
	"github.com/gogaruda/blog/blog/model"
	"github.com/gogaruda/pkg/apperror"
)

type TagRepository interface {
	GetAll() ([]model.TagModel, error)
}

type tagRepository struct {
	db *sql.DB
}

func NewTagRepository(db *sql.DB) TagRepository {
	return &tagRepository{db}
}

func (r *tagRepository) GetAll() ([]model.TagModel, error) {
	query := `SELECT id, name, slug, seo_title, seo_description FROM tags ORDER BY created_at DESC`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, apperror.New(apperror.CodeDBError, "gagal query GetAll", err)
	}
	defer rows.Close()

	var tags []model.TagModel
	for rows.Next() {
		var tag model.TagModel
		if err := rows.Scan(&tag.ID, &tag.Name, &tag.Slug, &tag.SeoTitle, &tag.SeoDescription); err != nil {
			return nil, apperror.New(apperror.CodeDBError, "gagal scan GetAll", err)
		}
		tags = append(tags, tag)
	}

	if err := rows.Err(); err != nil {
		return nil, apperror.New(apperror.CodeDBError, "gagal setelah iterasi rows", err)
	}

	return tags, nil
}
