package request

type CategoryRequest struct {
	Name           string  `json:"name" binding:"required"`
	Description    string  `json:"description" binding:"required"`
	ParentID       *string `json:"parent_id,omitempty"`
	SeoDescription string  `json:"seo_description,omitempty"`
}
