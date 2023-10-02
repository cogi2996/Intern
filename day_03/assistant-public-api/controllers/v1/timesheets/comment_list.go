package timesheets

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ideal-forward/assistant-public-api/controllers/resources"
	"github.com/ideal-forward/assistant-public-api/services"
)

func (h Handler) ListComment(c *gin.Context) {
	ctx := c.Request.Context()

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Errorf("invalid parameter"))
		return
	}

	data, err := h.Timesheets.ListComment(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	comments := make([]*resources.TimesheetsComment, 0)
	for _, val := range data {
		comments = append(comments, &resources.TimesheetsComment{
			ID: val.ID,
			Creator: &resources.User{
				ID:   val.GetCreator().GetID(),
				Name: val.GetCreator().GetName(),
			},
			Msg:  val.Comment,
			Time: services.FormatTime(val.CreatedAt),
		})
	}

	c.JSON(http.StatusOK, &resources.Response{
		Data: comments,
	})
}
