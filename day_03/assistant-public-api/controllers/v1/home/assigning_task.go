package home

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ideal-forward/assistant-public-api/controllers/resources"
	"github.com/ideal-forward/assistant-public-api/entities"
	"github.com/ideal-forward/assistant-public-api/middlewares"
	"github.com/ideal-forward/assistant-public-api/services"
)

func (h Handler) ListAssigningTask(c *gin.Context) {
	ctx := c.Request.Context()
	userID, _, _ := middlewares.ParseToken(c)

	publicTasks, err := h.PublicTask.List(ctx, &services.TaskFilter{
		CreatorID: userID,
	}, true)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	privateTasks, err := h.PrivateTask.List(ctx, &services.TaskFilter{
		CreatorID: userID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	resp := resources.ListAssigningTaskResponse{
		Tasks: make([]*resources.AssigningTask, 0),
	}
	for _, val := range publicTasks {
		resp.Tasks = append(resp.Tasks, &resources.AssigningTask{
			PublicID: val.ID,
			Star:     val.Star,
			Task: resources.Task{
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
			},
		})
	}

	for _, val := range privateTasks {
		if val.GetExecutor().GetID() == userID {
			continue
		}

		resp.Tasks = append(resp.Tasks, &resources.AssigningTask{
			PrivateID: val.ID,
			Task: resources.Task{
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
			},
		})
	}

	c.JSON(http.StatusOK, &resources.Response{
		Data: resp,
	})
}
