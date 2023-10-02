package timesheets

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ideal-forward/assistant-public-api/controllers/resources"
	"github.com/ideal-forward/assistant-public-api/middlewares"
	"github.com/ideal-forward/assistant-public-api/pkg/http_parser"
	"github.com/ideal-forward/assistant-public-api/pkg/log"
	"github.com/ideal-forward/assistant-public-api/services"
)

func (h Handler) ListTimesheets(c *gin.Context) {
	ctx := c.Request.Context()
	req := &resources.ListTimesheetsRequest{}

	err := http_parser.BindAndValid(c, req, "MaxSize", "DateFormat")
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	userID, _, _ := middlewares.ParseToken(c)
	data, err := h.Timesheets.List(ctx, req.ProjectID)
	if err != nil {
		log.For(c).Error("[timesheets-list] query failed", log.Field("user_id", userID), log.Field("project_id", req.ProjectID), log.Err(err))
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	resp := make([]*resources.Timesheets, 0)
	for _, val := range data {
		resp = append(resp, &resources.Timesheets{
			ID: val.ID,
			Creator: &resources.User{
				ID:   val.GetCreator().GetID(),
				Name: val.GetCreator().GetName(),
			},
			Date: services.FormatTimestamp(val.Date),
		})
	}

	c.JSON(http.StatusOK, &resources.Response{
		Data: resp,
	})
}
