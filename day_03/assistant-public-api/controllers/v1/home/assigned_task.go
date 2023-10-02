package home

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ideal-forward/assistant-public-api/controllers/resources"
	"github.com/ideal-forward/assistant-public-api/entities"
	"github.com/ideal-forward/assistant-public-api/middlewares"
	"github.com/ideal-forward/assistant-public-api/services"
)

func (h Handler) ListAssignedTask(c *gin.Context) {
	ctx := c.Request.Context()
	userID, _, _ := middlewares.ParseToken(c)

	privateTasks, err := h.PrivateTask.List(ctx, &services.TaskFilter{
		ExecutorID: userID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	resp := resources.ListTaskResponse{
		Tasks: make([]*resources.Task, 0),
	}
	for _, val := range privateTasks {
		if val.GetCreator().GetID() == userID {
			continue
		}

		resp.Tasks = append(resp.Tasks, &resources.Task{
			ID:          val.ID,
			Code:        val.Code,
			Name:        val.Name,
			ProjectID:   val.ProjectID,
			ProjectName: val.GetProject().GetName(),
			Project: &resources.Project{
				ID:        val.GetProject().GetID(),
				Name:      val.GetProject().GetName(),
				Code:      val.GetProject().GetCode(),
				Address:   val.GetProject().GetAddress(),
				StartDate: services.FormatTimestamp(val.GetProject().GetStartDate()),
				EndDate:   services.FormatTimestamp(val.GetProject().GetEndDate()),
			},
			Creator: &resources.User{
				ID:   val.GetCreator().GetID(),
				Name: val.GetCreator().GetName(),
			},
			Executor: &resources.Executor{
				ID:   val.GetExecutor().GetID(),
				Name: val.GetExecutor().GetName(),
			},
			Acceptor: &resources.User{
				ID:   val.GetAcceptor().GetID(),
				Name: val.GetAcceptor().GetName(),
			},
			StartTime:     services.FormatTimestamp(val.StartTime),
			EndTime:       services.FormatTimestamp(val.EndTime),
			Description:   val.Description,
			Status:        entities.NewTaskStatus(val.Status).String(),
			PriorityLevel: entities.NewPriorityLevel(val.PriorityLevel).String(),
		})
	}

	c.JSON(http.StatusOK, &resources.Response{
		Data: resp,
	})
}
