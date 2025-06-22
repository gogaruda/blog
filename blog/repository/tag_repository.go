package repository

import (
	"database/sql"
	"github.com/gogaruda/blog/blog/dto/request"
	"github.com/gogaruda/blog/blog/dto/response"
	"github.com/gogaruda/pkg/apperror"
	"github.com/gogaruda/pkg/utils"
	"github.com/gogaruda/seo"
)

type TagRepository interface {
	GetAll(limit, offset int) ([]response.TagResponse, int, error)
	CreateSlug(slug string) (string, error)
	UpdateSlug(input string, excludeID string) (string, error)
	Create(req request.TagRequest, slug, seoDescription string) error
	GetByID(tagID string) (*response.TagResponse, error)
	Update(req request.TagRequest, slug, tagID, seoDescription string) error
	Delete(tagID string) error
}

type tagRepository struct {
	db *sql.DB
}

func NewTagRepository(db *sql.DB) TagRepository {
	return &tagRepository{db}
}

func (r *tagRepository) GetAll(limit, offset int) ([]response.TagResponse, int, error) {
	query := `SELECT id, name, slug, description, seo_title, seo_description 
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

	var tags []response.TagResponse
	for rows.Next() {
		var tag response.TagResponse
		if err := rows.Scan(&tag.ID, &tag.Name, &tag.Slug, &tag.Description, &tag.SeoTitle, &tag.SeoDescription); err != nil {
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
	base := seo.Slugify(input)

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

	if err := rows.Err(); err != nil {
		return "", apperror.New(apperror.CodeDBError, "gagal setelah iterasi rows", err)
	}

	uniqueSlug := seo.GenerateUniqueSlug(input, existing)

	return uniqueSlug, nil
}

func (r *tagRepository) UpdateSlug(input string, excludeID string) (string, error) {
	base := seo.Slugify(input)

	// Ambil semua slug yang mirip, kecuali milik ID yang sedang diupdate
	rows, err := r.db.Query(`
		SELECT slug FROM tags 
		WHERE slug LIKE ? AND id != ?
	`, base+"%", excludeID)
	if err != nil {
		return "", apperror.New(apperror.CodeDBError, "gagal query slug saat update", err)
	}
	defer rows.Close()

	var existing []string
	for rows.Next() {
		var slug string
		_ = rows.Scan(&slug)
		existing = append(existing, slug)
	}
	if err := rows.Err(); err != nil {
		return "", apperror.New(apperror.CodeDBError, "error iterasi slug", err)
	}

	uniqueSlug := seo.GenerateUniqueSlug(input, existing)
	return uniqueSlug, nil
}

func (r *tagRepository) Create(req request.TagRequest, slug, seoDescription string) error {
	seoTitle := seo.Title(req.Name, 60)

	_, err := r.db.Exec(`INSERT INTO 
										tags(id, name, slug, description, seo_title, seo_description)
										VALUES(?, ?, ?, ?, ?, ?)`,
		utils.NewULID(), req.Name, slug, req.Description, seoTitle, seoDescription)

	if err != nil {
		return apperror.New(apperror.CodeDBError, "gagal query create tag", err)
	}

	return nil
}

func (r *tagRepository) GetByID(tagID string) (*response.TagResponse, error) {
	var tag response.TagResponse
	err := r.db.QueryRow(`SELECT id, name, slug, description, seo_title, seo_description FROM tags WHERE id = ?`, tagID).
		Scan(&tag.ID, &tag.Name, &tag.Slug, &tag.Description, &tag.SeoTitle, &tag.SeoDescription)

	if err != nil {
		return nil, apperror.New(apperror.CodeResourceNotFound, "query tag by id gagal", err)
	}

	return &tag, nil
}

func (r *tagRepository) Update(req request.TagRequest, slug, tagID, seoDescription string) error {
	seoTitle := seo.Title(req.Name, 60)

	_, err := r.db.Exec(`UPDATE tags SET name = ?, slug = ?, description = ?, seo_title = ?, seo_description = ? WHERE id = ?`,
		req.Name, slug, req.Description, seoTitle, seoDescription, tagID)
	if err != nil {
		return apperror.New(apperror.CodeDBError, "query update tag gagal", err)
	}

	return nil
}

func (r *tagRepository) Delete(tagID string) error {
	_, err := r.db.Exec(`DELETE FROM tags WHERE id = ?`, tagID)

	if err != nil {
		return apperror.New(apperror.CodeDBError, "query delete tag gagal", err)
	}

	return nil
}
