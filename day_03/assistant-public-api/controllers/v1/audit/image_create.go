package audit

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ideal-forward/assistant-public-api/controllers/resources"
	"github.com/ideal-forward/assistant-public-api/entities"
	"github.com/ideal-forward/assistant-public-api/middlewares"
)

const HeaderImage = "image"

func (h Handler) CreateTaskImage(c *gin.Context) {
	ctx := c.Request.Context()

	taskID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Errorf("invalid parameter"))
		return
	}

	userID, _, _ := middlewares.ParseToken(c)

	data := &entities.AuditTaskAttachFile{
		Base: entities.Base{
			CreatedBy: userID,
		},
		TaskID: taskID,
		Type:   entities.AttachFileByOwner,
	}

	_, err = h.Task.Read(ctx, taskID)
	if err == nil {
		data.Type = entities.AttachFileByExecutor
	}

	data.FileName, data.FilePath, data.Thumbnail, err = h.FileUploader.UploadImage(c, HeaderImage)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	data.FilePath = h.FileNamer.ToRelativeFile(data.FilePath)
	data.Thumbnail = h.FileNamer.ToRelativeFile(data.Thumbnail)

	id, err := h.TaskImage.Create(ctx, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, &resources.Response{
		Data: &resources.CreateTaskImageResponse{
			ID:        id,
			UUID:      data.UUID,
			FilePath:  h.FileNamer.ToPublicFile(data.FilePath),
			Thumbnail: h.FileNamer.ToPublicFile(data.Thumbnail),
		},
	})
}
