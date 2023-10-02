package private

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ideal-forward/assistant-public-api/controllers/resources"
	"github.com/ideal-forward/assistant-public-api/entities"
)

func (h Handler) ListTaskImage(c *gin.Context) {
	ctx := c.Request.Context()

	taskID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Errorf("invalid parameter"))
		return
	}

	images, err := h.TaskImage.List(ctx, taskID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	resp := &resources.ListTaskImageResponse{}
	for _, val := range images {
		el := &resources.Image{
			ID:        val.ID,
			Name:      val.FileName,
			Thumbnail: h.FileNamer.ToPublicFile(val.Thumbnail),
			FullFile:  h.FileNamer.ToPublicFile(val.FilePath),
			Creator: &resources.User{
				ID:   val.GetCreator().GetID(),
				Name: val.GetCreator().GetName(),
			},
		}

		if val.Type == entities.AttachFileByOwner {
			resp.OwnerImages = append(resp.OwnerImages, el)
			continue
		}
		resp.ExecutorImages = append(resp.ExecutorImages, el)
	}

	c.JSON(http.StatusOK, &resources.Response{
		Data: resp,
	})
}
