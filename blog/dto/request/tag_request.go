package request

type TagRequest struct {
	Name           string `json:"name" binding:"required"`
	Description    string `json:"description" binding:"required"`
	SeoDescription string `json:"seo_description,omitempty"`
}
