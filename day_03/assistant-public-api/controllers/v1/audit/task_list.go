package audit

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ideal-forward/assistant-public-api/controllers/resources"
	"github.com/ideal-forward/assistant-public-api/entities"
	"github.com/ideal-forward/assistant-public-api/pkg/http_parser"
	"github.com/ideal-forward/assistant-public-api/services"
)

func (h Handler) ListTask(c *gin.Context) {
	ctx := c.Request.Context()
	req := &resources.ListTaskRequest{}

	err := http_parser.BindAndValid(c, req, "MaxSize", "DateFormat")
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	// userID, _, _ := middlewares.ParseToken(c)

	/*
		var tasks []*entities.Task
		if req.FilterBy == "" {
			tasks, err = h.Task.ListAll(ctx)
		} else {
			tasks, err = h.Task.ListOwnTask(ctx, userID)
		}
	*/
	tasks, err := h.Task.List(ctx, &services.TaskFilter{
		ProjectID:     req.ProjectID,
		CreatorID:     req.CreatorID,
		AcceptorID:    req.AcceptorID,
		ExecutorID:    req.ExecutorID,
		StartTime:     req.StartTime,
		EndTime:       req.EndTime,
		Status:        entities.NewTaskStatusFromString(req.Status),
		PriorityLevel: entities.NewPriorityLevelFromString(req.PriorityLevel),
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	resp := resources.ListTaskResponse{}
	for _, val := range tasks {
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
