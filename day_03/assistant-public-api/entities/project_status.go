package entities

type ProjectStatus int

const (
	ProjectStatusUnknown      ProjectStatus = 0
	ProjectStatusImplementing ProjectStatus = 1
	ProjectStatusCancel       ProjectStatus = 2
	ProjectStatusCompleted    ProjectStatus = 3
)

var ProjectStatusMessage = map[ProjectStatus]string{
	ProjectStatusUnknown:      "unknown",
	ProjectStatusImplementing: "implementing",
	ProjectStatusCancel:       "cancel",
	ProjectStatusCompleted:    "completed",
}

var ProjectStatusComment = map[ProjectStatus]string{
	ProjectStatusUnknown:      "không xác định",
	ProjectStatusImplementing: "đang thực hiện",
	ProjectStatusCancel:       "huỷ",
	ProjectStatusCompleted:    "hoàn thành",
}

func NewProjectStatus(status int) ProjectStatus {
	if (status < int(ProjectStatusImplementing)) || (status > int(ProjectStatusCompleted)) {
		return ProjectStatusUnknown
	}
	return ProjectStatus(status)
}

func NewProjectStatusFromString(status string) ProjectStatus {
	switch status {
	case "implementing":
		return ProjectStatusImplementing
	case "cancel":
		return ProjectStatusCancel
	case "completed":
		return ProjectStatusCompleted
	default:
		return ProjectStatusUnknown
	}
}

func (s ProjectStatus) String() string {
	return ProjectStatusMessage[s]
}

func (s ProjectStatus) Value() int {
	return int(s)
}

func (s ProjectStatus) Comment() string {
	return ProjectStatusComment[s]
}
