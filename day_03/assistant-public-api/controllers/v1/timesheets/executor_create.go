package timesheets

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ideal-forward/assistant-public-api/controllers/resources"
	"github.com/ideal-forward/assistant-public-api/entities"
	"github.com/ideal-forward/assistant-public-api/middlewares"
	"github.com/ideal-forward/assistant-public-api/pkg/http_parser"
)

func (h Handler) CreateExecutor(c *gin.Context) {
	ctx := c.Request.Context()
	req := &resources.CreateTimesheetsExecutorRequest{}

	timesheetsID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Errorf("invalid parameter"))
		return
	}

	err = http_parser.BindAndValid(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	userID, _, _ := middlewares.ParseToken(c)

	data := &entities.TimesheetsExecutor{
		Base: entities.Base{
			CreatedBy: userID,
		},
		TimesheetsID:           timesheetsID,
		ExecutorID:             req.ExecutorID,
		MorningPersonPlanned:   req.MorningPersonPlanned,
		AfternoonPersonPlanned: req.AfternoonPersonPlanned,
		EveningPersonPlanned:   req.EveningPersonPlanned,
		MorningPerson:          req.MorningPerson,
		AfternoonPerson:        req.AfternoonPerson,
		EveningPerson:          req.EveningPerson,
		OvertimeHour:           req.OvertimeHour,
		Coefficient:            req.Coefficient,
	}
	id, err := h.Timesheets.CreateExecutor(ctx, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, &resources.Response{
		Data: &resources.CreateResponse{
			ID:   id,
			UUID: data.UUID,
		},
	})
}
