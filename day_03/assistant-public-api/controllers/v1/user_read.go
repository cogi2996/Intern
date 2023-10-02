package v1

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ideal-forward/assistant-public-api/controllers/resources"
)

func (h Handler) ReadUser(c *gin.Context) {
	ctx := c.Request.Context()

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Errorf("invalid parameter"))
		return
	}

	data, err := h.UserService.Read(ctx, id)
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
