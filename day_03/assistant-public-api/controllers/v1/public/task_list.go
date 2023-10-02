package public

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ideal-forward/assistant-public-api/controllers/resources"
	"github.com/ideal-forward/assistant-public-api/entities"
	"github.com/ideal-forward/assistant-public-api/middlewares"
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

	userID, _, _ := middlewares.ParseToken(c)
	var tasks []*entities.Task
	executor, err := h.Executor.ReadByRepresenterID(ctx, userID)
	if err == nil {
		tasks, err = h.Task.List(ctx, &services.TaskFilter{
			ExecutorID:    executor.ID,
			Status:        entities.NewTaskStatusFromString(req.Status),
			PriorityLevel: entities.NewPriorityLevelFromString(req.PriorityLevel),
		}, true)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
	} else {
		tasks, err = h.Task.List(ctx, &services.TaskFilter{
			ProjectID:     req.ProjectID,
			CreatorID:     req.CreatorID,
			AcceptorID:    req.AcceptorID,
			ExecutorID:    req.ExecutorID,
			StartTime:     req.StartTime,
			EndTime:       req.EndTime,
			Status:        entities.NewTaskStatusFromString(req.Status),
			PriorityLevel: entities.NewPriorityLevelFromString(req.PriorityLevel),
		}, true)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
	}

	resp := resources.ListTaskResponse{
		Tasks: make([]*resources.Task, 0),
	}
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
			Area: &resources.ProjectArea{
				ID:   val.GetArea().GetID(),
				Name: val.GetArea().GetName(),
			},
			Phase: &resources.ProjectPhase{
				ID:   val.GetPhase().GetID(),
				Name: val.GetPhase().GetName(),
			},
			Creator: &resources.User{
				ID:   val.GetCreator().GetID(),
				Name: val.GetCreator().GetName(),
			},
			Executor: &resources.Executor{
				ID:   val.GetExecutor().GetID(),
				Name: val.GetExecutor().GetName(),
			},
			Reporter: &resources.User{
				ID:   val.GetReporter().GetID(),
				Name: val.GetReporter().GetName(),
			},
			Acceptor: &resources.User{
				ID:   val.GetAcceptor().GetID(),
				Name: val.GetAcceptor().GetName(),
			},
			StartTime:     services.FormatTimestamp(val.StartTime),
			EndTime:       services.FormatTimestamp(val.EndTime),
			Unit:          val.Unit,
			Quantity:      val.Quantity,
			Price:         val.Price,
			Description:   val.Description,
			Status:        entities.NewTaskStatus(val.Status).String(),
			PriorityLevel: entities.NewPriorityLevel(val.PriorityLevel).String(),
			Star:          val.Star,
		})
	}

	// hardcode to test
	req.NeedExportFile = true
	if req.NeedExportFile {
		filePath := fmt.Sprintf("%stasks-%s.xlsx", services.Config.FileStorage.Folder, uuid.NewString())
		if err = h.Excel.Create(tasks, filePath); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		resp.ExcelURL = h.FileNamer.ToPublicFile(h.FileNamer.ToRelativeFile(filePath))
	}

	c.JSON(http.StatusOK, &resources.Response{
		Data: resp,
	})
}
