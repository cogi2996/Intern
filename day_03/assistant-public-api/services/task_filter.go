package services

import (
	"time"

	"github.com/ideal-forward/assistant-public-api/entities"
)

type TaskFilter struct {
	ProjectID     int64
	CreatorID     int64
	AcceptorID    int64
	ExecutorID    int64
	StartTime     string
	EndTime       string
	Status        entities.TaskStatus
	PriorityLevel entities.PriorityLevel
}

func (t *TaskFilter) ToMap() map[string]any {
	result := make(map[string]any)
	if t.ProjectID > 0 {
		result["project_id"] = t.ProjectID
	}

	if t.CreatorID > 0 {
		result["created_by"] = t.CreatorID
	}

	if t.AcceptorID > 0 {
		result["accepted_by"] = t.AcceptorID
	}

	if t.ExecutorID > 0 {
		result["executed_by"] = t.ExecutorID
	}

	if t.StartTime != "" {
		if startTime, err := StringToTime(t.StartTime); err == nil {
			result["start_time"] = startTime.UnixMilli()
		}
	}

	if t.EndTime != "" {
		if endTime, err := StringToTime(t.EndTime); err == nil {
			result["end_time"] = endTime.Add((24*3600 - 1) * time.Second).UnixMilli()
		}
	}

	if t.Status != entities.StatusUnknown {
		result["status"] = t.Status.Value()
	}

	if t.PriorityLevel != entities.PriorityLevelUnknown {
		result["priority_level"] = t.PriorityLevel.Value()
	}

	return result
}

type TaskFilters struct {
	ProjectIDs     []int64
	CreatorIDs     []int64
	AcceptorIDs    []int64
	ExecutorIDs    []int64
	StartTime      int64
	EndTime        int64
	Statuses       []entities.TaskStatus
	PriorityLevels []entities.PriorityLevel
}

func (t *TaskFilters) ToMap() map[string]any {
	result := make(map[string]any)
	if len(t.ProjectIDs) > 0 {
		result["project_id"] = t.ProjectIDs
	}

	if len(t.CreatorIDs) > 0 {
		result["created_by"] = t.CreatorIDs
	}

	if len(t.AcceptorIDs) > 0 {
		result["accepted_by"] = t.AcceptorIDs
	}

	if len(t.ExecutorIDs) > 0 {
		result["executed_by"] = t.ExecutorIDs
	}

	if t.StartTime != 0 {
		result["start_time"] = t.StartTime
	}

	if t.EndTime != 0 {
		result["end_time"] = t.EndTime
	}

	if len(t.Statuses) > 0 {
		statuses := make([]int, 0)
		for _, val := range t.Statuses {
			statuses = append(statuses, val.Value())
		}
		result["status"] = statuses
	}

	if len(t.PriorityLevels) > 0 {
		priorityLevels := make([]int, 0)
		for _, val := range t.PriorityLevels {
			priorityLevels = append(priorityLevels, val.Value())
		}
		result["priority_level"] = priorityLevels
	}

	return result
}
