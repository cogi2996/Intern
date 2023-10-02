package entities

type Timesheets struct {
	ID                            int64                 `gorm:"column:id;primaryKey;"`
	ProjectID                     int64                 `gorm:"column:project_id;"`
	Project                       *Project              `gorm:"references:project_id;foreignKey:id;"`
	Date                          int64                 `gorm:"column:date;"`
	ConsMorningPersonPlanned      int                   `gorm:"column:cons_morning_person_planned;"`
	ConsAfternoonPersonPlanned    int                   `gorm:"column:cons_afternoon_person_planned;"`
	ConsEveningPersonPlanned      int                   `gorm:"column:cons_evening_person_planned;"`
	ConsMorningPerson             int                   `gorm:"column:cons_morning_person;"`
	ConsAfternoonPerson           int                   `gorm:"column:cons_afternoon_person;"`
	ConsEveningPerson             int                   `gorm:"column:cons_evening_person;"`
	ConsOvertimeHour              int                   `gorm:"column:cons_overtime_hour;"`
	ConsCoefficient               float32               `gorm:"column:cons_coefficient;"`
	PartnerMorningPersonPlanned   int                   `gorm:"column:partner_morning_person_planned;"`
	PartnerAfternoonPersonPlanned int                   `gorm:"column:partner_afternoon_person_planned;"`
	PartnerEveningPersonPlanned   int                   `gorm:"column:partner_evening_person_planned;"`
	PartnerMorningPerson          int                   `gorm:"column:partner_morning_person;"`
	PartnerAfternoonPerson        int                   `gorm:"column:partner_afternoon_person;"`
	PartnerEveningPerson          int                   `gorm:"column:partner_evening_person;"`
	PartnerOvertimeHour           int                   `gorm:"column:partner_overtime_hour;"`
	PartnerCoefficient            float32               `gorm:"column:partner_coefficient;"`
	Executors                     []*TimesheetsExecutor `gorm:"references:id;foreignKey:timesheets_id;"`
	Comments                      []*TimesheetsComment  `gorm:"references:id;foreignKey:timesheets_id;"`
	Base
}

func (t *Timesheets) GetProject() *Project {
	if t == nil {
		return nil
	}

	return t.Project
}

func (t *Timesheets) TableName() string {
	return "timesheets"
}

type TimesheetsExecutor struct {
	ID                     int64       `gorm:"column:id;primaryKey;"`
	TimesheetsID           int64       `gorm:"column:timesheets_id;"`
	Timesheets             *Timesheets `gorm:"references:timesheets_id;foreignKey:id;"`
	ExecutorID             int64       `gorm:"column:executor_id;"`
	Executor               *Executor   `gorm:"references:executor_id;foreignKey:id;"`
	MorningPersonPlanned   int         `gorm:"column:morning_person_planned;"`
	AfternoonPersonPlanned int         `gorm:"column:afternoon_person_planned;"`
	EveningPersonPlanned   int         `gorm:"column:evening_person_planned;"`
	MorningPerson          int         `gorm:"column:morning_person;"`
	AfternoonPerson        int         `gorm:"column:afternoon_person;"`
	EveningPerson          int         `gorm:"column:evening_person;"`
	OvertimeHour           int         `gorm:"column:overtime_hour;"`
	Coefficient            float32     `gorm:"column:coefficient;"`
	Base
}

func (t *TimesheetsExecutor) GetExecutor() *Executor {
	if t == nil {
		return nil
	}

	return t.Executor
}

func (t *TimesheetsExecutor) TableName() string {
	return "timesheets_executors"
}

type TimesheetsComment struct {
	ID           int64       `gorm:"column:id;primaryKey;"`
	TimesheetsID int64       `gorm:"column:timesheets_id;"`
	Timesheets   *Timesheets `gorm:"references:timesheets_id;foreignKey:id;"`
	Comment      string      `gorm:"column:comment;"`
	Base
}

func (*TimesheetsComment) TableName() string {
	return "timesheets_comments"
}
