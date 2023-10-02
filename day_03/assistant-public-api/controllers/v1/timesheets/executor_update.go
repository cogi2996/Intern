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
)

func (h Handler) UpdateExecutor(c *gin.Context) {
	ctx := c.Request.Context()
	req := &resources.UpdateTimesheetsExecutorRequest{}

	err := http_parser.BindAndValid(c, req, "MaxSize", "DateFormat")
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	userID, _, _ := middlewares.ParseToken(c)

	id, err := strconv.ParseInt(c.Param("executor_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Errorf("invalid parameter"))
		return
	}

	newData := &entities.TimesheetsExecutor{
		Base: entities.Base{
			UpdatedBy: userID,
		},
		ID:                     id,
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

	err = h.Timesheets.UpdateExecutor(ctx, newData)
	if err != nil {
		log.For(c).Error("[timesheets-executor-update] update new data failed", log.Field("user_id", userID), log.Field("task_id", id), log.Err(err))
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	h.makeUpdateExecutorComment(ctx, userID, id, req.Reason)

	log.For(c).Info("[timesheets-executor-update] insert success", log.Field("user_id", userID), log.Field("task_id", id), log.Err(err))
	c.JSON(http.StatusOK, &resources.Response{
		Data: &resources.Empty{},
	})
}

func (h Handler) makeUpdateExecutorComment(ctx context.Context, userID, timesheetsID int64, reason string) error {
	user, err := h.User.Read(ctx, userID)
	if err != nil {
		return err
	}

	_, err = h.Timesheets.CreateComment(ctx, &entities.TimesheetsComment{
		Base: entities.Base{
			CreatedBy: userID,
		},
		TimesheetsID: timesheetsID,
		Comment:      user.GetName() + " cập nhật bảng với lí do \"" + reason + "\"",
	})

	return err
}
