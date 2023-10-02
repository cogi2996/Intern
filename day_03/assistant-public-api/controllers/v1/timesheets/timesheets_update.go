package timesheets

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ideal-forward/assistant-public-api/controllers/resources"
	"github.com/ideal-forward/assistant-public-api/entities"
	"github.com/ideal-forward/assistant-public-api/middlewares"
	"github.com/ideal-forward/assistant-public-api/pkg/http_parser"
	"github.com/ideal-forward/assistant-public-api/pkg/log"
	"github.com/ideal-forward/assistant-public-api/services"
)

func (h Handler) UpdateTimesheets(c *gin.Context) {
	ctx := c.Request.Context()
	req := &resources.UpdateTimesheetsRequest{}

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

	newData := &entities.Timesheets{
		Base: entities.Base{
			UpdatedBy: userID,
		},
		ID:                            id,
		ConsMorningPersonPlanned:      req.ConsMorningPersonPlanned,
		ConsAfternoonPersonPlanned:    req.ConsAfternoonPersonPlanned,
		ConsEveningPersonPlanned:      req.ConsEveningPersonPlanned,
		ConsMorningPerson:             req.ConsMorningPerson,
		ConsAfternoonPerson:           req.ConsAfternoonPerson,
		ConsEveningPerson:             req.ConsEveningPerson,
		ConsOvertimeHour:              req.ConsOvertimeHour,
		ConsCoefficient:               req.ConsCoefficient,
		PartnerMorningPersonPlanned:   req.PartnerMorningPersonPlanned,
		PartnerAfternoonPersonPlanned: req.PartnerAfternoonPersonPlanned,
		PartnerEveningPersonPlanned:   req.PartnerEveningPersonPlanned,
		PartnerMorningPerson:          req.PartnerMorningPerson,
		PartnerAfternoonPerson:        req.PartnerAfternoonPerson,
		PartnerEveningPerson:          req.PartnerEveningPerson,
		PartnerOvertimeHour:           req.PartnerOvertimeHour,
		PartnerCoefficient:            req.PartnerCoefficient,
	}
	newData.Date, _ = services.StringToTimestamp(req.Date)

	err = h.Timesheets.Update(ctx, newData)
	if err != nil {
		log.For(c).Error("[timesheets-update] update new data failed", log.Field("user_id", userID), log.Field("task_id", id), log.Err(err))
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	h.makeUpdateComment(ctx, userID, id)

	log.For(c).Info("[timesheets-update] insert success", log.Field("user_id", userID), log.Field("task_id", id), log.Err(err))
	c.JSON(http.StatusOK, &resources.Response{
		Data: &resources.Empty{},
	})
}

func (h Handler) makeUpdateComment(ctx context.Context, userID, timesheetsID int64) error {
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
