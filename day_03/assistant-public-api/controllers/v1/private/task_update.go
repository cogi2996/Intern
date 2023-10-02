package private

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ideal-forward/assistant-public-api/controllers/resources"
	"github.com/ideal-forward/assistant-public-api/entities"
	"github.com/ideal-forward/assistant-public-api/middlewares"
	"github.com/ideal-forward/assistant-public-api/pkg/http_parser"
	"github.com/ideal-forward/assistant-public-api/services"
)

func (h Handler) UpdateTask(c *gin.Context) {
	ctx := c.Request.Context()
	req := &resources.UpdateTaskRequest{}

	err := http_parser.BindAndValid(c, req, "MaxSize", "DateFormat")
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	userID, _, _ := middlewares.ParseToken(c)

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Errorf("invalid parameter"))
		return
	}

	data := &entities.PrivateTask{
		ID:            id,
		Name:          req.Name,
		ProjectID:     req.ProjectID,
		AreaID:        req.AreaID,
		PhaseID:       req.PhaseID,
		ExecutedBy:    req.ExecutorID,
		AcceptedBy:    req.AcceptorID,
		Quantity:      req.Quantity,
		Price:         req.Price,
		Unit:          req.Unit,
		Description:   req.Description,
		Status:        req.Status,
		PriorityLevel: entities.NewPriorityLevelFromString(req.PriorityLevel).Value(),
	}
	data.StartTime, _ = services.StringToTimestamp(req.StartTime)
	data.EndTime, _ = services.StringToTimestamp(req.EndTime)
	if data.AcceptedBy == 0 {
		data.AcceptedBy = userID
	}

	err = h.Task.Update(ctx, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, &resources.Response{
		Data: &resources.Empty{},
	})
}
