package resources

type CreateProjectCategoryRequest struct {
	Name string `json:"name"`
}

type UpdateProjectCategoryRequest struct {
	Name string `json:"name"`
}

type ListProjectCategoryResponse struct {
	Categories []*ProjectCategory `json:"categories"`
}
