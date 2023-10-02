package resources

// Create
type CreateProjectRequest struct {
	Name       string `json:"name"`
	Code       string `json:"code"`
	ManagerID  int64  `json:"manager_id"`
	Address    string `json:"address"`
	StartDate  string `json:"start_date"`
	EndDate    string `json:"end_date"`
	RegionID   int64  `json:"region_id"`
	CategoryID int64  `json:"category_id"`
}

type UpdateProjectRequest struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Code       string `json:"code"`
	ManagerID  int64  `json:"manager_id"`
	Address    string `json:"address"`
	StartDate  string `json:"start_date"`
	EndDate    string `json:"end_date"`
	RegionID   int64  `json:"region_id"`
	CategoryID int64  `json:"category_id"`
}

type ListProjectResponse struct {
	Projects []*Project `json:"projects"`
}

type AddProjectMemberRequest struct {
	UserID int64 `json:"user_id"`
}

type RemoveProjectMemberRequest struct {
	UserID int64 `json:"user_id"`
}

type AddProjectExecutorRequest struct {
	ExecutorID int64 `json:"executor_id"`
}

type RemoveProjectExecutorRequest struct {
	ExecutorID int64 `json:"executor_id"`
}

type ListProjectMemberResponse struct {
	Users []*User `json:"users"`
}

type ListProjectExecutorResponse struct {
	Executors []*Executor `json:"executors"`
}

type CreateProjectPhaseRequest struct {
	Name      string `json:"name"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type UpdateProjectPhaseRequest struct {
	Name      string `json:"name"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type ListProjectPhaseResponse struct {
	Phases []*ProjectPhase `json:"phases"`
}

type CreateProjectAreaRequest struct {
	Name string `json:"name"`
}

type UpdateProjectAreaRequest struct {
	Name string `json:"name"`
}

type ListProjectAreaResponse struct {
	Areas []*ProjectArea `json:"areas"`
}
