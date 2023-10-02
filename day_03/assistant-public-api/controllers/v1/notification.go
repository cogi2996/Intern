package v1

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ideal-forward/assistant-public-api/controllers/resources"
	"github.com/ideal-forward/assistant-public-api/services"
)

func (h Handler) ReadNotification(c *gin.Context) {
	ctx := c.Request.Context()

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Errorf("invalid parameter"))
		return
	}

	data, err := h.Notification.Read(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, &resources.Response{
		Data: &resources.Notification{
			ID:         data.ID,
			UUID:       data.UUID,
			Title:      data.Title,
			ReviewText: data.Message,
			Message:    data.Message,
			SendTime:   services.FormatTimestamp(data.CreatedAt.UnixMilli()),
			PublicTask: &resources.Task{
				ID:   data.GetPublicTask().GetID(),
				Name: data.GetPublicTask().GetName(),
			},
			Sender: &resources.User{
				ID:   data.GetCreator().GetID(),
				Name: data.GetCreator().GetName(),
			},
		},
	})
}

func (h Handler) ListNotification(c *gin.Context) {
	ctx := c.Request.Context()

	data, err := h.Notification.GetList(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	resp := &resources.ListNotificationResponse{
		Notifications: make([]*resources.Notification, 0),
	}
	for _, val := range data {
		resp.Notifications = append(resp.Notifications, &resources.Notification{
			ID:         val.ID,
			UUID:       val.UUID,
			Title:      val.Title,
			ReviewText: val.Message,
			Message:    val.Message,
			SendTime:   services.FormatTimestamp(val.CreatedAt.UnixMilli()),
			PublicTask: &resources.Task{
				ID:   val.GetPublicTask().GetID(),
				Name: val.GetPublicTask().GetName(),
			},
			Sender: &resources.User{
				ID:   val.GetCreator().GetID(),
				Name: val.GetCreator().GetName(),
			},
		})
	}

	c.JSON(http.StatusOK, &resources.Response{
		Data: resp,
	})
}
