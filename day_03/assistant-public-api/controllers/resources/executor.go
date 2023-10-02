package resources

type CreateExecutorRequest struct {
	Code          string `json:"code"`
	Name          string `json:"name"`
	RepresenterID int64  `json:"representer_id"`
	Address       string `json:"address"`
	Phone         string `json:"phone"`
	Email         string `json:"email"`
}

type UpdateExecutorRequest struct {
	Code          string `json:"code"`
	Name          string `json:"name"`
	RepresenterID int64  `json:"representer_id"`
	Address       string `json:"address"`
	Phone         string `json:"phone"`
	Email         string `json:"email"`
}

type ListExecutorResponse struct {
	Executors []*Executor `json:"executors"`
}
