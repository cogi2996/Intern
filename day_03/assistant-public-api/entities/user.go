package entities

type User struct {
	Base
	ID        int64  `gorm:"column:id;primaryKey"`
	UUID      string `gorm:"column:uuid;size:200"`
	Code      string `gorm:"column:code;size:200"`
	Name      string `gorm:"column:name;size:200"`
	Phone     string `gorm:"column:phone;size:200"`
	Email     string `gorm:"column:email;size:200"`
	Address   string `gorm:"column:address;size:200"`
	ManagerID int64  `gorm:"column:manager_id"`
	Manager   *User
	Members   []*User     `gorm:"references:id;foreignKey:manager_id;"`
	Executors []*Executor `gorm:"references:id;foreignKey:represented_by;"`
}

func (*User) TableName() string {
	return "users"
}

func (u *User) GetID() int64 {
	if u == nil {
		return 0
	}

	return u.ID
}

func (u *User) GetName() string {
	if u == nil {
		return ""
	}

	return u.Name
}

func (u *User) GetCode() string {
	if u == nil {
		return ""
	}

	return u.Code
}

func (u *User) GetPhone() string {
	if u == nil {
		return ""
	}

	return u.Phone
}

func (u *User) GetEmail() string {
	if u == nil {
		return ""
	}

	return u.Email
}

func (u *User) GetAddress() string {
	if u == nil {
		return ""
	}

	return u.Address
}

func (u *User) GetManager() *User {
	if u == nil {
		return nil
	}

	return u.Manager
}

func (u *User) GetMembers() []*User {
	if u == nil {
		return nil
	}

	return u.Members
}
