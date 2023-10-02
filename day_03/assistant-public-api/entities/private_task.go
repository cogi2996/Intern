package entities

type PrivateTask struct {
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
	ParentTask      *PrivateTask
	ChildTasks      []*PrivateTask              `gorm:"references:id;foreignKey:parent_task_id;"`
	ExecutedBy      int64                       `gorm:"column:executed_by;"`
	Executor        *User                       `gorm:"references:executed_by;foreignKey:id;"`
	AcceptedBy      int64                       `gorm:"column:accepted_by;"`
	Acceptor        *User                       `gorm:"references:accepted_by;foreignKey:id;"`
	StartTime       int64                       `gorm:"column:start_time;"`
	EndTime         int64                       `gorm:"column:end_time;"`
	Quantity        int64                       `gorm:"column:quantity;"`
	Price           int64                       `gorm:"column:price;"`
	Unit            string                      `gorm:"column:unit;"`
	Description     string                      `gorm:"column:description;"`
	Status          int                         `gorm:"column:status;"`
	PriorityLevel   int                         `gorm:"column:priority_level;"`
	AttachFiles     []*PrivateTaskAttachFile    `gorm:"references:id;foreignKey:task_id;"`
	Comments        []*PrivateTaskComment       `gorm:"references:id;foreignKey:task_id;"`
	AssignHistories []*PrivateTaskAssignHistory `gorm:"references:id;foreignKey:task_id;"`
	StatusHistories []*PrivateTaskStatusHistory `gorm:"references:id;foreignKey:task_id;"`
	Base
}

func (*PrivateTask) TableName() string {
	return "private_tasks"
}

func (t *PrivateTask) GetID() int64 {
	if t == nil {
		return 0
	}

	return t.ID
}

func (t *PrivateTask) GetName() string {
	if t == nil {
		return ""
	}

	return t.Name
}

func (t *PrivateTask) GetExecutor() *User {
	if t == nil {
		return nil
	}

	return t.Executor
}

func (t *PrivateTask) GetAcceptor() *User {
	if t == nil {
		return nil
	}

	return t.Acceptor
}

func (t *PrivateTask) GetProject() *Project {
	if t == nil {
		return nil
	}

	return t.Project
}

func (t *PrivateTask) GetArea() *ProjectArea {
	if t == nil {
		return nil
	}

	return t.Area
}

func (t *PrivateTask) GetPhase() *ProjectPhase {
	if t == nil {
		return nil
	}

	return t.Phase
}

func (t *PrivateTask) GetParentTask() *PrivateTask {
	if t == nil {
		return nil
	}

	return t.ParentTask
}

type PrivateTaskComment struct {
	ID      int64        `gorm:"column:id;primaryKey;"`
	TaskID  int64        `gorm:"column:task_id;"`
	Task    *PrivateTask `gorm:"references:task_id;foreignKey:id;"`
	Comment string       `gorm:"column:comment;"`
	Base
}

func (*PrivateTaskComment) TableName() string {
	return "private_task_comments"
}

func (c *PrivateTaskComment) GetTask() *PrivateTask {
	if c == nil {
		return nil
	}

	return c.Task
}

type PrivateTaskAttachFile struct {
	ID        int64        `gorm:"column:id;primaryKey;"`
	TaskID    int64        `gorm:"column:task_id;"`
	Task      *PrivateTask `gorm:"references:task_id;foreignKey:id;"`
	Type      int          `gorm:"column:type;"`
	FileName  string       `gorm:"column:file_name;"`
	Thumbnail string       `gorm:"column:thumbnail;"`
	FilePath  string       `gorm:"column:file_path;"`
	Base
}

func (*PrivateTaskAttachFile) TableName() string {
	return "private_task_attach_files"
}

type PrivateTaskAssignHistory struct {
	ID             int64  `gorm:"column:id;primaryKey;"`
	TaskID         int64  `gorm:"column:task_id;"`
	Task           *Task  `gorm:"references:task_id;foreignKey:id;"`
	AssignedUserID int64  `gorm:"column:assigned_user_id;"`
	Assigner       *User  `gorm:"references:assigned_user_id;foreignKey:id;"`
	Comment        string `gorm:"column:comment;"`
	Base
}

func (*PrivateTaskAssignHistory) TableName() string {
	return "private_task_assign_histories"
}

type PrivateTaskStatusHistory struct {
	ID     int64        `gorm:"column:id;primaryKey;"`
	TaskID int64        `gorm:"column:task_id;"`
	Task   *PrivateTask `gorm:"references:task_id;foreignKey:id;"`
	Status int          `gorm:"column:status;"`
	Base
}

func (*PrivateTaskStatusHistory) TableName() string {
	return "private_task_status_histories"
}
