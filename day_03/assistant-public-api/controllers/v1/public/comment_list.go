package public

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ideal-forward/assistant-public-api/controllers/resources"
	"github.com/ideal-forward/assistant-public-api/services"
)

func (h Handler) ListTaskComment(c *gin.Context) {
	ctx := c.Request.Context()

	taskID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Errorf("invalid parameter"))
		return
	}

	comments, err := h.TaskComment.List(ctx, taskID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	resp := &resources.ListTaskCommentResponse{}
	for _, val := range comments {
		resp.Comments = append(resp.Comments, &resources.Comment{
			ID: val.ID,
			Creator: &resources.User{
				ID:   val.GetCreator().GetID(),
				Name: val.GetCreator().GetName(),
			},
			Task: &resources.Task{
				ID:   val.GetTask().GetID(),
				Name: val.GetTask().GetName(),
			},
			Msg:  val.Comment,
			Time: services.FormatTime(val.CreatedAt),
		})
	}

	c.JSON(http.StatusOK, &resources.Response{
		Data: resp,
	})
}
