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

func (h Handler) ListExecutor(c *gin.Context) {
	ctx := c.Request.Context()

	timesheetsID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Errorf("invalid parameter"))
		return
	}

	userID, _, _ := middlewares.ParseToken(c)
	data, err := h.Timesheets.ListExecutor(ctx, timesheetsID)
	if err != nil {
		log.For(c).Error("[timesheets-executor-list] query failed", log.Field("user_id", userID), log.Field("timesheets_id", timesheetsID), log.Err(err))
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	resp := make([]*resources.TimesheetsExecutor, 0)
	for _, val := range data {
		resp = append(resp, &resources.TimesheetsExecutor{
			ID: val.ID,
			Creator: &resources.User{
				ID:   val.GetCreator().GetID(),
				Name: val.GetCreator().GetName(),
			},
			Executor: &resources.Executor{
				ID:   val.GetExecutor().GetID(),
				Name: val.GetExecutor().GetName(),
			},
			MorningPersonPlanned:   val.MorningPersonPlanned,
			AfternoonPersonPlanned: val.AfternoonPersonPlanned,
			EveningPersonPlanned:   val.EveningPersonPlanned,
			MorningPerson:          val.MorningPerson,
			AfternoonPerson:        val.AfternoonPerson,
			EveningPerson:          val.EveningPerson,
			OvertimeHour:           val.OvertimeHour,
			Coefficient:            val.Coefficient,
		})
	}

	c.JSON(http.StatusOK, &resources.Response{
		Data: resp,
	})
}
