package entities

type AuditTask struct {
	ID              int64    `gorm:"column:id;primaryKey;"`
	Code            string   `gorm:"column:code;size:200"`
	Name            string   `gorm:"column:name;size:1024"`
	ProjectID       int64    `gorm:"column:project_id;"`
	Project         *Project `gorm:"references:project_id;foreignKey:id;"`
	ParentTaskID    int64    `gorm:"column:parent_task_id;"`
	ParentTask      *AuditTask
	ChildTasks      []*AuditTask              `gorm:"references:id;foreignKey:parent_task_id;"`
	ExecutedBy      int64                     `gorm:"column:executed_by;"`
	Executor        *User                     `gorm:"references:executed_by;foreignKey:id;"`
	AcceptedBy      int64                     `gorm:"column:accepted_by;"`
	Acceptor        *User                     `gorm:"references:accepted_by;foreignKey:id;"`
	StartTime       int64                     `gorm:"column:start_time;"`
	EndTime         int64                     `gorm:"column:end_time;"`
	Description     string                    `gorm:"column:description;"`
	Status          int                       `gorm:"column:status;"`
	PriorityLevel   int                       `gorm:"column:priority_level;"`
	AttachFiles     []*AuditTaskAttachFile    `gorm:"references:id;foreignKey:task_id;"`
	Comments        []*AuditTaskComment       `gorm:"references:id;foreignKey:task_id;"`
	AssignHistories []*AuditTaskAssignHistory `gorm:"references:id;foreignKey:task_id;"`
	StatusHistories []*AuditTaskStatusHistory `gorm:"references:id;foreignKey:task_id;"`
	Base
}

func (*AuditTask) TableName() string {
	return "audit_tasks"
}

func (t *AuditTask) GetID() int64 {
	if t == nil {
		return 0
	}

	return t.ID
}

func (t *AuditTask) GetName() string {
	if t == nil {
		return ""
	}

	return t.Name
}

func (t *AuditTask) GetExecutor() *User {
	if t == nil {
		return nil
	}

	return t.Executor
}

func (t *AuditTask) GetAcceptor() *User {
	if t == nil {
		return nil
	}

	return t.Acceptor
}

func (t *AuditTask) GetProject() *Project {
	if t == nil {
		return nil
	}

	return t.Project
}

func (t *AuditTask) GetParentTask() *AuditTask {
	if t == nil {
		return nil
	}

	return t.ParentTask
}

type AuditTaskComment struct {
	ID      int64      `gorm:"column:id;primaryKey;"`
	TaskID  int64      `gorm:"column:task_id;"`
	Task    *AuditTask `gorm:"references:task_id;foreignKey:id;"`
	Comment string     `gorm:"column:comment;"`
	Base
}

func (*AuditTaskComment) TableName() string {
	return "audit_task_comments"
}

func (c *AuditTaskComment) GetTask() *AuditTask {
	if c == nil {
		return nil
	}

	return c.Task
}

type AuditTaskAttachFile struct {
	ID        int64      `gorm:"column:id;primaryKey;"`
	TaskID    int64      `gorm:"column:task_id;"`
	Task      *AuditTask `gorm:"references:task_id;foreignKey:id;"`
	Type      int        `gorm:"column:type;"`
	FileName  string     `gorm:"column:file_name;"`
	Thumbnail string     `gorm:"column:thumbnail;"`
	FilePath  string     `gorm:"column:file_path;"`
	Base
}

func (*AuditTaskAttachFile) TableName() string {
	return "audit_task_attach_files"
}

type AuditTaskAssignHistory struct {
	ID             int64  `gorm:"column:id;primaryKey;"`
	TaskID         int64  `gorm:"column:task_id;"`
	Task           *Task  `gorm:"references:task_id;foreignKey:id;"`
	AssignedUserID int64  `gorm:"column:assigned_user_id;"`
	Assigner       *User  `gorm:"references:assigned_user_id;foreignKey:id;"`
	Comment        string `gorm:"column:comment;"`
	Base
}

func (*AuditTaskAssignHistory) TableName() string {
	return "audit_task_assign_histories"
}

type AuditTaskStatusHistory struct {
	ID     int64      `gorm:"column:id;primaryKey;"`
	TaskID int64      `gorm:"column:task_id;"`
	Task   *AuditTask `gorm:"references:task_id;foreignKey:id;"`
	Status int        `gorm:"column:status;"`
	Base
}

func (*AuditTaskStatusHistory) TableName() string {
	return "audit_task_status_histories"
}
