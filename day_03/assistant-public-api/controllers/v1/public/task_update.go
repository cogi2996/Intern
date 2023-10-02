package public

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

func (h Handler) UpdateTask(c *gin.Context) {
	ctx := c.Request.Context()
	req := &resources.UpdateTaskRequest{}

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

	currentData, err := h.Task.Read(ctx, id)
	if err != nil {
		log.For(c).Error("[task-update] query old data failed", log.Field("user_id", userID), log.Field("task_id", id), log.Err(err))
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	newData := &entities.Task{
		ID:            id,
		Name:          req.Name,
		ProjectID:     req.ProjectID,
		AreaID:        req.AreaID,
		PhaseID:       req.PhaseID,
		ExecutedBy:    req.ExecutorID,
		ReportBy:      req.ReporterID,
		AcceptedBy:    req.AcceptorID,
		Quantity:      req.Quantity,
		Price:         req.Price,
		Unit:          req.Unit,
		Description:   req.Description,
		Status:        req.Status,
		PriorityLevel: entities.NewPriorityLevelFromString(req.PriorityLevel).Value(),
	}
	newData.StartTime, _ = services.StringToTimestamp(req.StartTime)
	newData.EndTime, _ = services.StringToTimestamp(req.EndTime)
	if newData.AcceptedBy == 0 {
		newData.AcceptedBy = userID
	}

	err = h.Task.Update(ctx, newData)
	if err != nil {
		log.For(c).Error("[task-update] update new data failed", log.Field("user_id", userID), log.Field("task_id", id), log.Err(err))
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	comment, err := h.makeUpdateComment(ctx, userID, currentData, newData)
	if err == nil {
		h.TaskComment.Create(ctx, &entities.TaskComment{
			Base: entities.Base{
				CreatedBy: userID,
			},
			TaskID:  id,
			Comment: comment,
		})
	}

	log.For(c).Info("[task-update] insert failed", log.Field("user_id", userID), log.Field("task_id", id), log.Err(err))
	c.JSON(http.StatusOK, &resources.Response{
		Data: &resources.Empty{},
	})
}

func (h Handler) makeUpdateComment(ctx context.Context, userID int64, currentData, newData *entities.Task) (string, error) {
	if (currentData.Quantity == newData.Quantity) && (currentData.Price == newData.Price) && (currentData.PriorityLevel == newData.PriorityLevel) {
		return "", fmt.Errorf("")
	}

	user, err := h.User.Read(ctx, userID)
	if err != nil {
		return "", err
	}

	result := fmt.Sprintf("%s đã cập nhật ", user.Name)
	if currentData.Quantity != newData.Quantity {
		result = result + fmt.Sprintf(", số lượng %d thành %d", currentData.Quantity, newData.Quantity)
	}

	if currentData.Price != newData.Price {
		result = result + fmt.Sprintf(", giá %d thành %d", currentData.Price, newData.Price)
	}

	if currentData.PriorityLevel != newData.PriorityLevel {
		result = result + fmt.Sprintf(", độ ưu tiên %s thành %s", entities.NewPriorityLevel(currentData.PriorityLevel).Comment(), entities.NewPriorityLevel(newData.PriorityLevel).Comment())
	}

	return result, nil
}
