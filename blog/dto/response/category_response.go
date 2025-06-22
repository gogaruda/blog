package response

type CategoryModel struct {
	ID             string
	Name           string
	Slug           string
	Description    string
	ParentID       *string
	SeoTitle       string
	SeoDescription string
}
