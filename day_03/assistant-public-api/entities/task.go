package entities

const (
	AttachFileByOwner    = 0
	AttachFileByExecutor = 1
)

type Task struct {
	ID              int64         `gorm:"column:id;primaryKey;"`
	Code            string        `gorm:"column:code;size:200"`
	Name            string        `gorm:"column:name;size:1024"`
	ProjectID       int64         `gorm:"column:project_id;"`
	Project         *Project      `gorm:"references:project_id;foreignKey:id;"`
	AreaID          int64         `gorm:"column:area_id;"`
	Area            *ProjectArea  `gorm:"references:area_id;foreignKey:id;"`
	PhaseID         int64         `gorm:"column:phase_id;"`
	Phase           *ProjectPhase `gorm:"references:phase_id;foreignKey:id;"`
	ParentTaskID    int64         `gorm:"column:parent_task_id;"`
	ParentTask      *Task
	ChildTasks      []*Task              `gorm:"references:id;foreignKey:parent_task_id;"`
	ExecutedBy      int64                `gorm:"column:executed_by;"`
	Executor        *Executor            `gorm:"references:executed_by;foreignKey:id;"`
	ReportBy        int64                `gorm:"column:report_by;"`
	Reporter        *User                `gorm:"references:report_by;foreignKey:id;"`
	AcceptedBy      int64                `gorm:"column:accepted_by;"`
	Acceptor        *User                `gorm:"references:accepted_by;foreignKey:id;"`
	StartTime       int64                `gorm:"column:start_time;"`
	EndTime         int64                `gorm:"column:end_time;"`
	Quantity        int64                `gorm:"column:quantity;"`
	Price           int64                `gorm:"column:price;"`
	Unit            string               `gorm:"column:unit;"`
	Description     string               `gorm:"column:description;"`
	Status          int                  `gorm:"column:status;"`
	PriorityLevel   int                  `gorm:"column:priority_level;"`
	AttachFiles     []*TaskAttachFile    `gorm:"references:id;foreignKey:task_id;"`
	Comments        []*TaskComment       `gorm:"references:id;foreignKey:task_id;"`
	AssignHistories []*TaskAssignHistory `gorm:"references:id;foreignKey:task_id;"`
	StatusHistories []*TaskStatusHistory `gorm:"references:id;foreignKey:task_id;"`
	Star            int                  `gorm:"column:star;"`
	Base
}

func (*Task) TableName() string {
	return "tasks"
}

func (t *Task) GetID() int64 {
	if t == nil {
		return 0
	}

	return t.ID
}

func (t *Task) GetName() string {
	if t == nil {
		return ""
	}

	return t.Name
}

func (t *Task) GetExecutor() *Executor {
	if t == nil {
		return nil
	}

	return t.Executor
}

func (t *Task) GetAcceptor() *User {
	if t == nil {
		return nil
	}

	return t.Acceptor
}

func (t *Task) GetReporter() *User {
	if t == nil {
		return nil
	}

	return t.Reporter
}

func (t *Task) GetProject() *Project {
	if t == nil {
		return nil
	}

	return t.Project
}

func (t *Task) GetArea() *ProjectArea {
	if t == nil {
		return nil
	}

	return t.Area
}

func (t *Task) GetPhase() *ProjectPhase {
	if t == nil {
		return nil
	}

	return t.Phase
}

func (t *Task) GetParentTask() *Task {
	if t == nil {
		return nil
	}

	return t.ParentTask
}

type TaskComment struct {
	ID      int64  `gorm:"column:id;primaryKey;"`
	TaskID  int64  `gorm:"column:task_id;"`
	Task    *Task  `gorm:"references:task_id;foreignKey:id;"`
	Comment string `gorm:"column:comment;"`
	Base
}

func (*TaskComment) TableName() string {
	return "task_comments"
}

func (c *TaskComment) GetTask() *Task {
	if c == nil {
		return nil
	}

	return c.Task
}

type TaskAttachFile struct {
	ID        int64  `gorm:"column:id;primaryKey;"`
	TaskID    int64  `gorm:"column:task_id;"`
	Task      *Task  `gorm:"references:task_id;foreignKey:id;"`
	Type      int    `gorm:"column:type;"`
	FileName  string `gorm:"column:file_name;"`
	Thumbnail string `gorm:"column:thumbnail;"`
	FilePath  string `gorm:"column:file_path;"`
	Base
}

func (*TaskAttachFile) TableName() string {
	return "task_attach_files"
}

type TaskAssignHistory struct {
	ID             int64  `gorm:"column:id;primaryKey;"`
	TaskID         int64  `gorm:"column:task_id;"`
	Task           *Task  `gorm:"references:task_id;foreignKey:id;"`
	AssignedUserID int64  `gorm:"column:assigned_user_id;"`
	Assigner       *User  `gorm:"references:assigned_user_id;foreignKey:id;"`
	Comment        string `gorm:"column:comment;"`
	Base
}

func (*TaskAssignHistory) TableName() string {
	return "task_assign_histories"
}

type TaskStatusHistory struct {
	ID     int64 `gorm:"column:id;primaryKey;"`
	TaskID int64 `gorm:"column:task_id;"`
	Task   *Task `gorm:"references:task_id;foreignKey:id;"`
	Status int   `gorm:"column:status;"`
	Base
}

func (*TaskStatusHistory) TableName() string {
	return "task_status_histories"
}
