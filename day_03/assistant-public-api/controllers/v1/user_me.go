package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ideal-forward/assistant-public-api/controllers/resources"
	"github.com/ideal-forward/assistant-public-api/middlewares"
)

func (h Handler) ReadSelfProfile(c *gin.Context) {
	ctx := c.Request.Context()

	userID, _, _ := middlewares.ParseToken(c)

	data, err := h.UserService.Read(ctx, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, &resources.Response{
		Data: &resources.User{
			ID:      data.ID,
			Code:    data.Code,
			Name:    data.Name,
			Phone:   data.Phone,
			Email:   data.Email,
			Address: data.Address,
			Manager: &resources.User{
				ID:   data.GetManager().GetID(),
				Code: data.GetManager().GetCode(),
				Name: data.GetManager().GetName(),
			},
		},
	})
}
