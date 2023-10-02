package resources

type CreateTaskImageRequest struct {
	IsOwner bool `json:"is_owner,omitempty" form:"is_owner"`
}

type CreateTaskImageResponse struct {
	ID        int64  `json:"id,omitempty"`
	UUID      string `json:"uuid,omitempty"`
	FilePath  string `json:"file_path,omitempty"`
	Thumbnail string `json:"thumbnail,omitempty"`
}

type ListTaskImageResponse struct {
	OwnerImages    []*Image `json:"owner_images,omitempty"`
	ExecutorImages []*Image `json:"executor_images,omitempty"`
}
