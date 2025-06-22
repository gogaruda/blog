package repository

import (
	"database/sql"
	"github.com/gogaruda/blog/blog/dto/request"
	"github.com/gogaruda/blog/blog/dto/response"
	"github.com/gogaruda/blog/pkg/helpers"
	"github.com/gogaruda/pkg/apperror"
	"github.com/gogaruda/pkg/utils"
	"github.com/gogaruda/seo"
)

type CategoryRepository interface {
	GetAll() ([]*response.CategoryResponse, error)
	CreateSlug(input string) (string, error)
	UpdateSlug(input string, excludeID string) (string, error)
	Create(req request.CategoryRequest, slug, seoDescription string) error
}

type categoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) CategoryRepository {
	return &categoryRepository{db}
}

func (r *categoryRepository) GetAll() ([]*response.CategoryResponse, error) {
	query := `SELECT id, name, slug, description, parent_id, seo_title, seo_description
            FROM categories ORDER BY updated_at DESC`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, apperror.New(apperror.CodeDBError, "query cateogires gagal", err)
	}
	defer rows.Close()

	var flat []response.CategoryResponse
	for rows.Next() {
		var category response.CategoryResponse
		if err := rows.Scan(
			&category.ID, &category.Name, &category.Slug,
			&category.Description, &category.ParentID,
			&category.SeoTitle, &category.SeoDescription,
		); err != nil {
			return nil, apperror.New(apperror.CodeDBError, "gagal scan categories", err)
		}
		category.Children = []*response.CategoryResponse{} // inisialisasi children dulu
		flat = append(flat, category)
	}

	if err := rows.Err(); err != nil {
		return nil, apperror.New(apperror.CodeDBError, "gagal terasi rows", err)
	}

	return helpers.BuildCategoryTree(flat), nil
}

func (r *categoryRepository) CreateSlug(input string) (string, error) {
	base := seo.Slugify(input)

	rows, err := r.db.Query(`SELECT slug FROM categories WHERE slug LIKE ?`, base+"%")
	if err != nil {
		return "", apperror.New(apperror.CodeDBError, "gagal query category slug", err)
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

func (r *categoryRepository) UpdateSlug(input string, excludeID string) (string, error) {
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

func (r *categoryRepository) Create(req request.CategoryRequest, slug, seoDescription string) error {
	seoTitle := seo.Title(req.Name, 60)
	_, err := r.db.Exec(`INSERT INTO categories(id, name, slug, description, parent_id, seo_title, seo_description)
													VALUES(?, ?, ?, ?, ?, ?, ?)`,
		utils.NewULID(), req.Name, slug, req.Description, req.ParentID, seoTitle, seoDescription,
	)
	if err != nil {
		return apperror.New(apperror.CodeDBError, "query insert category gagal", err)
	}

	return nil
}
