package repository

import (
	"database/sql"
	"github.com/gogaruda/blog/blog/dto/request"
	"github.com/gogaruda/blog/blog/model"
	"github.com/gogaruda/pkg/apperror"
	"github.com/gogaruda/pkg/utils"
	"github.com/gogaruda/seo/pkg"
)

type TagRepository interface {
	GetAll(limit, offset int) ([]model.TagModel, int, error)
	CreateSlug(slug string) (string, error)
	Create(req request.TagRequest, slug string) error
}

type tagRepository struct {
	db *sql.DB
}

func NewTagRepository(db *sql.DB) TagRepository {
	return &tagRepository{db}
}

func (r *tagRepository) GetAll(limit, offset int) ([]model.TagModel, int, error) {
	query := `SELECT id, name, slug, seo_title, seo_description 
						FROM tags ORDER BY created_at DESC
						LIMIT ? OFFSET ?`
	rows, err := r.db.Query(query, limit, offset)
	if err != nil {
		return nil, 0, apperror.New(apperror.CodeDBError, "gagal query GetAll", err)
	}
	defer rows.Close()

	queryCount := `SELECT COUNT(*) FROM tags`
	var total int
	err = r.db.QueryRow(queryCount).Scan(&total)
	if err != nil {
		return nil, 0, apperror.New(apperror.CodeDBError, "gagal query count tags", err)
	}

	var tags []model.TagModel
	for rows.Next() {
		var tag model.TagModel
		if err := rows.Scan(&tag.ID, &tag.Name, &tag.Slug, &tag.SeoTitle, &tag.SeoDescription); err != nil {
			return nil, 0, apperror.New(apperror.CodeDBError, "gagal scan GetAll", err)
		}
		tags = append(tags, tag)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, apperror.New(apperror.CodeDBError, "gagal setelah iterasi rows", err)
	}

	return tags, total, nil
}

func (r *tagRepository) CreateSlug(input string) (string, error) {
	base := pkg.Slugify(input)

	rows, err := r.db.Query(`SELECT slug FROM tags WHERE slug LIKE ?`, base+"%")
	if err != nil {
		return "", apperror.New(apperror.CodeDBError, "gagal query tag slug", err)
	}
	defer rows.Close()

	var existing []string
	for rows.Next() {
		var slug string
		_ = rows.Scan(&slug)
		existing = append(existing, slug)
	}

	uniqueSlug := pkg.GenerateUniqueSlug(input, existing)

	return uniqueSlug, nil
}

func (r *tagRepository) Create(req request.TagRequest, slug string) error {
	seoTitle := pkg.Title(req.Name, 60)
	seoDescription := pkg.Description(req.Description, 160)
	_, err := r.db.Exec(`INSERT INTO 
										tags(id, name, slug, description, seo_title, seo_description)
										VALUES(?, ?, ?, ?, ?, ?)`,
		utils.NewULID(), req.Name, slug, req.Description, seoTitle, seoDescription)

	if err != nil {
		return apperror.New(apperror.CodeDBError, "gagal query create tag", err)
	}

	return nil
}
