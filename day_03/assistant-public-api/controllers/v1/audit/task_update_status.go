package audit

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ideal-forward/assistant-public-api/controllers/resources"
	"github.com/ideal-forward/assistant-public-api/entities"
	"github.com/ideal-forward/assistant-public-api/middlewares"
	"github.com/ideal-forward/assistant-public-api/pkg/http_parser"
)

func (h Handler) UpdateTaskStatus(c *gin.Context) {
	ctx := c.Request.Context()
	req := &resources.UpdateTaskStatusRequest{}

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

	status := entities.NewTaskStatusFromString(req.Status)
	if status == entities.StatusUnknown {
		c.JSON(http.StatusBadRequest, fmt.Errorf("invalid status parameter"))
		return
	}

	err = h.Task.UpdateStatus(ctx, userID, id, status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	user, err := h.User.Read(ctx, userID)
	if err != nil {
		c.JSON(http.StatusOK, &resources.Response{
			Data: &resources.Empty{},
		})

		return
	}

	h.TaskComment.Create(ctx, &entities.AuditTaskComment{
		Base: entities.Base{
			CreatedBy: userID,
		},
		TaskID:  id,
		Comment: fmt.Sprintf("%s %s", user.Name, status.Comment()),
	})

	c.JSON(http.StatusOK, &resources.Response{
		Data: &resources.Empty{},
	})
}
