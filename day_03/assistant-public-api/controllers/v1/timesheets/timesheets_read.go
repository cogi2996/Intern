package timesheets

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ideal-forward/assistant-public-api/controllers/resources"
	"github.com/ideal-forward/assistant-public-api/middlewares"
	"github.com/ideal-forward/assistant-public-api/pkg/log"
	"github.com/ideal-forward/assistant-public-api/services"
)

func (h Handler) ReadTimesheets(c *gin.Context) {
	ctx := c.Request.Context()
	userID, _, _ := middlewares.ParseToken(c)

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Errorf("invalid parameter"))
		return
	}

	data, err := h.Timesheets.Read(ctx, id)
	if err != nil {
		log.For(c).Error("[timesheets-detail] query failed", log.Field("user_id", userID), log.Field("timesheets_id", id), log.Err(err))
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	resp := &resources.Timesheets{
		ID: data.ID,
		Creator: &resources.User{
			ID:   data.GetCreator().GetID(),
			Name: data.Project.GetName(),
		},
		Project: &resources.Project{
			ID:   data.GetProject().GetID(),
			Name: data.GetProject().GetName(),
			Code: data.GetProject().GetCode(),
		},
		Date:                          services.FormatTimestamp(data.Date),
		ConsMorningPersonPlanned:      data.ConsMorningPersonPlanned,
		ConsAfternoonPersonPlanned:    data.ConsAfternoonPersonPlanned,
		ConsEveningPersonPlanned:      data.ConsEveningPersonPlanned,
		ConsMorningPerson:             data.ConsMorningPerson,
		ConsAfternoonPerson:           data.ConsAfternoonPerson,
		ConsEveningPerson:             data.ConsEveningPerson,
		ConsOvertimeHour:              data.ConsOvertimeHour,
		ConsCoefficient:               data.ConsCoefficient,
		PartnerMorningPersonPlanned:   data.PartnerMorningPersonPlanned,
		PartnerAfternoonPersonPlanned: data.PartnerAfternoonPersonPlanned,
		PartnerEveningPersonPlanned:   data.PartnerEveningPersonPlanned,
		PartnerMorningPerson:          data.PartnerMorningPerson,
		PartnerAfternoonPerson:        data.PartnerAfternoonPerson,
		PartnerEveningPerson:          data.PartnerEveningPerson,
		PartnerOvertimeHour:           data.PartnerOvertimeHour,
		PartnerCoefficient:            data.PartnerCoefficient,
	}

	log.For(c).Info("[timesheets-detail] process success", log.Field("user_id", userID), log.Field("task_id", id), log.Field("resp", resp))
	c.JSON(http.StatusOK, &resources.Response{
		Data: resp,
	})
}
