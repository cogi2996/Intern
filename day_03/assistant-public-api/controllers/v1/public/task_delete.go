package public

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ideal-forward/assistant-public-api/controllers/resources"
	"github.com/ideal-forward/assistant-public-api/middlewares"
	"github.com/ideal-forward/assistant-public-api/pkg/http_parser"
	"github.com/ideal-forward/assistant-public-api/pkg/log"
)

func (h Handler) DeleteTask(c *gin.Context) {
	ctx := c.Request.Context()
	userID, _, _ := middlewares.ParseToken(c)

	req := &resources.DeleteTaskRequest{}

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

	err = h.Task.Delete(ctx, id)
	if err != nil {
		log.For(c).Error("[task-delete] insert failed", log.Field("user_id", userID), log.Field("task_id", id), log.Err(err))
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	log.For(c).Info("[task-delete] process success", log.Field("user_id", userID), log.Field("task_id", id))
	c.JSON(http.StatusOK, &resources.Response{
		Data: resources.Empty{},
	})
}
