package public

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ideal-forward/assistant-public-api/controllers/resources"
)

func (h Handler) DeleteTaskComment(c *gin.Context) {
	ctx := c.Request.Context()

	id, err := strconv.ParseInt(c.Param("comment_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Errorf("invalid parameter"))
		return
	}

	err = h.TaskComment.Delete(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, &resources.Response{
		Data: resources.Empty{},
	})
}
