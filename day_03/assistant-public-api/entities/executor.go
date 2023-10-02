package entities

type Executor struct {
	ID            int64   `gorm:"column:id;primaryKey"`
	Code          string  `gorm:"column:code;size:200"`
	RepresentedBy int64   `gorm:"column:represented_by;"`
	Representer   *User   `gorm:"references:represented_by;foreignKey:id;"`
	Name          string  `gorm:"column:name;size:200"`
	Address       string  `gorm:"column:address;size:200"`
	Phone         string  `gorm:"column:phone;size:200"`
	Email         string  `gorm:"column:email;size:200"`
	Tasks         []*Task `gorm:"references:id;foreignKey:executed_by;"`
	Base
}

func (*Executor) TableName() string {
	return "executors"
}

func (e *Executor) GetID() int64 {
	if e == nil {
		return 0
	}

	return e.ID
}

func (e *Executor) GetCode() string {
	if e == nil {
		return ""
	}

	return e.Code
}

func (e *Executor) GetName() string {
	if e == nil {
		return ""
	}

	return e.Name
}

func (e *Executor) GetPhone() string {
	if e == nil {
		return ""
	}

	return e.Phone
}

func (e *Executor) GetEmail() string {
	if e == nil {
		return ""
	}

	return e.Email
}

func (e *Executor) GetAddress() string {
	if e == nil {
		return ""
	}

	return e.Address
}

func (e *Executor) GetRepresenter() *User {
	if e == nil {
		return nil
	}

	return e.Representer
}
