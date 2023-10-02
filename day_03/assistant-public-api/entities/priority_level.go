package entities

type PriorityLevel int

const (
	PriorityLevelUnknown PriorityLevel = 0
	PriorityLevelLow     PriorityLevel = 1
	PriorityLevelNormal  PriorityLevel = 2
	PriorityLevelHigh    PriorityLevel = 3
)

var PriorityLevelMessage = map[PriorityLevel]string{
	PriorityLevelUnknown: "unknown",
	PriorityLevelLow:     "low",
	PriorityLevelNormal:  "normal",
	PriorityLevelHigh:    "high",
}

var PriorityLevelComment = map[PriorityLevel]string{
	PriorityLevelUnknown: "không xác định",
	PriorityLevelLow:     "không ưu tiên",
	PriorityLevelNormal:  "bình thường",
	PriorityLevelHigh:    "gấp",
}

func NewPriorityLevel(level int) PriorityLevel {
	if (level < int(PriorityLevelUnknown)) || (level > int(PriorityLevelHigh)) {
		return PriorityLevelUnknown
	}
	return PriorityLevel(level)
}

func NewPriorityLevelFromString(status string) PriorityLevel {
	switch status {
	case "low":
		return PriorityLevelLow
	case "normal":
		return PriorityLevelNormal
	case "high":
		return PriorityLevelHigh
	default:
		return PriorityLevelUnknown
	}
}

func (s PriorityLevel) String() string {
	return PriorityLevelMessage[s]
}

func (s PriorityLevel) Comment() string {
	return PriorityLevelComment[s]
}

func (s PriorityLevel) Value() int {
	return int(s)
}
