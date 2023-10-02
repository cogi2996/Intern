package entities

type Region struct {
	ID   int64  `gorm:"column:id;primaryKey"`
	Name string `gorm:"column:name"`
	Base
}

func (*Region) TableName() string {
	return "regions"
}

func (p *Region) GetID() int64 {
	if p == nil {
		return 0
	}

	return p.ID
}

func (p *Region) GetName() string {
	if p == nil {
		return ""
	}

	return p.Name
}

type ProjectCategory struct {
	ID   int64  `gorm:"column:id;primaryKey"`
	Name string `gorm:"column:name"`
	Base
}

func (*ProjectCategory) TableName() string {
	return "project_categories"
}

func (p *ProjectCategory) GetID() int64 {
	if p == nil {
		return 0
	}

	return p.ID
}

func (p *ProjectCategory) GetName() string {
	if p == nil {
		return ""
	}

	return p.Name
}

type Project struct {
	ID         int64              `gorm:"column:id;primaryKey"`
	Code       string             `gorm:"column:code;size:200"`
	Name       string             `gorm:"column:name;size:200"`
	Address    string             `gorm:"column:address;size:200"`
	StartDate  int64              `gorm:"column:start_date"`
	EndDate    int64              `gorm:"column:end_date"`
	RegionID   int64              `gorm:"column:region_id"`
	Region     *Region            `gorm:"references:region_id;foreignKey:id;"`
	CategoryID int64              `gorm:"column:category_id"`
	Category   *ProjectCategory   `gorm:"references:category_id;foreignKey:id;"`
	ManagerID  int64              `gorm:"column:manager_id;"`
	Manager    *User              `gorm:"references:manager_id;foreignKey:id;"`
	Phases     []*ProjectPhase    `gorm:"references:id;foreignKey:project_id;"`
	Areas      []*ProjectArea     `gorm:"references:id;foreignKey:project_id;"`
	Tasks      []*Task            `gorm:"references:id;foreignKey:project_id;"`
	Members    []*ProjectMember   `gorm:"references:id;foreignKey:project_id;"`
	Executors  []*ProjectExecutor `gorm:"references:id;foreignKey:project_id;"`
	Status     int                `gorm:"column:status;"`
	Base
}

func (*Project) TableName() string {
	return "projects"
}

func (p *Project) GetID() int64 {
	if p == nil {
		return 0
	}

	return p.ID
}

func (p *Project) GetName() string {
	if p == nil {
		return ""
	}

	return p.Name
}

func (p *Project) GetCode() string {
	if p == nil {
		return ""
	}

	return p.Code
}

func (p *Project) GetAddress() string {
	if p == nil {
		return ""
	}

	return p.Address
}

func (p *Project) GetStartDate() int64 {
	if p == nil {
		return 0
	}

	return p.StartDate
}

func (p *Project) GetEndDate() int64 {
	if p == nil {
		return 0
	}

	return p.EndDate
}

func (p *Project) GetManager() *User {
	if p == nil {
		return nil
	}

	return p.Manager
}

func (p *Project) GetCategory() *ProjectCategory {
	if p == nil {
		return nil
	}

	return p.Category
}

func (p *Project) GetRegion() *Region {
	if p == nil {
		return nil
	}

	return p.Region
}

type ProjectMember struct {
	ProjectID int64    `gorm:"column:project_id;primaryKey"`
	Project   *Project `gorm:"references:project_id;foreignKey:id;"`
	MemberID  int64    `gorm:"column:member_id;primaryKey"`
	Member    *User    `gorm:"references:member_id;foreignKey:id;"`
	Base
}

func (*ProjectMember) TableName() string {
	return "project_members"
}

func (p *ProjectMember) GetProject() *Project {
	if p == nil {
		return nil
	}

	return p.Project
}

func (p *ProjectMember) GetMember() *User {
	if p == nil {
		return nil
	}

	return p.Member
}

type ProjectExecutor struct {
	ProjectID  int64     `gorm:"column:project_id;primaryKey"`
	Project    *Project  `gorm:"references:project_id;foreignKey:id;"`
	ExecutorID int64     `gorm:"column:executor_id;primaryKey"`
	Executor   *Executor `gorm:"references:executor_id;foreignKey:id;"`
	Base
}

func (*ProjectExecutor) TableName() string {
	return "project_executors"
}

func (p *ProjectExecutor) GetProject() *Project {
	if p == nil {
		return nil
	}

	return p.Project
}

func (p *ProjectExecutor) GetExecutor() *Executor {
	if p == nil {
		return nil
	}

	return p.Executor
}

type ProjectPhase struct {
	ID        int64    `gorm:"column:id;primaryKey"`
	ProjectID int64    `gorm:"column:project_id"`
	Project   *Project `gorm:"references:project_id;foreignKey:id;"`
	Name      string   `gorm:"column:name"`
	StartDate int64    `gorm:"column:start_date"`
	EndDate   int64    `gorm:"column:end_date"`
	Status    int      `gorm:"column:status;"`
	Base
}

func (*ProjectPhase) TableName() string {
	return "project_phases"
}

func (p *ProjectPhase) GetName() string {
	if p == nil {
		return ""
	}

	return p.Name
}

func (p *ProjectPhase) GetID() int64 {
	if p == nil {
		return 0
	}

	return p.ID
}

func (p *ProjectPhase) GetProject() *Project {
	if p == nil {
		return nil
	}

	return p.Project
}

type ProjectArea struct {
	ID        int64    `gorm:"column:id;primaryKey"`
	ProjectID int64    `gorm:"column:project_id"`
	Project   *Project `gorm:"references:project_id;foreignKey:id;"`
	Name      string   `gorm:"column:name"`
	Base
}

func (*ProjectArea) TableName() string {
	return "project_areas"
}

func (p *ProjectArea) GetProject() *Project {
	if p == nil {
		return nil
	}

	return p.Project
}

func (p *ProjectArea) GetName() string {
	if p == nil {
		return ""
	}

	return p.Name
}

func (p *ProjectArea) GetID() int64 {
	if p == nil {
		return 0
	}

	return p.ID
}
