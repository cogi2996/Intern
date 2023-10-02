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
	"github.com/ideal-forward/assistant-public-api/pkg/log"
)

func (h Handler) Remind(c *gin.Context) {
	ctx := c.Request.Context()
	req := &resources.RemindTaskRequest{}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Errorf("invalid parameter"))
		return
	}

	err = http_parser.BindAndValid(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	userID, _, _ := middlewares.ParseToken(c)

	task, err := h.Task.Read(ctx, id)
	if err != nil {
		log.For(c).Error("[task-remind] query failed", log.Field("user_id", userID), log.Field("task_id", id), log.Err(err))
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	// send notification
	executor, err := h.Executor.Read(ctx, task.ExecutedBy)
	if err != nil {
		log.For(c).Error("[task-remind] query task failed", log.Field("user_id", userID), log.Field("task_id", id), log.Err(err))
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	_, err = h.Notification.Create(ctx, &entities.Notification{
		Base: entities.Base{
			CreatedBy: userID,
		},
		Title:        "Nhắc nhở từ giám sát",
		Message:      req.Message,
		PublicTaskID: id,
		ReceivedBy:   executor.RepresentedBy,
	})
	if err != nil {
		log.For(c).Error("[task-remind] insert failed", log.Field("user_id", userID), log.Field("task_id", id), log.Err(err))
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, &resources.Response{
		Data: &resources.CreateResponse{
			ID: id,
		},
	})
}
