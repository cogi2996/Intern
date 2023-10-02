package timesheets

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

func (h Handler) DeleteTimesheets(c *gin.Context) {
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

	err = h.Timesheets.Delete(ctx, id)
	if err != nil {
		log.For(c).Error("[timesheets-delete] insert failed", log.Field("user_id", userID), log.Field("timesheets_id", id), log.Err(err))
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	log.For(c).Info("[timesheets-delete] process success", log.Field("user_id", userID), log.Field("timesheets_id", id))
	c.JSON(http.StatusOK, &resources.Response{
		Data: resources.Empty{},
	})
}
