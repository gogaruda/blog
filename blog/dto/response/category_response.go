package response

type CategoryResponse struct {
	ID             string              `json:"id"`
	Name           string              `json:"name"`
	Slug           string              `json:"slug"`
	Description    string              `json:"description"`
	ParentID       *string             `json:"parent_id"`
	SeoTitle       string              `json:"seo_title"`
	SeoDescription string              `json:"seo_description"`
	Children       []*CategoryResponse `json:"children"`
}
