package resources

// Create
type CreateTaskRequest struct {
	Name          string `json:"name"`
	ParentTaskID  int64  `json:"parent_task_id"`
	ProjectID     int64  `json:"project_id"`
	AreaID        int64  `json:"area_id"`
	PhaseID       int64  `json:"phase_id"`
	ExecutorID    int64  `json:"executor_id"`
	ReporterID    int64  `json:"reporter_id"`
	AcceptorID    int64  `json:"acceptor_id"`
	StartTime     string `json:"start_time"`
	EndTime       string `json:"end_time"`
	Quantity      int64  `json:"quantity"`
	Price         int64  `json:"price"`
	Unit          string `json:"unit"`
	Description   string `json:"description"`
	PriorityLevel string `json:"priority_level"`
}

type CreateTaskResponse struct {
	ID   int64  `json:"id"`
	UUID string `json:"uuid"`
}

type UpdateTaskRequest struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	ProjectID     int64  `json:"project_id"`
	AreaID        int64  `json:"area_id"`
	PhaseID       int64  `json:"phase_id"`
	ExecutorID    int64  `json:"executor_id"`
	ReporterID    int64  `json:"reporter_id"`
	AcceptorID    int64  `json:"acceptor_id"`
	StartTime     string `json:"start_time"`
	EndTime       string `json:"end_time"`
	Quantity      int64  `json:"quantity"`
	Price         int64  `json:"price"`
	Unit          string `json:"unit"`
	Description   string `json:"description"`
	Status        int    `json:"status"`
	PriorityLevel string `json:"priority_level"`
}

type UpdateTaskStatusRequest struct {
	Status string `json:"status"`
	Star   int    `json:"star"`
}

type ListTaskRequest struct {
	ProjectID      int64  `json:"project_id" form:"project_id"`
	AreaID         int64  `json:"area_id"`
	PhaseID        int64  `json:"phase_id"`
	CreatorID      int64  `json:"creator_id" form:"creator_id"`
	AcceptorID     int64  `json:"acceptor_id" form:"acceptor_id"`
	ExecutorID     int64  `json:"executor_id" form:"executor_id"`
	StartTime      string `json:"start_time" form:"start_time"`
	EndTime        string `json:"end_time" form:"end_time"`
	Status         string `json:"status" form:"status"`
	PriorityLevel  string `json:"priority_level" form:"priority_level"`
	NeedExportFile bool   `json:"need_export_file" form:"need_export_file"`
}

type ListTaskResponse struct {
	Tasks    []*Task `json:"tasks"`
	ExcelURL string  `json:"excel_url"`
}

type ListAssigningTaskResponse struct {
	Tasks []*AssigningTask `json:"tasks"`
}

type DeleteTaskRequest struct {
	Force bool `json:"force" form:"force"`
}

type RemindTaskRequest struct {
	Message string `json:"message"`
}
