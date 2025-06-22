package response

type TagResponse struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Slug           string `json:"slug"`
	Description    string `json:"description"`
	SeoTitle       string `json:"seo_title"`
	SeoDescription string `json:"seo_description"`
}
