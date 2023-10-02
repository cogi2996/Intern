package resources

type Timesheets struct {
	ID                            int64    `json:"id"`
	Creator                       *User    `json:"creator"`
	Project                       *Project `json:"project"`
	Date                          string   `json:"date"`
	ConsMorningPersonPlanned      int      `json:"cons_morning_person_planned"`
	ConsAfternoonPersonPlanned    int      `json:"cons_afternoon_person_planned"`
	ConsEveningPersonPlanned      int      `json:"cons_evening_person_planned"`
	ConsMorningPerson             int      `json:"cons_morning_person"`
	ConsAfternoonPerson           int      `json:"cons_afternoon_person"`
	ConsEveningPerson             int      `json:"cons_evening_person"`
	ConsOvertimeHour              int      `json:"cons_overtime_hour"`
	ConsCoefficient               float32  `json:"cons_coefficient"`
	PartnerMorningPersonPlanned   int      `json:"partner_morning_person_planned"`
	PartnerAfternoonPersonPlanned int      `json:"partner_afternoon_person_planned"`
	PartnerEveningPersonPlanned   int      `json:"partner_evening_person_planned"`
	PartnerMorningPerson          int      `json:"partner_morning_person"`
	PartnerAfternoonPerson        int      `json:"partner_afternoon_person"`
	PartnerEveningPerson          int      `json:"partner_evening_person"`
	PartnerOvertimeHour           int      `json:"partner_overtime_hour"`
	PartnerCoefficient            float32  `json:"partner_coefficient"`
}

type TimesheetsComment struct {
	ID         int64       `json:"id"`
	Timesheets *Timesheets `json:"timesheets,omitempty"`
	Creator    *User       `json:"creator,omitempty"`
	Msg        string      `json:"msg,omitempty"`
	Time       string      `json:"time,omitempty"`
}

type TimesheetsExecutor struct {
	ID                     int64       `json:"id"`
	UUID                   string      `json:"uuid"`
	Creator                *User       `json:"creator"`
	Timesheets             *Timesheets `json:"timesheets"`
	Executor               *Executor   `json:"executor"`
	MorningPersonPlanned   int         `json:"morning_person_planned"`
	AfternoonPersonPlanned int         `json:"afternoon_person_planned"`
	EveningPersonPlanned   int         `json:"evening_person_planned"`
	MorningPerson          int         `json:"morning_person"`
	AfternoonPerson        int         `json:"afternoon_person"`
	EveningPerson          int         `json:"evening_person"`
	OvertimeHour           int         `json:"overtime_hour"`
	Coefficient            float32     `json:"coefficient"`
}

type CreateTimesheetsRequest struct {
	ProjectID int64  `json:"project_id"`
	Date      string `json:"date"`
}

type UpdateTimesheetsRequest struct {
	Date                          string  `json:"date"`
	ConsMorningPersonPlanned      int     `json:"cons_morning_person_planned"`
	ConsAfternoonPersonPlanned    int     `json:"cons_afternoon_person_planned"`
	ConsEveningPersonPlanned      int     `json:"cons_evening_person_planned"`
	ConsMorningPerson             int     `json:"cons_morning_person"`
	ConsAfternoonPerson           int     `json:"cons_afternoon_person"`
	ConsEveningPerson             int     `json:"cons_evening_person"`
	ConsOvertimeHour              int     `json:"cons_overtime_hour"`
	ConsCoefficient               float32 `json:"cons_coefficient"`
	PartnerMorningPersonPlanned   int     `json:"partner_morning_person_planned"`
	PartnerAfternoonPersonPlanned int     `json:"partner_afternoon_person_planned"`
	PartnerEveningPersonPlanned   int     `json:"partner_evening_person_planned"`
	PartnerMorningPerson          int     `json:"partner_morning_person"`
	PartnerAfternoonPerson        int     `json:"partner_afternoon_person"`
	PartnerEveningPerson          int     `json:"partner_evening_person"`
	PartnerOvertimeHour           int     `json:"partner_overtime_hour"`
	PartnerCoefficient            float32 `json:"partner_coefficient"`
}

type ListTimesheetsRequest struct {
	ProjectID int64 `json:"project_id" form:"project_id"`
}

type CreateTimesheetsExecutorRequest struct {
	ExecutorID             int64   `json:"executor_id"`
	MorningPersonPlanned   int     `json:"morning_person_planned"`
	AfternoonPersonPlanned int     `json:"afternoon_person_planned"`
	EveningPersonPlanned   int     `json:"evening_person_planned"`
	MorningPerson          int     `json:"morning_person"`
	AfternoonPerson        int     `json:"afternoon_person"`
	EveningPerson          int     `json:"evening_person"`
	OvertimeHour           int     `json:"overtime_hour"`
	Coefficient            float32 `json:"coefficient"`
}

type UpdateTimesheetsExecutorRequest struct {
	ExecutorID             int64   `json:"executor_id"`
	MorningPersonPlanned   int     `json:"morning_person_planned"`
	AfternoonPersonPlanned int     `json:"afternoon_person_planned"`
	EveningPersonPlanned   int     `json:"evening_person_planned"`
	MorningPerson          int     `json:"morning_person"`
	AfternoonPerson        int     `json:"afternoon_person"`
	EveningPerson          int     `json:"evening_person"`
	OvertimeHour           int     `json:"overtime_hour"`
	Coefficient            float32 `json:"coefficient"`
	Reason                 string  `json:"reason"`
}
