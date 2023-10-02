package private

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ideal-forward/assistant-public-api/controllers/resources"
	"github.com/ideal-forward/assistant-public-api/entities"
	"github.com/ideal-forward/assistant-public-api/middlewares"
	"github.com/ideal-forward/assistant-public-api/pkg/http_parser"
	"github.com/ideal-forward/assistant-public-api/services"
)

func (h Handler) CreateTask(c *gin.Context) {
	ctx := c.Request.Context()
	req := &resources.CreateTaskRequest{}

	err := http_parser.BindAndValid(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	userID, _, _ := middlewares.ParseToken(c)

	data := &entities.PrivateTask{
		Base: entities.Base{
			CreatedBy: userID,
		},
		Name:          req.Name,
		ParentTaskID:  req.ParentTaskID,
		ProjectID:     req.ProjectID,
		AreaID:        req.AreaID,
		PhaseID:       req.PhaseID,
		ExecutedBy:    req.ExecutorID,
		AcceptedBy:    req.AcceptorID,
		Quantity:      req.Quantity,
		Price:         req.Price,
		Unit:          req.Unit,
		Description:   req.Description,
		PriorityLevel: entities.NewPriorityLevelFromString(req.PriorityLevel).Value(),
	}

	data.StartTime, _ = services.StringToTimestamp(req.StartTime)
	data.EndTime, _ = services.StringToTimestamp(req.EndTime)
	if data.AcceptedBy == 0 {
		data.AcceptedBy = userID
	}

	id, err := h.Task.Create(ctx, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, &resources.Response{
		Data: &resources.CreateResponse{
			ID:   id,
			UUID: data.UUID,
		},
	})
}
