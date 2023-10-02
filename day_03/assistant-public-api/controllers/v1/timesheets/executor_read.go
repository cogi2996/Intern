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

func (h Handler) ReadExecutor(c *gin.Context) {
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

	data, err := h.Timesheets.ReadExecutor(ctx, id)
	if err != nil {
		log.For(c).Error("[timesheets-detail-executor] query failed", log.Field("user_id", userID), log.Field("timesheets_id", timesheetsID), log.Field("id", id), log.Err(err))
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	resp := &resources.TimesheetsExecutor{
		ID: data.ID,
		Creator: &resources.User{
			ID: data.GetCreator().GetID(),
		},
		Executor: &resources.Executor{
			Name: data.GetExecutor().GetName(),
		},
		MorningPersonPlanned:   data.MorningPersonPlanned,
		AfternoonPersonPlanned: data.AfternoonPersonPlanned,
		EveningPersonPlanned:   data.EveningPersonPlanned,
		MorningPerson:          data.MorningPerson,
		AfternoonPerson:        data.AfternoonPerson,
		EveningPerson:          data.EveningPerson,
		OvertimeHour:           data.OvertimeHour,
		Coefficient:            data.Coefficient,
	}

	log.For(c).Info("[timesheets-detail-executor] process success", log.Field("user_id", userID), log.Field("timesheets_id", id), log.Field("resp", resp))
	c.JSON(http.StatusOK, &resources.Response{
		Data: resp,
	})
}
