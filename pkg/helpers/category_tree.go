package helpers

import "github.com/gogaruda/blog/blog/dto/response"

func BuildCategoryTree(flat []response.CategoryResponse) []*response.CategoryResponse {
	idMap := make(map[string]*response.CategoryResponse)
	var roots []*response.CategoryResponse

	for i := range flat {
		flat[i].Children = []*response.CategoryResponse{} // inisialisasi pointer slice
		idMap[flat[i].ID] = &flat[i]
	}

	for i := range flat {
		node := &flat[i]
		if node.ParentID != nil && *node.ParentID != "" {
			if parent, ok := idMap[*node.ParentID]; ok {
				parent.Children = append(parent.Children, node)
			}
		} else {
			roots = append(roots, node)
		}
	}

	return roots
}
