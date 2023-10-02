package timesheets

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ideal-forward/assistant-public-api/controllers/resources"
	"github.com/ideal-forward/assistant-public-api/middlewares"
	"github.com/ideal-forward/assistant-public-api/pkg/log"
)

func (h Handler) DeleteExecutor(c *gin.Context) {
	ctx := c.Request.Context()
	userID, _, _ := middlewares.ParseToken(c)

	timesheetsID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Errorf("invalid parameter"))
		return
	}

	id, err := strconv.ParseInt(c.Param("executor_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Errorf("invalid parameter"))
		return
	}

	err = h.Timesheets.DeleteExecutor(ctx, id)
	if err != nil {
		log.For(c).Error("[timesheets-executor-delete] execute failed", log.Field("user_id", userID), log.Field("timesheetsID", timesheetsID), log.Err(err))
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	log.For(c).Info("[timesheets-executor-delete] process success", log.Field("user_id", userID), log.Field("timesheets_id", id))
	c.JSON(http.StatusOK, &resources.Response{
		Data: resources.Empty{},
	})
}
