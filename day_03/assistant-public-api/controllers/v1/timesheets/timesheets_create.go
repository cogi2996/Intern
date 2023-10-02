package timesheets

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ideal-forward/assistant-public-api/controllers/resources"
	"github.com/ideal-forward/assistant-public-api/entities"
	"github.com/ideal-forward/assistant-public-api/middlewares"
	"github.com/ideal-forward/assistant-public-api/pkg/http_parser"
	"github.com/ideal-forward/assistant-public-api/pkg/log"
	"github.com/ideal-forward/assistant-public-api/services"
)

func (h Handler) CreateTimesheets(c *gin.Context) {
	ctx := c.Request.Context()
	req := &resources.CreateTimesheetsRequest{}

	err := http_parser.BindAndValid(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	userID, _, _ := middlewares.ParseToken(c)

	data := &entities.Timesheets{
		Base: entities.Base{
			CreatedBy: userID,
		},
		ProjectID: req.ProjectID,
	}
	data.Date, _ = services.StringToTimestamp(req.Date)

	id, err := h.Timesheets.Create(ctx, data)
	if err != nil {
		log.For(c).Error("[timesheets-create] insert failed", log.Field("user_id", userID), log.Err(err))
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	h.makeCreateComment(ctx, userID, id)

	log.For(c).Info("[timesheets-create] process success", log.Field("user_id", userID), log.Field("timesheets_id", id))
	c.JSON(http.StatusOK, &resources.Response{
		Data: &resources.CreateResponse{
			ID:   id,
			UUID: data.UUID,
		},
	})
}

func (h Handler) makeCreateComment(ctx context.Context, userID, timesheetsID int64) error {
	user, err := h.User.Read(ctx, userID)
	if err != nil {
		return err
	}

	_, err = h.Timesheets.CreateComment(ctx, &entities.TimesheetsComment{
		Base: entities.Base{
			CreatedBy: userID,
		},
		TimesheetsID: timesheetsID,
		Comment:      user.GetName() + " khởi tạo bảng.",
	})

	return err
}
