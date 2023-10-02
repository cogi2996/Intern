package public

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

	// hardcode
	status := entities.NewTaskStatusFromString(req.Status)
	if status == entities.StatusUnknown {
		c.JSON(http.StatusBadRequest, fmt.Errorf("invalid status parameter"))
		return
	}

	if status == entities.StatusAccepted {
		status = entities.StatusApproved
	}
	if (status != entities.StatusApproved) && (status != entities.StatusCompleted) {
		req.Star = 0
	}

	err = h.Task.UpdateStatus(ctx, id, status, req.Star)
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

	task, err := h.Task.Read(ctx, id)
	if err != nil {
		c.JSON(http.StatusOK, &resources.Response{
			Data: &resources.Empty{},
		})

		return
	}

	h.TaskComment.Create(ctx, &entities.TaskComment{
		Base: entities.Base{
			CreatedBy: userID,
		},
		TaskID:  id,
		Comment: fmt.Sprintf("%s %s", user.Name, status.Comment()),
	})

	if status == entities.StatusApproved {
		h.TaskComment.Create(ctx, &entities.TaskComment{
			Base: entities.Base{
				CreatedBy: userID,
			},
			TaskID:  id,
			Comment: fmt.Sprintf("%s %s", user.Name, entities.StatusCompleted.Comment()),
		})

		if req.Star < 2 {
			executor, err := h.Executor.Read(ctx, task.ExecutedBy)
			if err == nil {
				h.Notification.Create(ctx, &entities.Notification{
					Base: entities.Base{
						CreatedBy: userID,
					},
					Title:        "Đánh giá một sao!",
					Message:      "Nhiệm vụ đã hoàn thành nhưng bị đánh giá chất lượng không tốt!",
					PublicTaskID: id,
					ReceivedBy:   executor.RepresentedBy,
				})
			}
		}
	}

	c.JSON(http.StatusOK, &resources.Response{
		Data: &resources.Empty{},
	})
}
