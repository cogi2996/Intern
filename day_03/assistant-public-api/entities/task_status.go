package entities

type TaskStatus int

const (
	StatusUnknown      TaskStatus = 0
	StatusCreated      TaskStatus = 1
	StatusImplementing TaskStatus = 2
	StatusAccepted     TaskStatus = 3
	StatusRequested    TaskStatus = 4
	StatusApproved     TaskStatus = 5
	StatusRejected     TaskStatus = 6
	StatusCancel       TaskStatus = 7
	StatusCompleted    TaskStatus = 8
)

var TaskStatusMessage = map[TaskStatus]string{
	StatusUnknown:      "unknown",
	StatusCreated:      "created",
	StatusImplementing: "implementing",
	StatusAccepted:     "accepted",
	StatusRequested:    "requested",
	StatusApproved:     "approved",
	StatusRejected:     "rejected",
	StatusCancel:       "cancel",
	StatusCompleted:    "completed",
}

var TaskStatusComment = map[TaskStatus]string{
	StatusUnknown:      "không xác định",
	StatusCreated:      "tạo",
	StatusImplementing: "thực hiện",
	StatusAccepted:     "nhận nhiệm vụ",
	StatusRequested:    "đề nghị nghiệm thu",
	StatusApproved:     "đã nghiệm thu",
	StatusRejected:     "từ chối, yêu cầu làm lại",
	StatusCancel:       "huỷ",
	StatusCompleted:    "xác nhận hoàn thành",
}

func NewTaskStatus(status int) TaskStatus {
	if (status < int(StatusCreated)) || (status > int(StatusCompleted)) {
		return StatusUnknown
	}
	return TaskStatus(status)
}

func NewTaskStatusFromString(status string) TaskStatus {
	switch status {
	case "created":
		return StatusCreated
	case "accepted":
		return StatusAccepted
	case "implementing":
		return StatusImplementing
	case "requested":
		return StatusRequested
	case "approved":
		return StatusApproved
	case "rejected":
		return StatusRejected
	case "cancel":
		return StatusCancel
	case "completed":
		return StatusCompleted
	default:
		return StatusUnknown
	}
}

func (s TaskStatus) String() string {
	return TaskStatusMessage[s]
}

func (s TaskStatus) Value() int {
	return int(s)
}

func (s TaskStatus) Comment() string {
	return TaskStatusComment[s]
}
