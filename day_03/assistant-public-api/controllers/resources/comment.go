package resources

type CreateTaskCommentRequest struct {
	TaskID int64  `json:"task_id,omitempty"`
	Msg    string `json:"msg,omitempty"`
}

type ListTaskCommentResponse struct {
	Comments []*Comment `json:"comments"`
}
