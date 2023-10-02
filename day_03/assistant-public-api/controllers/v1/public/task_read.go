package public

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ideal-forward/assistant-public-api/controllers/resources"
	"github.com/ideal-forward/assistant-public-api/entities"
	"github.com/ideal-forward/assistant-public-api/middlewares"
	"github.com/ideal-forward/assistant-public-api/pkg/log"
	"github.com/ideal-forward/assistant-public-api/services"
)

func (h Handler) ReadTask(c *gin.Context) {
	ctx := c.Request.Context()
	userID, _, _ := middlewares.ParseToken(c)

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Errorf("invalid parameter"))
		return
	}

	data, err := h.Task.Read(ctx, id)
	if err != nil {
		log.For(c).Error("[task-detail] query failed", log.Field("user_id", userID), log.Field("task_id", id), log.Err(err))
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	task := &resources.Task{
		ID:          data.ID,
		Code:        data.Code,
		Name:        data.Name,
		ProjectID:   data.GetProject().GetID(),
		ProjectName: data.GetProject().GetName(),
		Project: &resources.Project{
			ID:        data.GetProject().GetID(),
			Name:      data.GetProject().GetName(),
			Code:      data.GetProject().GetCode(),
			Address:   data.GetProject().GetAddress(),
			StartDate: services.FormatTimestamp(data.GetProject().GetStartDate()),
			EndDate:   services.FormatTimestamp(data.GetProject().GetEndDate()),
		},
		Area: &resources.ProjectArea{
			ID:   data.GetArea().GetID(),
			Name: data.GetArea().GetName(),
		},
		Phase: &resources.ProjectPhase{
			ID:   data.GetPhase().GetID(),
			Name: data.GetPhase().GetName(),
		},
		Creator: &resources.User{
			ID:   data.GetCreator().GetID(),
			Name: data.GetCreator().GetName(),
		},
		Executor: &resources.Executor{
			ID:   data.GetExecutor().GetID(),
			Name: data.GetExecutor().GetName(),
		},
		Reporter: &resources.User{
			ID:   data.GetReporter().GetID(),
			Name: data.GetReporter().GetName(),
		},
		Acceptor: &resources.User{
			ID:   data.GetAcceptor().GetID(),
			Name: data.GetAcceptor().GetName(),
		},
		StartTime:     services.FormatTimestamp(data.StartTime),
		EndTime:       services.FormatTimestamp(data.EndTime),
		Quantity:      data.Quantity,
		Price:         data.Price,
		Unit:          data.Unit,
		Description:   data.Description,
		Status:        entities.NewTaskStatus(data.Status).String(),
		PriorityLevel: entities.NewPriorityLevel(data.PriorityLevel).String(),
		Star:          data.Star,
	}

	if data.GetParentTask() != nil {
		task.ParentTask = &resources.Task{
			ID:   data.GetParentTask().GetID(),
			Name: data.GetParentTask().GetName(),
		}
	}

	if len(data.ChildTasks) > 0 {
		task.ChildTask = &resources.Task{
			ID:   data.ChildTasks[0].GetID(),
			Name: data.ChildTasks[0].GetName(),
		}
	}

	for _, val := range data.Comments {
		task.Comments = append(task.Comments, &resources.Comment{
			ID:       val.ID,
			UserID:   val.CreatedBy,
			UserName: val.Creator.GetName(),
			Creator: &resources.User{
				ID:   val.GetCreator().GetID(),
				Name: val.GetCreator().GetName(),
			},
			Msg:  val.Comment,
			Time: services.FormatTime(val.CreatedAt),
		})
	}

	for _, val := range data.AttachFiles {
		if val.Type == entities.AttachFileByOwner {
			task.OwnerImages = append(task.OwnerImages, &resources.Image{
				ID:        val.ID,
				Name:      val.FileName,
				Thumbnail: h.FileNamer.ToPublicFile(val.Thumbnail),
				FullFile:  h.FileNamer.ToPublicFile(val.FilePath),
			})
			continue
		}

		task.ExecutorImages = append(task.ExecutorImages, &resources.Image{
			ID:        val.ID,
			Name:      val.FileName,
			Thumbnail: h.FileNamer.ToPublicFile(val.Thumbnail),
			FullFile:  h.FileNamer.ToPublicFile(val.FilePath),
		})
	}

	log.For(c).Info("[task-detail] process success", log.Field("user_id", userID), log.Field("task_id", id), log.Field("resp", task))
	c.JSON(http.StatusOK, &resources.Response{
		Data: task,
	})
}
