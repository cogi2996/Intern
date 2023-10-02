package entities

type Notification struct {
	Base
	ID           int64  `gorm:"column:id;primaryKey"`
	UUID         string `gorm:"column:uuid;size:200"`
	Title        string `gorm:"column:code;size:200"`
	Message      string `gorm:"column:name;size:200"`
	PublicTaskID int64  `gorm:"column:public_task_id;"`
	PublicTask   *Task  `gorm:"references:public_task_id;foreignKey:id;"`
	ReceivedBy   int64  `gorm:"column:received_by;"`
	Receiver     *User  `gorm:"references:received_by;foreignKey:id;"`
}

func (n *Notification) GetTitle() string {
	if n == nil {
		return ""
	}
	return n.Title
}

func (n *Notification) GetMessage() string {
	if n == nil {
		return ""
	}
	return n.Message
}

func (n *Notification) GetPublicTask() *Task {
	if n == nil {
		return nil
	}
	return n.PublicTask
}
