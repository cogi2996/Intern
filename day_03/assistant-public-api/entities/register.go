package entities

type Register struct {
	ID       int64  `gorm:"column:id;primaryKey"`
	Username string `gorm:"column:username;size:200"`
	Password string `gorm:"column:password;size:200"`
	Status   int    `gorm:"column:status;size:200"`
	Base
}

func (*Register) TableName() string {
	return "registers"
}

func (r *Register) GetID() int64 {
	if r == nil {
		return 0
	}

	return r.ID
}

func (r *Register) GetPassword() string {
	if r == nil {
		return ""
	}

	return r.Password
}
