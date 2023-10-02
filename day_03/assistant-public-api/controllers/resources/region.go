package resources

type CreateRegionRequest struct {
	Name string `json:"name"`
}

type UpdateRegionRequest struct {
	Name string `json:"name"`
}

type ListRegionResponse struct {
	Regions []*Region `json:"regions"`
}
